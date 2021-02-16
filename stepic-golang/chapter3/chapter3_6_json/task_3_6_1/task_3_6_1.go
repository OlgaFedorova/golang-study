package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type myData struct {
	Average float64
}

func main() {
	//data, err := ioutil.ReadAll(os.Stdin)
	file, err := os.Open("stepic-golang/chapter3/chapter3_6_json/task_3_6_1/text.json")
	data, err := ioutil.ReadAll(file)
	if err == nil {
		group := Group{}
		err = json.Unmarshal(data, &group)
		if err == nil {
			var count int
			for _, s := range group.Students {
				count = count + len(s.Rating)
			}

			result := myData{
				Average: float64(count) / float64(len(group.Students)),
			}

			dr, err := json.MarshalIndent(result, "", "    ")
			if err == nil {
				fmt.Printf("%s", dr)
			}
		}
	}
}
