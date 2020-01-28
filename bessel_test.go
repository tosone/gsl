package gsl

import (
	"fmt"
	"log"
	"testing"
)

func TestBesselJ0(t *testing.T) {
	if out, err := BesselJ0(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J0(1) = %f\n", out)
	}
}

func BenchmarkBesselJ0(b *testing.B) {
	if out, err := BesselJ0(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J0(1) = %f\n", out)
	}
}

func TestBesselJ1(t *testing.T) {
	if out, err := BesselJ1(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J1(1) = %f\n", out)
	}
}

func BenchmarkBesselJ1(b *testing.B) {
	if out, err := BesselJ1(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J1(1) = %f\n", out)
	}
}

func TestBesselJn(t *testing.T) {
	if out, err := BesselJn(0, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J0(1) = %f\n", out)
	}
	if out, err := BesselJn(1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J1(1) =  %f\n", out)
	}
}

func BenchmarkBesselJn(b *testing.B) {
	if out, err := BesselJn(0, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J0(1) = %f\n", out)
	}
	if out, err := BesselJn(1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J1(1) =  %f\n", out)
	}
}

func TestBesselJnArr(t *testing.T) {
	if out, err := BesselJnArr(0, 1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J0(1) = %f\n", out[0])
		fmt.Printf("J1(1) = %f\n", out[1])
	}
}

func BenchmarkBesselJnArr(b *testing.B) {
	if out, err := BesselJnArr(0, 1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("J0(1) = %f\n", out[0])
		fmt.Printf("J1(1) = %f\n", out[1])
	}
}

func TestBesselY0(t *testing.T) {
	if out, err := BesselY0(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y0(1) = %f\n", out)
	}
}

func BenchmarkBesselY0(b *testing.B) {
	if out, err := BesselY0(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y0(1) = %f\n", out)
	}
}

func TestBesselY1(t *testing.T) {
	if out, err := BesselY1(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y1(1) = %f\n", out)
	}
}

func BenchmarkBesselY1(b *testing.B) {
	if out, err := BesselY1(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y1(1) = %f\n", out)
	}
}

func TestBesselYn(t *testing.T) {
	if out, err := BesselYn(0, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y0(1) = %f\n", out)
	}
	if out, err := BesselYn(1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y1(1) =  %f\n", out)
	}
}

func BenchmarkBesselYn(b *testing.B) {
	if out, err := BesselYn(0, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y0(1) = %f\n", out)
	}
	if out, err := BesselYn(1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y1(1) =  %f\n", out)
	}
}

func TestBesselYnArr(t *testing.T) {
	if out, err := BesselYnArr(0, 1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y0(1) = %f\n", out[0])
		fmt.Printf("Y1(1) = %f\n", out[1])
	}
}

func BenchmarkYnArr(b *testing.B) {
	if out, err := BesselYnArr(0, 1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Y0(1) = %f\n", out[0])
		fmt.Printf("Y1(1) = %f\n", out[1])
	}
}

func TestBesselI0(t *testing.T) {
	if out, err := BesselI0(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I0(1) = %f\n", out)
	}
}

func BenchmarkBesselI0(b *testing.B) {
	if out, err := BesselI0(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I0(1) = %f\n", out)
	}
}

func TestBesselI1(t *testing.T) {
	if out, err := BesselI1(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I1(1) = %f\n", out)
	}
}

func BenchmarkBesselI1(b *testing.B) {
	if out, err := BesselI1(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I1(1) = %f\n", out)
	}
}

func TestBesselIn(t *testing.T) {
	if out, err := BesselIn(0, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I0(1) = %f\n", out)
	}
	if out, err := BesselIn(1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I1(1) =  %f\n", out)
	}
}

func BenchmarkBesselIn(b *testing.B) {
	if out, err := BesselIn(0, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I0(1) = %f\n", out)
	}
	if out, err := BesselIn(1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I1(1) =  %f\n", out)
	}
}

func TestBesselInArr(t *testing.T) {
	if out, err := BesselInArr(0, 1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I0(1) = %f\n", out[0])
		fmt.Printf("I1(1) = %f\n", out[1])
	}
}

func BenchmarkBesselInArr(b *testing.B) {
	if out, err := BesselInArr(0, 1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("I0(1) = %f\n", out[0])
		fmt.Printf("I1(1) = %f\n", out[1])
	}
}

func TestBesselK0(t *testing.T) {
	if out, err := BesselK0(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("K0(1) = %f\n", out)
	}
}

func BenchmarkBesselK0(b *testing.B) {
	if out, err := BesselK0(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("K0(1) = %f\n", out)
	}
}

func BenchmarkBesselK1(b *testing.B) {
	if out, err := BesselK1(1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("K1(1) = %f\n", out)
	}
}

func BenchmarkBesselKn(b *testing.B) {
	if out, err := BesselKn(0, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("K0(1) = %f\n", out)
	}
	if out, err := BesselKn(1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("K1(1) =  %f\n", out)
	}
}

func BenchmarkBesselKnArr(b *testing.B) {
	if out, err := BesselKnArr(0, 1, 1.0); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("K0(1) = %f\n", out[0])
		fmt.Printf("K1(1) = %f\n", out[1])
	}
}
