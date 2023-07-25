package structs

import (
	"fmt"
	"testing"
)

func TestSFPrintln(t *testing.T) {
	s := &struct {
		name string
		age  int
	}{
		name: "mrwhen",
		age:  18,
	}
	fmt.Println(s)
}
