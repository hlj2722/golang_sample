package main

import (
	"fmt"
	"strings"
	"time"
)
/*time包中的Add和Sub的用法，
 *Add用于计算某个时间之前和之后的时间点，
 *Sub用于计算两个时间差
*/
func main() {
	// Add 时间相加
	now := time.Now()
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 10分钟前
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(10*m)
	fmt.Println("10分钟前的时间:",m1)

	// 8个小时前
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(8 * h)
	fmt.Println("8个小时前的时间:",h1)

	// 一天前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d)
	fmt.Println("一天前的时间:",d1)

	printSplit(50)

	// 10分钟后
	mm, _ := time.ParseDuration("1m")
	mm1 := now.Add(10*mm)
	fmt.Println("10分钟后的时间:",mm1)

	// 8小时后
	hh, _ := time.ParseDuration("1h")
	hh1 := now.Add(8*hh)
	fmt.Println("8个小时后的时间:",hh1)

	// 一天后
	dd, _ := time.ParseDuration("24h")
	dd1 := now.Add(dd)
	fmt.Println("一天后的时间:",dd1)

	printSplit(50)

	// Sub 计算两个时间差
	fmt.Println("之前的时间为:",m1)
	subM := now.Sub(m1)
	fmt.Println("现在时间相对之前时间为:",subM.Minutes(), "分钟之前")
	printSplit(50)
	fmt.Println("之前时间为:",h1)
	sumH := now.Sub(h1)
	fmt.Println("现在时间相对之前时间为:",sumH.Hours(), "小时")
	printSplit(50)
	fmt.Println("之前时间为:",d1)
	sumD := now.Sub(d1)
	fmt.Printf("现在时间相对之前时间为:%v 天\n", sumD.Hours()/24)

}

func printSplit(count int) {
	fmt.Println(strings.Repeat("#", count))
}