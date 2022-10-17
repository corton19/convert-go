package orc

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/scritchley/orc"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Converts any type of input to string (assumed)
func convert_to_string(s any) string {
	// Debug code
	//obtype := reflect.TypeOf(s)
	//fmt.Println("obtype: ", obtype)

	switch s.(type) {
	case int, float64, string, int64, uint32:
		b := reflect.TypeOf(s).Name()
		//fmt.Println("b: ", b)
		return b
	case []string, []int, []int64, []float32, []float64:
		b := reflect.TypeOf(s).Elem().Name()
		//fmt.Println("s.([]string): ", s.([]string))
		//fmt.Println("b: ", b)
		return b
	}
	return "err"
}

// Creates .orc file
func createFile(filename string) *os.File {
	// Create .orc file
	f, err := os.Create(filename)
	check(err)
	return f
}

// Gets array size of a given input
func getLength(s ...interface{}) interface{} {
	//fmt.Println("s[0]: ", s[0])
	switch s[0].(type) {
	case int, float64, string, int64, uint32:
		a := 1
		return a
	// a smarter way to do this MUST be possible
	case []string:
		a := len(s[0].([]string))
		return a
	case []int:
		a := len(s[0].([]int))
		return a
	case []int64:
		a := len(s[0].([]int64))
		return a
	case []float32:
		a := len(s[0].([]float32))
		return a
	case []float64:
		a := len(s[0].([]float64))
		return a
	}
	return "XXXX"
}

func writeFile(filename string, s ...interface{}) {
	// Variable declaration
	max_length := len(s) - 1             // amount of inputs
	arr_length := getLength(s...).(int)  // array size of given input
	parString := make([]string, 0)       // used to create Schema ("string1:string" etc.)
	var res string                       // concates all strings from parString to create Schema
	a := make([]interface{}, arr_length) // creates a duplicate of our inputs
	copy(a, s)                           // (used to write rows for the orc file)

	// Create .orc file
	f := createFile(filename + ".orc")

	// Create Schema
	for i := 0; i <= max_length; i++ {
		if i < max_length {
			parString = append(parString, convert_to_string(s[i])+strconv.Itoa(i)+":"+convert_to_string(s[i])+",")
		} else if i == max_length { // aka. the final 'string'
			parString = append(parString, convert_to_string(s[i])+strconv.Itoa(i)+":"+convert_to_string(s[i]))
		}
		res += parString[i]
	}

	variables := res // example: res = "string1:string,int1:int,string3:string,string4:string"

	schema, err := orc.ParseSchema("struct<" + variables + ">") // the schema is used as template to write orc file
	check(err)

	w, err := orc.NewWriter(f, orc.SetSchema(schema))
	check(err)

	for i := 0; i < arr_length; i++ {
		for j := 0; j <= max_length; j++ {
			switch a[j].(type) {
			case ([]string):
				s[j] = a[j].([]string)[i]
				fmt.Println("s[", j, "]: ", s[j]) // debug
			case ([]int):
				s[j] = a[j].([]int)[i]
				fmt.Println("s[", j, "]: ", s[j]) // debug
			case ([]float64):
				s[j] = a[j].([]float64)[i]
				fmt.Println("s[", j, "]: ", s[j]) // debug
			case ([]int64):
				s[j] = a[j].([]int64)[i]
				fmt.Println("s[", j, "]: ", s[j]) // debug
			}
		}
		err = w.Write(s...)
		if err != nil {
			fmt.Println("Error (ORC.Write): ", err)
		}
	}

	w.Close()
	f.Sync()
	f.Close()
}
