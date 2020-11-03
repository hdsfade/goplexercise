//@author: hdsfade
//@date: 2020-11-02-17:05
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	day   = 60 * 1000000 * time.Nanosecond
	month = 30 * 24 * day
	year  = 12 * month
)

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		//输出一个月内的结果
		if pasttime(item) < month {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

func pasttime(item *Issue) time.Duration {
	return time.Since(item.CreatedAt)
}
