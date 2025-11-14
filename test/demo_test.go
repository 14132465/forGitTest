package main

import (
	"fmt"
	"testing"
)

func TestFive(t *testing.T) {
	i := 1
	i_p := &i
	defer fmt.Println(i)
	defer func() {
		xxx := *i_p
		fmt.Println(xxx)
	}()
	i++
	panic("panic")
	*i_p++
}
