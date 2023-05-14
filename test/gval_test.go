package test

import (
	"fmt"
	"github.com/PaesslerAG/gval"
	"testing"
)

func TestGval(t *testing.T) {
	val, err := gval.Evaluate("a & 2 > 0", map[string]interface{}{"a": 2})
	fmt.Println(val, err)
}
