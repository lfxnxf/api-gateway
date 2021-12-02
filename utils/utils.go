package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type WrapResp struct {
	Code int         `json:"dm_error"`
	Msg  string      `json:"error_msg"`
	Data interface{} `json:"data"`
}

func InStringArray(item string, items []string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func InInt64Array(item int64, items []int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func Random(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}


func String2Unix(stringTime string) int64 {
	loc, _ := time.LoadLocation("Local")

	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)

	return theTime.Unix()
}

func NowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 分组
func SplitArray(arr []string, num int64) (splits [][]string) {
	var ltSlices = make([][]string, 0)
	max := int64(len(arr))
	if max < num || num <= 0 {
		ltSlices[0] = arr
		return ltSlices
	}
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	var start, end, i int64
	for i = 1; i <= num; i++ {
		end = i * quantity
		if i != num {
			ltSlices = append(ltSlices, arr[start:end])
		} else {
			ltSlices = append(ltSlices, arr[start:])
		}
		start = i * quantity
	}

	return ltSlices
}

// 根据生日计算年龄
func GetAgeFromBirthday(birthday string) (int, error) {
	now := time.Now()
	t, err := time.ParseInLocation("2006-01-02 15:04:05", birthday, time.Local)
	if err != nil {
		return 0, err
	}

	age := 0
	if t.Before(now) {
		age = now.Year() - t.Year()
		next := t.AddDate(age, 0, 0)
		if now.Before(next) {
			age = age - 1
		}
	}

	return age, nil
}

func SplitIntArray(arr []int64, num int64) [][]int64 {
	max := int64(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]int64{arr}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]int64, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= num; i++ {
		end = i * quantity
		if i != num {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * quantity
	}
	return segments
}

// 从src中去除except
func ExceptArray(src []int64, except []int64) []int64 {
	newArray := make([]int64, 0)

	excMap := make(map[int64]bool)
	for _, value := range except {
		excMap[value] = true
	}

	for _, value := range src {
		if _, ok := excMap[value]; ok {
			continue
		}

		newArray = append(newArray, value)
	}

	return newArray
}

//int64 数组去重
func UniqueInt64Slice(src []int64) []int64 {
	result := make([]int64, 0)
	tempMap := map[int64]byte{}
	for _, e := range src {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

//获取第二天凌晨时间戳
func GetTomorrowZeroTimeUnix() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	return t.Unix() + 1
}

func GetNowMonth() string {
	//获取当前月份
	return fmt.Sprintf("%d-%s", time.Now().Year(), time.Now().Format("01"))
}

func GetLastMonth() string {
	return fmt.Sprintf("%d-%s", time.Now().Year(), time.Now().AddDate(0, -1, 0).Format("01"))
}

func GetLastMonthLastDayUnix() int64 {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	end := thisMonth.AddDate(0, 0, -1).Unix() + 86399
	return end
}

func GetFirstDayUnix() int64 {
	return 0
}

// 调用栈
func Stack() []byte {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return buf[:n]
}

/**
获取本周周一的日期
*/
func GetNowDateOfWeek() (weekMonday string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format("2006-01-02")
	return
}

/**
获取上周的周一日期
*/
func GetLastWeekFirstDate() (weekMonday string) {
	thisWeekMonday := GetNowDateOfWeek()
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -7)
	weekMonday = lastWeekMonday.Format("2006-01-02")
	return
}

func YearWeekByDate(date string) string {
	week := WeekByDate(date)
	l, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02", date, l)
	return fmt.Sprintf("%d-%d", t.Year(), week)
}

//判断时间是当年的第几周
func WeekByDate(date string) int {
	l, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02", date, l)
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())
	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return week
}

func Md5(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

// 生成指定长度随机字符串
// 0 - 数字
// 1 - 小写字母
// 2 - 大写字母
// 3 - 数字、小写、大写字母
func GenRandString(size int, kind int) string {
	randKind := kind
	kinds := [][]int{{10, 48}, {26, 97}, {26, 65}}
	result := make([]byte, size)

	isAll := kind > 2 || kind < 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		if isAll { // random randKind
			randKind = rand.Intn(3)
		}
		scope, base := kinds[randKind][0], kinds[randKind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}

	return string(result)
}

// 字符串编码SHA1
func GenFaceInitSignSHA(src string) string {
	h := sha1.New()
	h.Write([]byte(src))
	bs := h.Sum(nil)
	return fmt.Sprintf("%X", bs)
}

// 逗号分割string转[]int64
func StringToInt64Slice(src string) []int64 {
	var result []int64
	stringList := strings.Split(src, ",")
	for _, v := range stringList {
		vInt, _ := strconv.ParseInt(v, 10, 64)
		result = append(result, vInt)
	}
	return result
}