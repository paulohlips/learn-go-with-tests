package main

import "fmt"

func main() {
	fmt.Println(Hello("Paulo"))
}

func Hello(name string) string{
	if name == "" {
		return "Hello, World"
	} else {
		return "Hello, " + name
	}
}