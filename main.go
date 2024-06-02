package main

import (
	"fmt"
	// "math"
	// "sort"
	// "strings"
)

// using a function outside of a function
// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }
// func sayBye(n string) {
// 	fmt.Printf("Good Bye %v \n", n)
// }
// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }



func main() {

	// fmt.Println("Hello guys, testing go!")

	// Strings

	// strings in go are double quote, using a single quote will throw an error
	// var nameOne string = "bright 1"
	// it reading the initial value of a variable if both lines look alike
	// var nameTwo = "bright 2"
	// setting a variable without a value
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// when declaring a var its not necesary to make use of the var key word we use := note it is used only the first time you are declaring a variable
	// it cannot be used outside of a function note!
	// nameFour := "bright 4"

	// fmt.Println(nameFour)

	// Integers
	// here are several methods of an integar
	// var ageOne int = 20

	// var ageTwo = 10

	// ageThree := 30

	// fmt.Println(ageOne, ageThree, ageTwo)

	// Bits & memory (check docs)
	// var numOne int8 = 25
	// using unsigned int does not allow minus
	// var numThree uint8 = 213

	// FLoat (default float is float64)
	// var scoreOne float32 = -44432342.3
	// var scoreTwo float64 = 3244432342.3

	// Printing and formatting strings
	// fmt.Print("hello,")
	// fmt.Print("world! \n")
	// adding new line \n
	// fmt.Print("new line")

	// outputing variables
	// age := 43
	// name := "Gbemi"

	// fmt.Println("my name is", name, "and i am", age, "years old")

	// Formatting string: embedding different varuable inside it
	// fmt.Printf("my name is %q and my age is %v", name, age)

	// Array and Slices
	// number put inside the square box is the  number of items that should be int the array
	// var ages [6]int = [6]int{21, 12, 32, 23, 44, 65}

	// shorthand
	// var ages = [6]int{21, 12, 32, 23, 44, 65}

	// names := [3]string{"best", "tobi", "gbemi"}

	// To change the value an item i array
	// names[1] = "bright"

	// fmt.Println(ages, len(ages), "\n", names, len(names))

	// Slices (uses array under the hood) they are arrays you can be able to edit the length
	// var score = []int{12, 234, 343, 54}

	// score[2] = 32

	// append a slice
	// score = append(score, 85)

	// fmt.Println(score)

	// slice ranges - used to get a range of element in a slice or array
	// from element 0 - 3
	// rangeOne := names[0:3]
	// from element 1 - the last one
	// rangeTwo := names[1:]
	// from element the fitst one - 3
	// rangeThree := names[:3]

	// rangeOne = append(rangeOne, "yinka")

	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	//  Standard Library

	// String Package

	// greeting := "hey guys!"
	// this is used for a search it the sting contains what the variable defines
	// fmt.Println(strings.Contains(greeting, "hey"))
	// uses to replace one term with another
	// fmt.Println(strings.ReplaceAll(greeting, "hey", "Hello"))
	// to uppercase
	// fmt.Println(strings.ToUpper(greeting))
	// getting the index of a term
	// fmt.Println(strings.Index(greeting, "gu"))
	// splits the character of a term into an array
	// fmt.Println(strings.Split(greeting, " "))

	// Sort Package
	// ages := []int{32, 43, 65, 7, 87, 31, 55, 88}

	// sort.Ints(ages)
	// fmt.Println(ages)

	// search a slice of integers
	// index := sort.SearchInts(ages, 7)
	// fmt.Println(index)

	// names := []string{"best", "tobi", "gbemi"}
	// sort.Strings(names)
	// fmt.Println(names)
	// fmt.Println(sort.SearchStrings(names, "tobi"))

	// Loops
	// method one
	// x := 0
	// for x <= 5 {
	// 	fmt.Println("this is the value of x", x)
	// 	x++
	// }

	// method two
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("this is i", i)
	// }

	// for strings
	// names := []string{"bright", "silas", "gbemi", "tobi", "best"}

	// for i := 0; i < len(names); i++ {
	// 	sort.Strings(names)
	// 	fmt.Println(names[i])
	// }

	// four in loop - range
	// for index, value := range names {
	// 	fmt.Printf("the postionn at index %v is %v \n", index, value)
	// }

	// Boolean and conditions
	// age := 54

	// fmt.Println(age <= 344)
	// fmt.Println(age >= 344)
	// fmt.Println(age == 54)
	// fmt.Println(age != 60)

	// if age < 30 {
	// 	fmt.Println("age is greater than 30")
	// } else if age == 30 {
	// 	fmt.Println("age is less than 30")
	// 	} else {
	// 	fmt.Println("age is", age)
	// }

	// names := []string{"bright", "silas", "gbemi", "tobi", "best"}

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continueing at pos", index)
	// 		continue
	// 	}
	// 	if index > 3 {
	// 		fmt.Println("break at pos", index)
	// 		break
	// 	}
	// 	fmt.Printf("the value at pos %v is %v \n", index, value)
	// }

	// implementing functions
	// cycleNames([]string{"bright", "silas", "gbemi", "tobi", "best"}, sayGreeting)
	// cycleNames([]string{"bright", "silas", "gbemi", "tobi", "best"}, sayBye)

	// a1 := circleArea(3)
	// a2 := circleArea(10)
	// a3 := circleArea(32.4)

	// fmt.Println(a1, a2, a3)
}
