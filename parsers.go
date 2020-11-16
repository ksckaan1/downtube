//ALL PARSER FUNCTIONS
package main

import (
	"fmt"
	"strconv"
)

func viewCountParser(views int) string {
	var result string
	vl := len(strconv.Itoa(views))
	if vl < 4 {
		result = strconv.Itoa(views)
	} else if vl > 3 && vl < 7 {
		result = fmt.Sprintf("%.1F K", (float64(views) / 1000))
	} else if vl > 6 && vl < 10 {
		result = fmt.Sprintf("%.1F M", (float64(views) / 1000000))
	} else {
		result = fmt.Sprintf("%.1F B", (float64(views) / 1000000000))
	}
	return fmt.Sprintf("%s views", result)
}

func fileSizeParser(fileSize int) string {
	var result string
	fl := len(strconv.Itoa(fileSize))
	if fileSize < 4 {
		result = strconv.Itoa(fileSize)
	} else if fl > 3 && fl < 7 {
		result = fmt.Sprintf("%.1F KB", (float64(fileSize) / 1000))
	} else if fl > 6 && fl < 10 {
		result = fmt.Sprintf("%.1F MB", (float64(fileSize) / 1000000))
	} else if fl > 9 && fl < 13 {
		result = fmt.Sprintf("%.1F GB", (float64(fileSize) / 1000000000))
	} else {
		result = fmt.Sprintf("%.1F TB", (float64(fileSize) / 1000000000000))
	}
	return result
}

func durationParser(duration int) string {
	hour := (duration / 60) / 60
	min := (duration / 60) % 60
	sec := duration % 60
	var time string
	if hour == 0 {
		time = fmt.Sprintf("%dm:%ds", min, sec)
	} else if hour == 0 && min == 0 {
		time = fmt.Sprintf("%ds", sec)
	} else {
		time = fmt.Sprintf("%dh:%dm:%ds", hour, min, sec)
	}
	return time
}
