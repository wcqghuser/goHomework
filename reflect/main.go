package main

import (
	"reflect"
	"fmt"
)

type Person struct {
	Name string
	Age  uint
}
type Student struct {
	Person Person
	Score  uint
}

func (stu *Student) SetScore(score uint) bool {
	stu.Score = score
	return true
}

func reflectAllFields(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			if v.Field(i).Type().Kind() == reflect.Struct {
				reflectAllFields(v.Field(i).Interface())
			} else {
				field := t.Field(i)
				fmt.Println(field.Name, "type:", field.Type)
				fmt.Println(field.Name, "value:", v.Field(i))
			}
		}
	}
}

func main() {
	stu := Student{Person{"wcq", 18}, 99}
	fmt.Println(stu)

	v := reflect.ValueOf(&stu)
	v = v.Elem()
	field := v.FieldByName("Score")
	field.SetUint(100)
	fmt.Println(stu)

	t := reflect.TypeOf(stu)
	fmt.Println("Class Name is", t.Name())

	reflectAllFields(stu)
}

