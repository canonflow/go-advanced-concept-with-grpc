package main

import (
	"fmt"
	"reflect"
)

// ============== METHODS ==============
type Greeter struct{}

func (g Greeter) Greet(fname, lname string) string {
	return "Hello " + fname + " " + lname
}

func main() {
	greeter := Greeter{}
	t := reflect.TypeOf(greeter)
	v := reflect.ValueOf(greeter)
	var method reflect.Method
	fmt.Println("Type:", t)

	for i := range t.NumMethod() {
		method = t.Method(i)
		fmt.Printf("Method %d: %v\n", i, method.Name)
	}
	/*
		Type: main.Greeter
		Method 0: Greet
	*/

	m := v.MethodByName(method.Name)
	results := m.Call([]reflect.Value{
		reflect.ValueOf("Alice"),
		reflect.ValueOf("Doe"),
	})
	fmt.Println("Greet Result:", results[0].String()) // Greet Result: Hello Alice Doe
}

// ============== STRUCT ==============
// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	person := Person{
// 		Name: "Alice",
// 		Age:  30,
// 	}

// 	v := reflect.ValueOf(person)

// 	for i := range v.NumField() {
// 		fmt.Printf("Field %d: %v\n", i, v.Field(i))
// 	}
// 	/*
// 		Field 0: Alice
// 		Field 1: 30
// 	*/

// 	v1 := reflect.ValueOf(&person).Elem()
// 	nameField := v1.FieldByName("Name")
// 	if nameField.CanSet() {
// 		nameField.SetString("Jane")
// 	} else {
// 		fmt.Println("Cannot set")
// 	}
// 	fmt.Println("Modified Person:", person) // Modified Person: {Jane 30}
// }

// ============== BASIC ==============
// func main() {
// 	x := 42
// 	v := reflect.ValueOf(x)

// 	t := v.Type()

// 	fmt.Println("Value:", v)
// 	fmt.Println("Type:", t)
// 	fmt.Println("Kind:", t.Kind())
// 	fmt.Println("Is Int:", t.Kind() == reflect.Int)
// 	fmt.Println("Is String:", t.Kind() == reflect.String)
// 	fmt.Println("Is Zero:", v.IsZero())
// 	/*
// 		Value: 42
// 		Type: int
// 		Kind: int
// 		Is Int: true
// 		Is String: false
// 		Is Zero: false
// 	*/

// 	y := 10
// 	v1 := reflect.ValueOf(&y).Elem()
// 	v2 := reflect.ValueOf(&y)

// 	fmt.Println("V2 Type:", v2.Type())
// 	fmt.Println("Original Value:", v1.Int())
// 	v1.SetInt(18)
// 	fmt.Println("Modified Value:", v1.Int())
// 	/*
// 		V2 Type: *int
// 		Original Value: 10
// 		Modified Value: 18
// 	*/

// 	var itf interface{} = "Hello"
// 	v3 := reflect.ValueOf(itf)

// 	fmt.Println("V3 Type:", v3.Type())
// 	if v3.Kind() == reflect.String {
// 		fmt.Println("String value:", v3.String())
// 	}
// 	/*
// 		V3 Type: string
// 		String value: Hello
// 	*/
// }
