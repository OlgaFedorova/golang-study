package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type MyStruct struct {
	GlobalID int `json:"global_id"`
}

func main() {
	//data, err := ioutil.ReadAll(os.Stdin)
	file, err := os.Open("stepic-golang/chapter3/chapter3_6_json/task_3_6_2/data-20190514T0100.json")
	data, err := ioutil.ReadAll(file)
	if err == nil {
		md := []MyStruct{}
		err = json.Unmarshal(data, &md)
		if err == nil {
			var sum int
			for _, s := range md {
				sum = sum + s.GlobalID
			}
			fmt.Println(sum)
		}
	}
}
