package main

import "github.com/dankore/greetings/greetings"
import "fmt"

func main(){
	message := greetings.Hello("Adamu")

	fmt.Println(message)
}