package gsl

// #include <gsl/gsl_sf.h>
// #include <gsl/gsl_errno.h>
import "C"
import (
	"errors"
	"unsafe"
)

// BesselJ0 ..
func BesselJ0(in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_J0_e(C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselJ1 ..
func BesselJ1(in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_J1_e(C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselJn ..
func BesselJn(level int, in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_Jn_e(C.int(level), C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselJnArr ..
func BesselJnArr(min, max int, in float64) (out []float64, err error) {
	if max-min+1 <= 0 {
		err = errors.New("max should be large than min")
		return
	}
	out = make([]float64, max-min+1) //[max - min + 1]float64{}
	if code := int(C.gsl_sf_bessel_Jn_array(C.int(min), C.int(max), C.double(in), (*C.double)(unsafe.Pointer(&out[0])))); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	return
}

// BesselY0 ..
func BesselY0(in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_Y0_e(C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselY1 ..
func BesselY1(in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_Y1_e(C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselYn ..
func BesselYn(level int, in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_Yn_e(C.int(level), C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselYnArr ..
func BesselYnArr(min, max int, in float64) (out []float64, err error) {
	if max-min+1 <= 0 {
		err = errors.New("max should be large than min")
		return
	}
	out = make([]float64, max-min+1) //[max - min + 1]float64{}
	if code := int(C.gsl_sf_bessel_Yn_array(C.int(min), C.int(max), C.double(in), (*C.double)(unsafe.Pointer(&out[0])))); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	return
}

// BesselI0 ..
func BesselI0(in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_I0_e(C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselI1 ..
func BesselI1(in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_I1_e(C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselIn ..
func BesselIn(level int, in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_In_e(C.int(level), C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselInArr ..
func BesselInArr(min, max int, in float64) (out []float64, err error) {
	if max-min+1 <= 0 {
		err = errors.New("max should be large than min")
		return
	}
	out = make([]float64, max-min+1) //[max - min + 1]float64{}
	if code := int(C.gsl_sf_bessel_In_array(C.int(min), C.int(max), C.double(in), (*C.double)(unsafe.Pointer(&out[0])))); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	return
}

// BesselK0 ..
func BesselK0(in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_K0_e(C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselK1 ..
func BesselK1(in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_K1_e(C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselKn ..
func BesselKn(level int, in float64) (out float64, err error) {
	var result C.gsl_sf_result
	if code := int(C.gsl_sf_bessel_Kn_e(C.int(level), C.double(in), &result)); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	out = float64(result.val)
	return
}

// BesselKnArr ..
func BesselKnArr(min, max int, in float64) (out []float64, err error) {
	if max-min+1 <= 0 {
		err = errors.New("max should be large than min")
		return
	}
	out = make([]float64, max-min+1) //[max - min + 1]float64{}
	if code := int(C.gsl_sf_bessel_Kn_array(C.int(min), C.int(max), C.double(in), (*C.double)(unsafe.Pointer(&out[0])))); code != 0 {
		err = errors.New(C.GoString(C.gsl_strerror(C.int(code))))
	}
	return
}
