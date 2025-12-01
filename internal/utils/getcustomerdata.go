package utils

import (
	"net/http"

	"github.com/go-playground/form/v4"
)

func ParseFormData(r *http.Request, data any) (err error) {
	decoder := form.NewDecoder()
	if err := r.ParseMultipartForm(20000000); err != nil {
		return err
	}

	if err := decoder.Decode(data, r.Form); err != nil {
		return err
	}

	//we are going to skip this implementation since things are getting out of hand with the parser and we are goint to use an external library instead
	//
	// v := reflect.ValueOf(data)
	// if v.Kind() != reflect.Ptr {
	// 	return errors.New("data must be a pointer to a struct")
	// }
	//
	// v = v.Elem()
	// if v.Kind() != reflect.Struct {
	// 	return errors.New("error getting the actual struct from the pointer")
	// }
	//
	// t := v.Type()
	//
	// for i := 0; i < t.NumField(); i++ {
	//
	// 	key := t.Field(i).Tag.Get("form")
	// 	if key == "" {
	// 		key = t.Field(i).Name
	// 	}
	//
	// 	values, exists := r.Form[key]
	// 	if !exists || len(values) == 0 {
	// 		continue
	// 	}
	//
	// 	switch v.Field(i).Kind() {
	// 	case reflect.String:
	// 		v.Field(i).SetString(values[0])
	// 	case reflect.Int:
	// 		n, _ := strconv.Atoi(values[0])
	// 		v.Field(i).SetInt(int64(n))
	// 	case reflect.Float64:
	// 		n, _ := strconv.ParseFloat(values[0], 64)
	// 		v.Field(i).SetFloat(n)
	// 	case reflect.Bool:
	// 		b, _ := strconv.ParseBool(values[0])
	// 		v.Field(i).SetBool(b)
	// 	//this only supports slices of string in the structs given to be recieve the form data
	// 	case reflect.Slice:
	// 		//you can add a switch around the v.Field(i).Type().Elem().Kind() to handle []int and other types of slices as well
	// 		if v.Field(i).Type().Elem().Kind() == reflect.String {
	// 			v.Field(i).Set(reflect.ValueOf(values))
	// 		}
	//
	// 	}
	// }

	return nil
}
