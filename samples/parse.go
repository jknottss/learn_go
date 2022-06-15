package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	type myStruct struct {
		Id int `json:"global_id"`
	}

	var unmsh []myStruct
	res := 0
	file, _ := os.Open("big.json")
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(data, &unmsh)
	for _, nbr := range unmsh {
		res = nbr.Id + res
	}
	fmt.Println(res)
}
