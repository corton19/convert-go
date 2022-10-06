package orc

import "testing"

func TestCreateFile(t *testing.T) {
	data := [][]string{{"1", "2", "3", "4"}, {"5", "6", "7", "8"}}
	CreateFile("test", data)
}
