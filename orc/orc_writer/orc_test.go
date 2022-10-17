package orc

import "testing"

func TestCreateFile(t *testing.T) {
	Name := []string{"Rose", "Smith", "William", "James", "Rolf", "Matteo"}
	Age := []int{28, 24, 29, 31, 21, 26}
	Country := []string{"U.K.", "U.S.", "France", "Norway", "Denmark", "Italy"}
	writeFile("test", Name, Age, Country)
}
