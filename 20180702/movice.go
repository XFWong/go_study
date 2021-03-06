package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Movice struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}
	var movices = []Movice{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	//1.未格式化输出
	// data, err := json.Marshal(movices)
	// if err != nil {
	// 	fmt.Printf("err code:%s", err)
	// }
	// fmt.Printf("%s\n", data)
	// 2.格式化输出
	data, err := json.MarshalIndent(movices, "", "    ")
	if err != nil {
		fmt.Printf("errcode %d", err)
	}
	fmt.Printf("%s\n", data)
}
