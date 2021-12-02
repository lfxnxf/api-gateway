package dispersed_lock

import (
	"context"
	"fmt"
	"git.inke.cn/inkelogic/daenerys"
	"testing"
	"time"
)

func init() {
	daenerys.Init(
		daenerys.ConfigPath("./config/config.toml"),
	)
}

func TestLock(t *testing.T) {
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			dispersedLock := New(context.Background(), "im.family.cache", "sg:test:lock", 1)
			dispersedLock.LoopLock(context.Background(), 1)
			fmt.Println("这里是第几个线程", i)
			fmt.Println("执行时间", time.Now().UnixNano() / 1000000)
			time.Sleep(1 * time.Second)
			c <- true
			dispersedLock.Unlock(context.Background())
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}
