package try

import (
    "fmt"
	"os"
	"io"
	"encoding/json"
)

type user struct {
	Value int `json:"Score"`
	Name  string `json:"Nom"`
}

type users struct {
	Users []user `json:"Users"`
}


func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return byteValue, nil
}

func ParseUsers(byteValue []byte) users {
	var listUsers users
	err := json.Unmarshal(byteValue, &listUsers)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return listUsers
	}

	return listUsers
}