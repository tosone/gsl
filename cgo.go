package gsl

// #cgo LDFLAGS: -lgsl -lgslcblas
import "C"

// #cgo CFLAGS: -I${SRCDIR}/drivers/include/
// #cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/drivers/darwin-lib
// #cgo linux,amd64 LDFLAGS: -L${SRCDIR}/drivers/linux-lib
