package main

import (
	"fmt"
	"strconv"
)

func doSomething(){
	var age string
	age = "13"
	fmt.Println(age + " ans")


	var ageNombre int
	var err error
	ageNombre, err = strconv.Atoi(age)
	if err != nil {
		fmt.Println("Erreur de conversion")
	} else {
		fmt.Println(ageNombre)
	}
}

func main(){
	doSomething();
}