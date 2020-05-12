package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Math struct which contains name, type, age and Social struct
type Math struct {
	Operation string `json:"Operation"`
	Units     Units  `json:"Units"`
}

// Units struct which contains list of links
type Units struct {
	One int `json:"one"`
	Two int `json:"two"`
}

//NewStruct struct which contains firstname and lastname
type NewStruct struct {
	Addition    string `json:"Addition"`
	Subtraction string `json:"Subtraction"`
	Mutliply    string `json:"Mutliply"`
	Division    string `json:"Divison"`
}

//Person struct which contains firstname and lastname
type Person struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address,omitempty"`
}

//Address struct which contains city and state
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	// jsonFile, err := os.Open("users.json")
	Read("assignment.txt", "txt")
	Read("humans.csv", "csv")
	Read("users.json", "json")

}

//Read file based on file Type
func Read(filename string, fileType string) {

	// Open file
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nSuccessfully Opened file")
	defer f.Close()

	// Read File into a Variable
	if fileType == "csv" {
		info := csv.NewReader(bufio.NewReader(f))
		var people []Person
		for {
			data, err := info.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			people = append(people, Person{
				Firstname: data[0],
				Lastname:  data[1],
				Address: &Address{
					City:  data[2],
					State: data[3],
				},
			})

		}
		peopleJson, _ := json.Marshal(people)
		fmt.Println(string(peopleJson))

	} else if fileType == "json" {
		data, _ := ioutil.ReadAll(f)
		var operations []Math

		json.Unmarshal(data, &operations)
		var myStruct NewStruct

		for i := 0; i < len(operations); i++ {
			if operations[i].Operation == "Add" {
				var sum = operations[i].Units.One + operations[i].Units.Two
				fmt.Println(sum)
				myStruct.Addition = strconv.Itoa(sum)
			} else if operations[i].Operation == "Subtract" {
				var sub = operations[i].Units.One - operations[i].Units.Two
				fmt.Println(sub)
				myStruct.Addition = strconv.Itoa(sub)
			} else if operations[i].Operation == "Multiply" {
				var multi = operations[i].Units.One * operations[i].Units.Two
				fmt.Println(multi)
				myStruct.Addition = strconv.Itoa(multi)
			} else if operations[i].Operation == "Divide" {
				var div = operations[i].Units.One / operations[i].Units.Two
				fmt.Println(div)
				myStruct.Addition = strconv.Itoa(div)
			}
		}
		//Tried to append values to existing json file but didn't work
		result, _ := json.Marshal(myStruct)
		new, _ := io.WriteString(f, string(result))
		fmt.Println(new)
		// till here
	} else if fileType == "txt" {
		data, _ := ioutil.ReadFile(filename)
		fmt.Print(string(data))
	}
}
