package orc

import (
	"fmt"
	"os"

	"github.com/scritchley/orc"
)

// CreateFile creates a .orc file
func CreateFile(filename string, data [][]string) {
	orc_file := filename + ".orc"
	f, err := os.Create(orc_file)
	if err != nil {
		fmt.Println("Error creating .orc file")
		panic(err)
	}

	w, err := orc.NewWriter(f, orc.SetSchema(SetSchemaStruct()))
	if err != nil {
		fmt.Println("Error: NewWriter in CreateFile")
		panic(err)
	}

	for a, b := range data {
		for c, _ := range b {
			err = w.Write(data[a][c])
		}
	}
	err = w.Close()
	f.Sync()
	f.Close()

}

func SetSchemaStruct() *orc.TypeDescription {
	schema, err := orc.ParseSchema("struct<string1:string>")
	if err != nil {
		fmt.Println("Error parsing Schema (ParseSchema)")
		panic(err)
	}
	return schema
}
