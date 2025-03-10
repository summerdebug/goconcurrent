package sqrt

import (
	"fmt"
	"math"
	"sync"
)

// Start square root calculation using caching
func Start() {
	for i := 0; i < 100; i++ {
		fmt.Println(sqrt(i * 1000))
	}
}

func sqrt(i int) float64 {
	var m = cachedMap()
	return m[i]
}

var cachedMap = sync.OnceValue(buildMap)

func buildMap() map[int]float64 {
	m := make(map[int]float64)
	for i := 0; i < 100_000; i++ {
		m[i] = math.Sqrt(float64(i))
	}
	return m
}
