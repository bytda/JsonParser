package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type info struct {
	Cars []struct {
		Name string
		Age  int
	}
}

func main() {
	//var text string = `{"cars":[{"name":"Bob","age":15}, {"name":"Max","age":20}, {"name":"German","age":22}]}`

	path := "/"
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var infos info
	//json.Unmarshal([]byte(text), &infos)
	json.Unmarshal(byteValue, &infos)
	infos.Cars[1].Name = "Cris"
	//var tom = info {Cars: [Name:"Kemran", Age:42]}
	//fmt.Println(infos)
	//fmt.Println(infos.Cars[2])
	for i := range infos.Cars {
		fmt.Println(i, infos.Cars[i].Name, infos.Cars[i].Age)
	}
}
