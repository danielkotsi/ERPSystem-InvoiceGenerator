package utils

import (
	"fmt"

	"github.com/signintech/gopdf"
)

func MakePDF() (resultpdf []byte, err error) {
	pdf := &gopdf.GoPdf{}
	fmt.Println(pdf)

	return resultpdf, err
}
