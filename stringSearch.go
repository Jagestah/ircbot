package main

import "fmt"
import "io/ioutil"
import "strings"
import "bufio"
//import "os"


func main() {

	scanner := bufio.NewScanner(data)

	data, err := ioutil.ReadFile("points")
		if err != nil {
			panic(err)
		}
	file := string(data)
	line := 0
	temp := strings.Split(file, "\n")
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "jagestah") {
			return line, nil
		}
		line++
	
	}

}

//	var text string
//	fmt.Print("Enter text: ")
//	fmt.Scanln(&text)
//
//	b, err := ioutil.ReadFile("points")
//		if err != nil {
//			panic(err)
//		}
//	s := string(b)
//	s = strings.TrimRight(s, ",")
//	fmt.Println(s)
//	a := strings.SplitAfterN(s, ",", 2)
//	fmt.Println(strings.Contains(a, text))
	
//	func readLines("points") ([]string, error) {
//	points, err := os.Open("points")
//		if err != nil {
//			return nil, err
//		}
//	defer points.Close()
//
//	var lines []string
//	scanner := bufio.NewScanner(points)
//	for scanner.Scan() {
//		lines = append(lines, scanner.Text())
//	}
//	return lines, scanner.Err()
//}


