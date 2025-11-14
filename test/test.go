package main

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	name string `json:name-field`
	age  int
}

func main() {

	var resChan = make(chan int)
	// do request

	select {
	case data := <-resChan:
		fmt.Println("data=", data)
		//doData(data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}

}

func getStructTag(f reflect.StructField) string {
	return string(f.Tag)
}
