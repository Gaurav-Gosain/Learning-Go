package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"log"
	"reflect"
)

var print = fmt.Println

func main() {
	print("Hello", stuff.Name)
	intArr := []int{1, 2, 3}
	print(reflect.TypeOf(intArr))
	strArr := stuff.IntArrToStrArr(intArr)
	print(strArr)
	print(reflect.TypeOf(strArr))

	date := stuff.Date{}
	err := date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(1974)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d/%d/%d\n", date.GetDay(), date.GetMonth(), date.GetYear())
	print(date.GetDate())
}
