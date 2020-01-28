package gsl

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	fmt.Printf("gsl version: %s\n", Version())
}
