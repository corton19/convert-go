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

	"github.com/scritchley/orc"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writer() {
	// Data
	Name := []string{"Rose", "Smith", "William", "James", "Rolf"}
	Age := []int{28, 24, 29, 31, 21}
	Country := []string{"U.K.", "U.S.", "France", "Norway", "Denmark"}
	Skills := []string{"C", "Java", "PHP", "Python", "C++"}

	// Creates .orc file
	f, err := os.Create("employee.orc")
	check(err)

	// set ParseSchema to the following:
	t_start := "struct<"
	variables := "string1:string,int1:int,string3:string,string4:string"
	t_end := ">"
	schema, err := orc.ParseSchema(t_start + variables + t_end)
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
	writer()
}
