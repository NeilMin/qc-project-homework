package main

import (
    "flag"
    "fmt"
    "os"
)
	
func main() {
    var t int
    flag.Usage = func() {
        fmt.Printf("Usage of %s:\n", os.Args[0])
        fmt.Printf("使用-type来决定输出唐诗或宋词，使用-h获得帮助。\n")
        flag.PrintDefaults()
    }

    flag.IntVar(&t, "type", 100, "决定输出唐诗或宋词")

    flag.Parse()

    if t == 1 {
        fmt.Println("大漠孤烟直,长河落日圆。")
    }

    if t == 2 {
        fmt.Println("琵琶弦上说相思，当时明月在，曾照彩云归。")
    }
}
