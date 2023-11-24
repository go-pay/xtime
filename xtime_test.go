package xtime

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestXtime(t *testing.T) {
	minutes := Time(1609066441).Time().Add(time.Minute * 30).Sub(time.Now()).Minutes()
	log.Println(minutes)
	if minutes < 0 { // 30分钟超时
		//更新订单状态为订单超时
		log.Println("超时")
	}
}

type TimeParser struct {
	T1 Duration `json:"t1"`
	T2 Duration `json:"t2"`
	T3 Duration `json:"t3"`
	T4 Duration `json:"t4"`
}

func TestParseTime(t *testing.T) {
	parseText := `
{
    "t1":"10s",
    "t2":"1m20s",
    "t3":"5m",
    "t4":"1h10m10s"
}`
	tp := new(TimeParser)
	err := json.Unmarshal([]byte(parseText), tp)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v", tp)

	log.Printf("t1: %s", tp.T1.UnitTime())
	log.Printf("t2: %s", tp.T2.UnitTime())
	log.Printf("t3: %s", tp.T3.UnitTime())
	log.Printf("t4: %s", tp.T4.UnitTime())
}
