package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func fibonacciArray(n int) {
	n1 := 0
	n2 := 1
	var n3 int
	resutl := make([]int, 0, n)
	resutl = append(resutl, 0, 1)

	for i := 2; i < n; i++ {
		n3 = n1 + n2
		resutl = append(resutl, n3)
		n1 = n2
		n2 = n3
	}
	fmt.Println(resutl)

}
func WriteFileByIO() {
	//Bước 1: tạo file
	file, err := os.Create("write_by_ioutil.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Bước 2: ghi file
	message := []byte("Heloo cac ban")
	err = ioutil.WriteFile("write_by_ioutil.txt", message, 0644)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	//Bước 3: đóng fileText
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	fibonacciArray(10)

	// Open our jsonFile
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	// resutl := make(map[string]string)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)

		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	WriteFileByIO()
}
