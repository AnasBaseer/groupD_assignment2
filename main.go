package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	result := add(3, 7)
	fmt.Println("Sum:", result)
	
	var year int
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)
	if isLeapYear(year) {
	fmt.Printf("%d is a leap year.\n", year)
		}
 	else {
	fmt.Printf("%d is not a leap year.\n", year)
		}
}
func add(a, b int) int {
	return a + b
}
func revString(str string) string {
	var theArray []string
	theArray = strings.Split(str, "")
	var strOutput string
	for i := len(str) - 1; i >= 0; i-- {
		strOutput += theArray[i]
	}
	return strOutput
}
func isLeapYear(year int) bool {
if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
return true
}
return false
}
