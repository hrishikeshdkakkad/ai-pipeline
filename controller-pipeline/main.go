package main

import (
	"fmt"
)

func main() {
	normalBuilder := getBuilder("normal")

	director := newDirector(normalBuilder)
	res := director.buildResponse()
	fmt.Println(res)
}
