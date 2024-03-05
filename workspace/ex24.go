//JSON Parsing
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Movie struct {
	Src  string
	Name string
	HOffset   int
	VOffset   int
	Alignment string
}

func (m *Movie) Print() {
	fmt.Printf("Title: %v\n", m.Src)
	fmt.Printf("Year: %v\n", m.Name)
	fmt.Printf("Type: %v\n", m.HOffset)
	fmt.Printf("imdbID: %v\n", m.VOffset)
	fmt.Printf("Poster: %v\n", m.Alignment)
}

func main() {
	file, err := os.Open("hehe.json")

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}

	var movie Movie

	if err := json.Unmarshal(bytes, &movie); err != nil {
		panic(err)
	}

	movie.Print()
}