package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "green tea",
		Price: 1,
		SKU: "a-a-a",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
