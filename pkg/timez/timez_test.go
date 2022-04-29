package timez

import (
	"testing"
	"time"
)

func TestTimeHello(t *testing.T) {
	// 获取当前时间
	now := time.Now()
	t.Logf("%#v \n", now) // time.Date(2022, time.April, 25, 9, 14, 34, 379809500, time.Local)
	// 获取年
	t.Log(now.Year()) // 2022
	// 获取月
	t.Log(now.Month()) // April
	// 获取日
	t.Log(now.Day()) // 25
	// 获取小时
	t.Log(now.Hour()) // 9
	// 获取分钟
	t.Log(now.Minute()) // 14
	// 获取秒
	t.Log(now.Second()) // 34
	// 获取纳秒
	t.Log(now.Nanosecond()) // 379809500
	// 时间戳
	t.Log(now.UnixMicro()) // 1650849274379809
	t.Log(now.UnixMilli()) // 1650849274379
}

func TestTimeParse(t *testing.T) {
	// 获取当前时间
	now := time.Now()

	// 将时间格式化成字符串
	timeStr := now.Format("2006年1月2日 03:04:05")
	t.Log(timeStr) // 2022年4月25日 09:18:33
	// 将字符串解析成时间
	timeStr2 := "2022-04-25 09:18:30"
	// 时间格式化样式
	layout := "2006-01-02 03:04:05"
	// 指定时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Log("parse in location failed: ", err)
		return
	}
	// 解析指定时区的时间
	parseTime, err := time.ParseInLocation(layout, timeStr2, loc)
	// 不指定时区, 返回UTC时间
	// parseTime, err := time.Parse(layout, timeStr2)
	if err != nil {
		t.Log("time parse failed: ", err)
	}
	t.Log(parseTime) // 2022-04-25 09:18:30 +0800 CST
}

func TestTimeStamp(t *testing.T) {
	// 获取当前时间
	now := time.Now()

	// 获取秒时间戳
	// 当前时间距离1970年1月1日所过去多少秒
	timeUnix := now.Unix()
	t.Log("时间转换为秒时间戳: ", timeUnix) // 1650849690
	// 获取纳秒时间戳
	// 当前时间距离1970年1月1日所过去多少纳秒
	timeUnixNano := now.UnixNano()
	t.Logf("时间转换为纳秒时间戳: %d", timeUnixNano) // 1650849690853584200
	// 将时间戳转为时间
	/*
	 我们指定一个时间戳, 然后将它转换为对应的时间
	*/
	t1 := time.Unix(timeUnix+int64(60*60), 0)
	t.Log(t1) // 2022-04-25 10:21:30 +0800 CST
}

func TestTimeCalc(t *testing.T) {
	// 获取当前时间
	now := time.Now()

	// 时间运算
	t.Log("当前时间: ", now)
	// add
	// 1小时后
	t.Log("1小时后: ", now.Add(time.Hour*1))
	// 1天后
	t.Log("1天后: ", now.Add(time.Hour*24))
	// 1小时前
	t.Log("1小时前: ", now.Add(-time.Hour*1))
	// 1天前
	t.Log("1天前: ", now.Add(-time.Hour*24))
	// sub
	time1 := time.Now()
	time2 := time.Now().Add(-time.Hour * 10)
	subTime := time1.Sub(time2)
	t.Log(subTime.Hours())
	t.Log(subTime.Minutes())
}
