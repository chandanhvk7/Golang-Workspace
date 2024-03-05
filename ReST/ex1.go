
package main

import ("fmt"
"net/http")

func home(resp http.ResponseWriter,req *http.Request){
	fmt.Fprintf(resp,"<html><h1>Hello World</h1></html>")
}

func main(){
	fmt.Println("Starting the http server at port 7777")
	http.HandleFunc("/",home)
	http.ListenAndServe("0.0.0.0:7777",nil)
}