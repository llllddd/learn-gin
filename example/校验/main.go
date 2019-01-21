package main

import (
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/binding"
	"github.com/smartwalle/validator"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2019-01-22"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2019-01-21"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	router.GET("/bookable", getBookable)
	router.Run(":3333")
}
