package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("当前时间戳:",time.Now().Unix())
    fmt.Println("当前时间戳格式化:",time.Now().Format("2006-01-02 15:04:05"))
    str_time := time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")
    fmt.Println("指定时间戳格式化:",str_time)
    //（月份必须是month）
    the_time := time.Date(2016, 1, 5, 0, 0, 0, 0, time.Local)
    unix_time1 := the_time.Unix()
    fmt.Println("时间反格式化成时间戳:",unix_time1)
    the_time, err := time.ParseInLocation("2006-01-02", "2015-12-01", time.Local)
    if err != nil {
        fmt.Println(err)
    }
    unix_time2 := the_time.Unix()
    fmt.Println("指定时间反格式化成时间戳:",unix_time2)
}
