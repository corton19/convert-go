// TODO: 1st iteration: Convert data to csv/orc
// Input is assumed to be an array of type string

package main

import (
	"encoding/csv"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func data() {
	csvData := [][]string{
		{"Name", "City", "Country", "Skills"},
		{"Rose", "London", "U.K.", "C"},
		{"Smith", "Austin", "U.S.", "Java"},
		{"William", "Paris", "France", "PHP"},
		{"James", "Oslo", "Norway", "Python"},
	}

	//fmt.Println(data)

	recordFile, err := os.Create("employee.csv")
	check(err)

	defer recordFile.Close()

	w := csv.NewWriter(recordFile)
	err = w.WriteAll(csvData)
	check(err)

	recordFile.Sync()

}

func main() {
	data()

}
