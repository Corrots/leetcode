package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomData(n, rangeL, rangeR int) []int {
	var res []int
	if rangeL > rangeR {
		return res
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		// 控制随机数的生成范围[rangeL, rangeR]
		res = append(res, rand.Int()%(rangeR-rangeL+1)+rangeL)
	}
	return res
}
