package main

import (
	"goconcurrent/sqrt"
	"goconcurrent/threegoroutines"
	"goconcurrent/twogoroutines"
)

func main() {
	threegoroutines.Start()
	twogoroutines.Start()
	sqrt.Start()
}
