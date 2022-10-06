// 2nd iteration: Convert data to orc
// Input is assumed to be a 2D array of type string

// TODO (for 3rd iteration)
// Create a function which can determine array size of an unknown imported array
// The code should be less hard-coded (i.e., improve nested for loop at 49-61)
// Q: Is the array going to be 2D? Otherwise, edit/improve it

package main

import (
	"fmt"
	"os"

	"github.com/scritchley/orc"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func data() {
	data_to_file := [][]string{
		{"Name", "City", "Country", "Skills"},
		{"Rose", "London", "U.K.", "C"},
		{"Smith", "Austin", "U.S.", "Java"},
		{"William", "Paris", "France", "PHP"},
		{"James", "Oslo", "Norway", "Python"},
	}

	// Creates .orc file
	f, err := os.Create("employee.orc")
	check(err)
	//fmt.Println("Checkpoint NR. 1")

	// set ParseSchema to the following:
	schema, err := orc.ParseSchema("struct<string1:string>")
	check(err)
	//fmt.Println("Checkpoint NR. 2")

	w, err := orc.NewWriter(f, orc.SetSchema(schema))
	check(err)
	//fmt.Println("Checkpoint NR. 3")

	//j := len(data_to_file)
	//fmt.Println("length of array: ", j)

	for a, b := range data_to_file {
		//fmt.Println("length of a: ", a)
		//fmt.Println("length of b: ", b)
		for c, d := range b {
			err = w.Write(data_to_file[a][c])

			fmt.Println(b, d)
			// d is unnecessary, can be replaced with _ (blank identifiers)

			//fmt.Println("length of c: ", c)
			//fmt.Println("length of d: ", d)
		}
	}
	//fmt.Println(string1)
	//err = w.Write(string1)
	//check(err)

	err = w.Close()
	f.Sync()
	f.Close()
}

func main() {
	data()

}
