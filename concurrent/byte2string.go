package main

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func main() {
	/***************************************/
	byteArray1 := []byte{'J', 'O', 'H', 'N'}
	fmt.Println("Byte:", byteArray1)

	str1 := BytesToString(byteArray1)
	fmt.Println("String:", str1)

	/****************************************/

	str2 := string(byteArray1[:])
	fmt.Println("String:", str2)

	/****************************************/
	str3 := bytes.NewBuffer(byteArray1).String()
	fmt.Println("String:", str3)
}
