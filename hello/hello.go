package main

import "github.com/dankore/handy-go-packages/greetings"
import "fmt"
import "log"

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	var names = []string {"Adamu", "Fatima", "Adam", "Kado", "Zahra"}

	message, err:= greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}