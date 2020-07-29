package fhttp

import (
	"fmt"
	"testing"
)

func TestFContext(t *testing.T) {
	resp, e := Request(FContext{}).Get("https://www.fox.one")
	if e != nil {
		fmt.Println()
	}

	fmt.Println(resp)
}
