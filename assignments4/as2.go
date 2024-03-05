package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)
func handleURL(url string){
	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}


	filename:=url[strings.LastIndex("/",url)+1:]
	file, err := os.Create(filename)
	fmt.Println("Downloading file ",filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Fprintln(file,string(data))
	fmt.Printf("File %v downloaded!\n",filename)
	wg.Done()
}
var wg sync.WaitGroup
func main() {

	urls:=[]string{"https://repo1.maven.org/maven2/com/mysql/mysql-connector-j/8.3.0/mysql-connector-j-8.3.0.jar"}

	var num int
	fmt.Print("Enter the number of urls:")
	fmt.Scan(&num)
	fmt.Print("Enter the urls separated by space:")
	for num > 0 {
		var url string
		fmt.Scan(&url)
		wg.Add(1)
		go handleURL(url)
		num--
	}
	wg.Wait()
}
