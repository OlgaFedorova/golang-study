package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
)

func main() {
	reader, err := zip.OpenReader("stepic-golang/chapter3/chapter3_5_files/task3_5_2/task.zip")
	if err == nil {
		defer reader.Close()
		for _, file := range reader.File {
			if !file.FileInfo().IsDir() {
				fc, err := file.Open()
				if err == nil {
					defer fc.Close()
					content, err := ioutil.ReadAll(fc)
					if err == nil {
						buf := bytes.NewBuffer(content)
						cr := csv.NewReader(buf)
						data, err := cr.ReadAll()
						if err == nil && len(data) > 1 {
							fmt.Println(file.Name)
							fmt.Println(data[4][2])
						}
					}
				}
			}
		}
	}

}
