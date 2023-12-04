package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.
		NewVerifier().
		EnableSMTPCheck()
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an email as an argument.")
		return
	}

	email := os.Args[1]
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		fmt.Println("Invalid email format.")
		return
	}

	username := parts[0]
	domain := parts[1]

	ret, err := verifier.CheckSMTP(domain, username)
	if err != nil {
		fmt.Println("check smtp failed: ", err)
		return
	}

	val := reflect.ValueOf(ret).Elem()
	typeOfT := val.Type()

	deliverable := false
	for i := 0; i < val.NumField(); i++ {
		fieldName := typeOfT.Field(i).Name
		fieldValue := val.Field(i).Interface()
		fmt.Printf("Field: %s\tValue: %v\n", fieldName, fieldValue)
		if fieldName == "Deliverable" && fieldValue == true {
			deliverable = true
		}
	}

	resultsDir := filepath.Join("..", "RESULTS")
	if _, err := os.Stat(resultsDir); os.IsNotExist(err) {
		os.Mkdir(resultsDir, 0755)
	}

	var fileName string
	if deliverable {
		fileName = "valid.txt"
	} else {
		fileName = "invalid.txt"
	}

	filePath := filepath.Join(resultsDir, fileName)
	ioutil.WriteFile(filePath, []byte(email+"\n"), 0644)
}
