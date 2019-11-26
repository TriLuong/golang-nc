package main

import (
	"fmt"
	"github.com/Triluong/golang-nc/type_system/student"
	"reflect"
)

func main() {
	aStruct := struct {
		Name string
		Age  int
	}{Name: "Thang", Age: 35}

	aType := reflect.TypeOf(aStruct)
	fmt.Printf("type %s, kind %s \n", aType.Name(), aType.Kind())

	bStruct := student.Student{FirstName: "A", LastName: "B"}
	bType := reflect.TypeOf(bStruct)
	fmt.Printf("type %s, kind %s\n", bType.Name(), bType.Kind())

	fmt.Println(bType.NumField())

}
