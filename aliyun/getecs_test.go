package aliyun

import (
	"fmt"
	"testing"
)

func TestT0(t *testing.T) {
	dump(true)
}

func TestGetProd(t *testing.T) {
	bl := getProd(true)
	fmt.Println(bl)
}



