// 2nd iteration: Convert data to orc
// Input is assumed to be a 2D array of type string

// TODO (for 3rd iteration)
// Create a function which can determine array size of an unknown imported array
// The code should be less hard-coded (i.e., improve nested for loop at 49-61)
// Q: Is the array going to be 2D? Otherwise, edit/improve it
//

package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/scritchley/orc"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Converts any type of input to string
func convert_to_string(s any) string {
	obtype := reflect.TypeOf(s)
	fmt.Println("obtype: ", obtype)
	b := reflect.TypeOf(s).Elem().Name()
	fmt.Println("a: ", b)
	return b
}

// Input data
func input(x int) ([]string, []int) {
	ERROR_MSG := []string{"XXXX"}
	ERROR_NUM := []int{0, 0, 0, 0}
	if x > 0 && x <= 4 {
		if x == 1 {
			Name := []string{"Rose", "Smith", "William", "James", "Rolf"}
			err := []int{0}
			return Name, err
		}
		if x == 2 {
			err := []string{"0"}
			Age := []int{28, 24, 29, 31, 21}
			return err, Age
		}
		if x == 3 {
			Country := []string{"U.K.", "U.S.", "France", "Norway", "Denmark"}
			err := []int{0}
			return Country, err
		}
		if x == 4 {
			Skills := []string{"C", "Java", "PHP", "Python", "C++"}
			err := []int{0}
			return Skills, err
		}
	}
	return ERROR_MSG, ERROR_NUM
}

func test() {
	Name, _ := input(1)
	_, Age := input(2)
	Country, _ := input(3)
	Skills, _ := input(4)

	var1 := "string1" + convert_to_string(Name)
	fmt.Println("1-----")
	var2 := "int1" + convert_to_string(Age)
	fmt.Println("2-----")
	var3 := "string3" + convert_to_string(Country)
	fmt.Println("3-----")
	var4 := "string4" + convert_to_string(Skills)
	fmt.Println("4-----")

	variables := var1 + var2 + var3 + var4
	fmt.Println(variables)

}

func writer() {
	// Data
	Name, _ := input(1)
	_, Age := input(2)
	Country, _ := input(3)
	Skills, _ := input(4)

	// Create .orc file
	f, err := os.Create("employee.orc")
	check(err)

	// set ParseSchema to the following:
	var1 := "Name:" + convert_to_string(Name) + ","
	var2 := "Age:" + convert_to_string(Age) + ","
	var3 := "Country:" + convert_to_string(Country) + ","
	var4 := "Skills:" + convert_to_string(Skills)
	variables := var1 + var2 + var3 + var4
	//variables := "string1:string,int1:int,string3:string,string4:string"
	schema, err := orc.ParseSchema("struct<" + variables + ">")
	check(err)

	w, err := orc.NewWriter(f, orc.SetSchema(schema))
	check(err)

	for i := 0; i <= 4; i++ {
		err = w.Write(Name[i], Age[i], Country[i], Skills[i])
		if err != nil {
			fmt.Println("Error in for loop: ", err)
		}
	}
	w.Close()
	f.Sync()
	f.Close()
}

func main() {
	//test()
	writer()
}
