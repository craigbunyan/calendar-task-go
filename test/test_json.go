package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Meeting struct {
	Meeting_name string `json:"meeting_name"`
	Date         string `json:"date"`
	Time         string `json:"start_time"`
	Duration     int    `json:"duration"`
}

func main() {

	json_file, err := ioutil.ReadFile("sample_data.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened JSON file")
	}

	data := Meeting{}

	unmarshal_err := json.Unmarshal([]byte(json_file), &data)
	fmt.Println(unmarshal_err)

	fmt.Println(data)
}
