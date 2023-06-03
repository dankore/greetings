package main

import "github.com/dankore/handy-go-packages/greetings"
import "fmt"

func main(){
	message := greetings.Hello("Adamu")

	fmt.Println(message)
}