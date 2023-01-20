package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("error, need help? -h")
	} else if os.Args[1] == "now" {
		fmt.Println(0)
		return
	} else if os.Args[1] == "-h" {
		fmt.Println("参数：\n   now 立刻执行\n   执行时间的小时数\n   执行时间的小时数 执行时间的分钟数")
	} else {
		var M = "0"
		if len(os.Args) == 3 {
			M = os.Args[2]
		}
		h, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("error, need help? -h")
			return
		}
		m, err := strconv.Atoi(M)
		if err != nil {
			fmt.Println("error, need help? -h")
			return
		}
		for {
			if h <= time.Now().Hour() {
				for {
					if m <= time.Now().Minute() {
						fmt.Print(0)
						return
					} else {
						time.Sleep(10 * 100 * time.Millisecond)
					}
				}
			} else {
				time.Sleep(30 * 100 * time.Millisecond)
			}
		}
	}
}
