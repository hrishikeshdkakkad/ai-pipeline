package main

import (
	"fmt"
)

func main() {
	normalBuilder := getBuilder("normal")

	director := newDirector(normalBuilder)
	res := director.buildResponse()
	//time.Sleep(12 * time.Second)
	fmt.Println(res)
}
