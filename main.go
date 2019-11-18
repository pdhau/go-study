package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
	var result map[string]interface{}
	json.Unmarshal(jsonData, &result)
	datas := result["data"].([]interface{})
	fmt.Print("code -> ")
	fmt.Println(result["code"])
	fmt.Println("message -> " + result["message"].(string))
	for value := range datas {
		fmt.Println("-----------")
		fmt.Println("index -> " + strconv.Itoa(value))
		item := datas[value].(map[string]interface{})
		fmt.Print("id -> ")
		fmt.Println(item["id"])
		fmt.Println("name -> " + item["name"].(string))
	}
}
