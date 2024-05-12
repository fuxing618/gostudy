package simple_factory

import (
	"fmt"
	"testing"
)

func TestProduct(t *testing.T) {
	f :=  &Factory{}
	product1 := f.createProduct(p1)
	if product1 == nil {
		t.Error("NewProduct should return a Product")
	}

	product2 := f.createProduct(p2)
	if product2 == nil {
		t.Error("NewProduct should return a Product")
	}
	fmt.Printf("product1:%+v\n", product1)
	fmt.Printf("product2:%+v\n", product2)
}

