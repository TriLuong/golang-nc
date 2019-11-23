package main

import (
	"fmt"

	"strconv"

	"reflect"

	"github.com/Triluong/golang-nc/type_system/student"
)

func main() {
	//map
	aMap := map[string]int{"a": 1, "b": 3, "d": 5}

	for i, v := range aMap {
		fmt.Printf("Index at %s has value %d \n", i, v)
	}

	//map[string]interface
	aStudent := student.Student{FirstName: "Tri", LastName: "Luong", Age: 25, Email: "lxt0094@gmail.com"}

	fmt.Printf("%+v", aStudent)
	// var studentMap map[string]interface{ aStudent }
	// fmt.Println(studentMap["FirstName"])

	//channel
	// ch := make(chan int)
	// go func() {
	// 	ch <- 1
	// 	ch <- 2
	// 	ch <- 3
	// 	ch <- 4
	// 	close(ch)
	// }()

	// for v := range ch {
	// 	fmt.Println(v)
	// }
	aString := "hello world"

	bs := []byte(aString)

	bString := string(bs)
	fmt.Println(bString)

	aInt := 100

	aString = strconv.Itoa(aInt)
	bInt, err := strconv.Atoi(aString)
	if err != nil {
		panic(err)
	}
	fmt.Println(bInt + 10)

	// int -> float

	aInt = 300
	aFloat := float32(aInt)
	fmt.Println(aFloat + 10.2)

	type1 := reflect.TypeOf(aFloat)
	fmt.Println(aFloat, " type = ", type1.Name(), " kind = ", type1.Kind())

	cInt := int(aFloat)
	type2 := reflect.TypeOf(cInt)
	fmt.Println(cInt, " type = ", type2.Name(), " kind = ", type2.Kind())

	//struct to struct

	type Boy struct {
		FirstName string
		LastName  string
		Age       int
		Email     string
	}

	boy := Boy{"Anh", "Nguyen", 12, "anhnguyen@gmail.com"}

	type3 := reflect.TypeOf(boy)
	fmt.Println(boy, "type=", type3.Name(), "kind=", type3.Kind())

	bStudent := student.Student(boy)
	type4 := reflect.TypeOf(bStudent)
	fmt.Println(bStudent, "type=", type4.Name(), "kind=", type4.Kind())

	//type conversion vs type assertion
	// type assertion interface{} -> string, int, struct
	var i interface{}

	i = 100
	aInt = i.(int)
	fmt.Println(aInt)
}
