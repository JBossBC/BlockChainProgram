package util

import "strings"

func HandleURLToPath(url string) []string {
	//使用rune防止因为编码集导致出错
	var result = len(url) - 1
	for index, value := range url {
		if value == '?' {
			result = index - 1
			break
		}
	}
	var discussRawQuery = url[:result+1]
	return strings.Split(discussRawQuery, "/")[1:]
}
