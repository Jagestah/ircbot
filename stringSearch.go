package main

import "fmt"
import "io/ioutil"
//import "strings"
//import "bufio"
//import "os"

func main() {
	file, err := ioutil.ReadFile("points")
		if err != nil {
			panic(err)
		}

	type pointsTable struct {
	name string
	points int
	}
	fileString := string(file)
//	fmt.Printf("%s", (strings.Split(fileString, ",")
	fmt.Println(fileString)

}

