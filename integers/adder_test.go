package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6

	/**
	1. install godoc by "go install golang.org/x/tools/cmd/godoc"
	2. run godoc -http=:6060
	3. browse http://localhost:6060/pkg
	4. find my package(integers, hello)
	5. find example for function Add
	*/
}
