package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var filename string

	fmt.Print("Enter filename to read: ")
	fmt.Scan(&filename)
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list:=[]string{}
	fmt.Println("The file contents are:")
	for i := 1; scanner.Scan(); i++ {
		content := scanner.Text()
		fmt.Printf("%2d. %s\n", i, content)
		list=append(list,content)
	}
	fmt.Println("The longest line:")
	longline,longi:=longest(list)
	fmt.Printf("Line %d:%v\n",longi,longline)
	fmt.Println("The smallest line:")
	smallLine,smalli:=smallest(list)
	fmt.Printf("Line %d:%v\n",smalli,smallLine)
	wc:=countWords(list)
	fmt.Println("Word count of each line:")
	for i,value:=range(wc){
		fmt.Printf("Line %d:%v words\n",i+1,value)
	}
	fmt.Println("Sorted word count:")
	sortWordCount(wc)
	for _,value:=range(wc){
		fmt.Printf("%d words\n",value)
	}
}

func longest(list []string)(string,int){
	max:=0
	var maxi int
	for i,value:=range(list){
		if len(value)>max{
			max=len(value)
			maxi=i
		}
	}
	return list[maxi],maxi
}

func smallest(list []string)(string,int){
	min:=999
	var mini int
	for i,value:=range(list){
		if len(value)<min{
			min=len(value)
			mini=i
		}
	}
	return list[mini],mini
}

func countWords( list [ ] string) [ ] int {
// Implementation to count words in each line
	wc:=[]int{}
	for i:=0;i<len(list);i++{
		words:=strings.Split(list[i]," ")
		count:=len(words)
		wc=append(wc,count)
	}
	return wc
}

func sortWordCount(wordCounts [ ] int) {
	sort.Ints(wordCounts)
	return
}