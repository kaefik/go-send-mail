package main

import (
	"fmt"
	"testing"
)

func TestgetExtInFileName(t *testing.T) {
	str := "время 00-00 до 8-59 -  - 2017-3-23 по 2017-3-23 - лог звонков.xlsx"
	fmt.Println(getExtInFileName(str))
	//t.Error("Expected 1.5, got ", v)
}
