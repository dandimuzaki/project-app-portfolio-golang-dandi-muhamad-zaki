package utils

import (
	"math"
	"strings"
)

func StrToSlice(str string) []string {
	var slice []string
	slice = strings.Split(str, ", ")
	return slice
}

func TotalPage(limit int, totalData int) int {
	if totalData <= 0 {
		return 0
	}

	flimit := float64(limit)
	fdata := float64(totalData)

	res := math.Ceil(fdata / flimit)

	return int(res)
}
