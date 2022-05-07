package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFunc(t *testing.T) {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TestGetTimestamp(t *testing.T) {
	//基于时间对象获取时间戳
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //将时间对象转变为时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	//使用time.Unix()函数可以将时间戳转变为时间格式。
	timeObj := time.Unix(timestamp1, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func TestTimeAdd(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Minute)
	fmt.Println(later)
}

func TestTimeSub(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Minute)
	fmt.Println(later)

	fmt.Println(later.Sub(now))
}

func TestTimeEqual(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Minute)
	fmt.Println(later)

	fmt.Println(later.Equal(now))
}

func TestTimeBeforeAfter(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Minute)
	fmt.Println(later)

	fmt.Println(later.Before(now))
	fmt.Println(later.After(now))
}
