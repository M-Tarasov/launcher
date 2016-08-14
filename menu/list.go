package menu

import (
	"os"
	"log"
	"encoding/json"
)

type ListItem struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Cmd  []string `json:"cmd"`
}

type List struct {
	List []ListItem
}

func  Load(file string) *List {
	fin, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}

	res := new(List)

	decoder := json.NewDecoder(fin)

	if err := decoder.Decode(&res.List); err != nil {
		log.Fatalln(err)
	}

	return res
}

