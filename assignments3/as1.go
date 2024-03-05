package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"
)

type Customer struct {
	Id    int
	Name  string
	City  string
	Email string
	phone string
}

type CustomerDAO struct {
	Slice []Customer
}

func (cust *Customer) Read() {
	fmt.Print("Enter Customer ID:")
	fmt.Scan(&cust.Id)
	fmt.Print("Enter Customer Name:")
	fmt.Scan(&cust.Name)
	fmt.Print("Enter Customer City:")
	fmt.Scan(&cust.City)
	fmt.Print("Enter Customer Email:")
	fmt.Scan(&cust.Email)
	fmt.Print("Enter Customer phone:")
	fmt.Scan(&cust.phone)
}
func (cust *Customer) Display() {
	fmt.Println("Name:", cust.Name)
	fmt.Println("Id:", cust.Id)
	fmt.Println("Email:", cust.Email)
	fmt.Println("City:", cust.City)
	fmt.Println("phone:", cust.phone)
	fmt.Println()
}
func (dao CustomerDAO) Display() {
	for i, value := range dao.Slice {
		fmt.Println("Customer", i+1)
		value.Display()
	}
}
func (dao *CustomerDAO) Add(cust Customer) {
	dao.Slice = append(dao.Slice, cust)
}

func (dao CustomerDAO) Get(id int) Customer {
	// returns the Customer based on id
	for _, value := range dao.Slice {
		if value.Id == id {
			return value
		} else {
			panic("ID not found")
		}
	}
	return Customer{}
}

func (dao *CustomerDAO) Update(cust Customer,id int) {
	// updates the [ ] Customer based on id of the customer
	for _, value := range dao.Slice {
		if value.Id == id {
			value = cust
		}
	}
}

func (dao *CustomerDAO) Delete(id int) Customer {
	// returns the deleted Customer
	for i, value := range dao.Slice {
		if value.Id == id {
			dao.Slice = append(dao.Slice[:i], dao.Slice[i+1:]...)
			return value
		} else {
			panic("ID not found")
		}
	}
	return Customer{}
}

func (dao CustomerDAO) GetAll() []Customer {
	// returns all Customer records
	return dao.Slice
}

func (dao CustomerDAO) GetByCity(city string) []Customer {
	// returns all Customer records with the given city
	cust := []Customer{}
	for _, value := range dao.Slice {
		if value.City == city {
			cust = append(cust, value)
		}
	}
	return cust
}

func (dao CustomerDAO) SaveAsXML() {
	// saves all customer records into a file named `customers_xxxxx.xml'
	// where 'xxxxx* is the current timestamp
	timestamp := time.Now().Unix()
		filename := "customers_" + string(timestamp) + ".xml"
		file, err := os.Create(filename)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
	for _, value := range dao.Slice {
		
		if bytes, err := xml.MarshalIndent(value, "", "    "); err != nil {
			fmt.Println("There was an error:", err)
		} else {
			file.WriteString(string(bytes))
		}
	}
}

func (dao CustomerDAO) SaveAsJSON() {
	// saves all customer records into a file named `customers_xxxxx.xml'
	// where 'xxxxx* is the current timestamp
	timestamp := time.Now().Unix()
		filename := "customers_" + string(timestamp) + ".json"
		file, err := os.Create(filename)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
	for _, value := range dao.Slice {
		if bytes, err := json.MarshalIndent(value, "", "    "); err != nil {
			fmt.Println("There was an error:", err)
		} else {
			file.WriteString(string(bytes))
		}
	}
}

func (dao CustomerDAO) LoadFromFile(filename, filetype string) {
	// loads data from a JSON or XML file into the []Customer
	// maintained by this DAO class
	var cust Customer
	if filetype == "XML" || filetype == "xml" {
		file, err := os.Open(filename)
		if err != nil {
			panic(err.Error())
		}

		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err.Error())
		}

		if err := json.Unmarshal(bytes, &cust); err != nil {
			panic(err)
		}

		dao.Slice = append(dao.Slice, cust)
	} else {
		file, err := os.Open(filename)
		if err != nil {
			panic(err.Error())
		}

		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err.Error())
		}

		if err := xml.Unmarshal(bytes, &cust); err != nil {
			panic(err)
		}

		dao.Slice = append(dao.Slice, cust)
	}
}

func main() {
	fmt.Println("Main menu:")
	fmt.Println("1. Load data from JSON/XML file\n2. Add new customer\n3. View all customers\n4. Search customers by city\n5. Delete a customer (by id)\n6. Search customer by id and edit/update details\n7. Save to an XML file\n8. Save to a JSON file\n9. Exit")
	var choice int
	var dao CustomerDAO
	for {
		fmt.Print("\nEnter Choice:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var filename, filetype string
			fmt.Print("Enter the filename and filetype:")
			fmt.Scan(&filename, &filetype)
			dao.LoadFromFile(filename, filetype)
		case 2:
			var cust Customer
			cust.Read()
			dao.Add(cust)
		case 3:
			dao.Display()
		case 4:
			var city string
			fmt.Print("Enter the city:")
			fmt.Scan(&city)
			cust := []Customer{}
			cust = dao.GetByCity(city)
			for i, value := range cust {
				fmt.Println("Customer", i+1)
				value.Display()
			}
		case 5:
			var id int
			fmt.Print("Enter the id of the customer to be deleted:")
			fmt.Scan(&id)
			cust:=dao.Delete(id)
			fmt.Println("Details of the customer deleted:")
			cust.Display()
		case 6:
			var cust Customer
			var id int
			fmt.Print("Enter the id of the customer to be updated:")
			fmt.Scan(&id)
			fmt.Println("Enter new values of the customer:")
			cust.Read()
			dao.Update(cust,id)
			fmt.Print("Customer data Updated!")
		case 8:
			dao.SaveAsXML()
		case 9:
			dao.SaveAsJSON()
		default:
			fmt.Print("Invalid Choice")
		}
	}
}
