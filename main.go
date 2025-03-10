package main

import (
	"goconcurrent/threegoroutines"
	"goconcurrent/twogoroutines"
)

func main() {
	threegoroutines.Start()
	twogoroutines.Start()
}
