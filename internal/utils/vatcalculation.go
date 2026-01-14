package utils

type vatCat int

const (
	Key1 vatCat = iota + 1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	Key10
	vatCategory
)

var vatValue = []float64{
	0.0,
	0.24,
	0.13,
	0.06,
	0.17,
	0.09,
	0.04,
	0.00,
	0.00,
	0.03,
	0.04,
}

var vatNames = []int{
	24,
	13,
	6,
	17,
	9,
	4,
	0,
	0,
	3,
	4,
}

func VatNames(i int) int {
	if i < 1 || i >= len(vatValue) {
		return 0
	}
	return vatNames[i]
}

func Vat(i int) float64 {
	if i < 1 || i >= len(vatValue) {
		return 0
	}
	return vatValue[i]
}
