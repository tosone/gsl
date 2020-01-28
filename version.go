package gsl

// #include <gsl/gsl_version.h>
import "C"

func Version() string {
	return string(C.GSL_VERSION)
}
