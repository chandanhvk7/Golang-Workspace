package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Directory path
	dir := "C:\\Dummy"

	// Read directory contents
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic("There was problem opening the file!")
	}
	wordMap := make(map[string]int)
	for _, filename := range files {
		filepath := dir + "\\" + filename.Name()
		file, err := os.Open(filepath)
		var mx sync.Mutex
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		content, _ := io.ReadAll(file)
		wg.Add(1)
		go func(s []byte) {
			defer wg.Done()
			parts := strings.Split(string(s), " ")
			mx.Lock()
			for _, word := range parts {
				fmt.Println("Got word", word)
				wordMap[word]++
			}
			mx.Unlock()

		}(content)
		wg.Wait()
	}
	fmt.Println("Word count:")
	for key, value := range wordMap {
		fmt.Printf("%-12s: %d\n", key, value)
	}
}

/*func wordFreq(words []string){
	 wordMap := make(map[string]int)
	 for _,value:=range words{
		wordMap[value]++
	 }
	fmt.Println("Word count:")
    for key, value := range wordMap {
        fmt.Printf("%-12s: %d\n", key, value)
    }
}*/
