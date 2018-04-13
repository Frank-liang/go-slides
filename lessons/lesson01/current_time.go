package main

import (
	"fmt"
	"time"
)

func main() {
	var t = time.Now()
	//t := time.Now()
	fmt.Println(t) // 2018-04-13 22:46:24.23096157 +0800 CST m=+0.000607127

	y := t.Year()                       // 年
	m := t.Month()                      // 月
	d := t.Day()                        // 日
	h := t.Hour()                       // 时
	min := t.Minute()                   // 分
	s := t.Second()                     // 秒
	fmt.Printf("%d-%d-%d\n", y, m, d)   // 2018-04-01
	fmt.Printf("%d%02d%02d\n", y, m, d) // 20180401

	// 2018-04-13 22:46:24
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
		y, m, d,
		h, min, s)
}
