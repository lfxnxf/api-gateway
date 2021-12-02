package dispersed_lock

import (
	"github.com/garyburd/redigo/redis"
	"github.com/lfxnxf/frame/logic/inits/proxy"
	"github.com/lfxnxf/school/api-gateway/utils"
	"golang.org/x/net/context"
	"strings"
	"sync"
	"time"
)

const (
	lockScript = "if redis.call('get', KEYS[1]) == ARGV[1] " +
		"then redis.call('del', KEYS[1]) return 1 " +
		"else " +
		"return 0 " +
		"end"

	lockMaxLoopNum = 1000 //加锁最大循环数量
)

var clientMap sync.Map
var scriptMap sync.Map

type option func() (bool, error)

type DispersedLock struct {
	key          string       `json:"key"`           // 锁key
	value        int64        `json:"value"`         // 锁的值，随机数
	expire       int64        `json:"expire"`        // 锁过期时间
	lockClient   *proxy.Redis `json:"lock_client"`   // 锁客户端，暂时只有redis
	scriptString string       `json:"script_string"` // lua脚本
	options      []option     `json:"options"`       // 事件
}

func New(ctx context.Context, clientKey, key string, expire int64) *DispersedLock {
	d := &DispersedLock{
		key:    key,
		expire: expire,
		value:  utils.Random(100000000, 999999999), // 随机值作为锁的值
	}

	//初始化连接
	client, _ := clientMap.LoadOrStore(clientKey, proxy.InitRedis(clientKey))
	d.lockClient = client.(*proxy.Redis)

	//初始化lua script
	scriptString, _ := scriptMap.LoadOrStore(clientKey, d.getScript(ctx))
	d.scriptString = scriptString.(string)

	return d
}

func (d *DispersedLock) getScript(ctx context.Context) string {
	scriptString, _ := redis.String(d.lockClient.Do(ctx, "SCRIPT", "LOAD", lockScript))
	return scriptString
}

//注册事件
func (d *DispersedLock) RegisterOptions(f ...option) {
	d.options = append(d.options, f...)
}

//加锁
func (d *DispersedLock) Lock(ctx context.Context) bool {
	args := []interface{}{
		d.key,    // 锁的key
		d.value,  // 锁的随机值
		"ex",     // 设置过期
		d.expire, // 过期时间
		"nx",     // 设置不可重复设置
	}
	ret, _ := redis.String(d.lockClient.Do(ctx, "Set", args...))
	if strings.ToLower(ret) == "ok" {
		return true
	} else {
		return false
	}
}

//循环加锁
func (d *DispersedLock) LoopLock(ctx context.Context, sleepTime int) bool {
	t := time.NewTicker(time.Duration(sleepTime) * time.Millisecond)
	w := utils.NewWhile(lockMaxLoopNum)
	w.For(func() {
		if d.Lock(ctx) {
			t.Stop()
			w.Break()
		} else {
			<-t.C
		}
	})
	if !w.IsNormal() {
		return false
	}
	return true
}

//循环获取锁并且绑定事件
//eg:单个线程获取缓存、其它线程等待
func (d *DispersedLock) LoopLockWithOption(ctx context.Context, sleepTime int) (bool, error) {
	t := time.NewTicker(time.Duration(sleepTime) * time.Millisecond)
	w := utils.NewWhile(lockMaxLoopNum)
	var err error
	w.For(func() {
		locked := d.Lock(ctx)
		if locked { // 获取到锁，跳出循环
			t.Stop()
			w.Break()
		}

		var flag bool
		for _, option := range d.options {
			flag, err = option()
			if err != nil { //事件代码出现异常，跳出循环
				t.Stop()
				w.Break()
			}
			if !flag {
				break
			}
		}

		//所有事件全部为true，不用等到获取锁，直接跳出
		if flag {
			t.Stop()
			w.Break()
		}

		<-t.C
	})
	return true, err
}

//解锁
func (d *DispersedLock) Unlock(ctx context.Context) bool {
	args := []interface{}{
		d.scriptString, // 脚本sha1值
		1,              // key的length
		d.key,          // key
		d.value,        // 脚本中的argv
	}
	flag, _ := redis.Int64(d.lockClient.Do(ctx, "EvalSha", args...))
	return lockRes(flag)
}

func lockRes(flag int64) bool {
	if flag > 0 {
		return true
	} else {
		return false
	}
}
