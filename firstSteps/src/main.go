package main

import (
	"fmt"
	"strconv"
	"errors"
	"sync"
)

var wg = sync.WaitGroup{}

func doSomething() string{
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

	var myBool bool = true
	newBool := false
	fmt.Println(!myBool == newBool) // !myBool = ! true = false == false -> true

	const myName string = "Willem"
	fmt.Println(myName)
	return (myName)
}

func additionINT(nombre1 int, nombre2 int) int{
	return nombre1 + nombre2
}

func divisionINT(nombre1 int, nombre2 int) (int, error){
	var err error
	if nombre2 == 0{
		err = errors.New("cannot divide by 0")
		return 0, err
	}
	return nombre1 / nombre2, err
}

func useArray(){
	var mainArray [10]int
	fmt.Println(mainArray[1:5])

	var adresse0 *int = &mainArray[0] // l'adresse de la liste en position 0
	var a int = 17
	*adresse0 = a // On écrit là où pointe l'adresse 17
	fmt.Println(mainArray)
	mainSlice := mainArray[:]
	mainSlice = append(mainSlice, 89)
	fmt.Println(mainSlice, cap(mainSlice), len(mainSlice))
	wg.Done()
}

func useMap(){
	var myMap = map[string]int {"Willem" : 22, "Chloe" : 24}
	for name := range myMap{
		fmt.Println(name + " -> " + strconv.Itoa(myMap[name]))
	}

	for name, age := range myMap{
		fmt.Println(name + " -> " + strconv.Itoa(age))
	}
	wg.Done()
}

func useLoop(){
	for i:= 0; i < 5000; i++{
		fmt.Println("Valeur : " + strconv.Itoa(i))
	}
	wg.Done()
}

func main(){
	var name string = doSomething();
	fmt.Println(name)
	fmt.Println(additionINT(10, 78))
	
	valRetour, err := divisionINT(10, 0)
	if err == nil{
		fmt.Println("La valeur de retour est : " + strconv.Itoa(valRetour))
	} else {
		fmt.Println("Error : " + err.Error())
	}

	go useArray()
	wg.Add(1)
	go useMap()
	wg.Add(1)
	go useLoop()
	wg.Add(1)

	wg.Wait()
}