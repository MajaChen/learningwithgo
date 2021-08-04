package golang

import (
	"testing"
	"time"
)

func TestUnixTime(t *testing.T) {
	past := time.Now().Unix()

	time.Sleep(30 * time.Second)

	current := time.Now().Unix()
	t.Errorf("time elasped %v seconds\n", current-past)
}

func TestTimeTransform(t *testing.T) {
	//时间到时间戳
	now := time.Now()
	timestamp := now.Unix()
	t.Errorf("now is %v\n", now)

	//时间戳到时间
	tm := time.Unix(timestamp, 0)

	// 时间到字符串
	str := tm.Format(time.RFC3339)

	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm, _ = time.Parse(time.RFC3339, str)
	t.Errorf("tm is %v\n", tm)

}
