package main

import "fmt"

func main() {
	fmt.Println(Hello("Paulo"))
}

func Hello(name string) string{
	return "Hello, " + name
}