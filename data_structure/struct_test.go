package data_structure

import (
	"fmt"
	"testing"
)

func TestTest1(t *testing.T) {
	//p1 := &People{Name: "Tom", Age: 18}
	p1 := &People{"Tom", 18}
	p1.SetName("Tom", 18)
	fmt.Printf(p1.getName())

	p2 := new(People)
	p2.Name = "Tom"
	p2.Age = 18
}

