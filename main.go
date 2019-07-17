package main

import (
	"fmt"
)

func main(){
	fmt.Println(greeting("1. Ed", "Dev"))

	fmt.Println(greetingNamedFunc("2. Ed", "Dev"))

	fmt.Println(greetingMultipleValuesFunc("3. Ed", "Dev"))
}

func greeting(firstName, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName)
} 

func greetingNamedFunc(firstName, lastName string) (s string) {
	s = fmt.Sprintf("%s %s", firstName, lastName)
	return
} 

func greetingMultipleValuesFunc(firstName, lastName string) (string, string) {
	return fmt.Sprintf("%s %s", firstName, lastName), "!!!"
} 