//Displaying the struct in JSON and XML format
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Address struct {
	Street string `json:"street" xml:"street"`
	Area   string `json:"locality" xml:"locality"`
	City   string `json:"city" xml:"city"`
	State  string `json:"state" xml:"state"`
}

type Customer struct {
	Id      int     `json:"customerId" xml:"customerId"`
	Name    string  `json:"customerName" xml:"customerName"`
	Email   string  `json:"emailAddress" xml:"emailAddress"`
	Address Address `json:"address" xml:"address"`
}

func main() {

	c1 := Customer{Id: 1, Name: "Chandan", Email: "chand@n.com"}
	c1.Address = Address{"1st cross, 1st main", "ISRO layout", "Bangalore", "Karnataka"}
	// c1 := Customer{1, "Vinod", "vinod@vinod.co"}

	if bytes, err := json.MarshalIndent(c1, "", "    "); err != nil {
		// if bytes, err := json.Marshal(c1); err != nil {
		fmt.Println("There was an error:", err)
	} else {
		fmt.Println(string(bytes))
	}

	if bytes, err := xml.MarshalIndent(c1, "", "    "); err != nil {
		fmt.Println("There was an error:", err)
	} else {
		fmt.Println(string(bytes))
	}

}