package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 小文字が含まれるかどうか
func IncludeLowercase(fl validator.FieldLevel) bool {
	return checkRegexp("[a-z]", fl.Field().String())
}

// 大文字が含まれるかどうか
func IncludeUppercase(fl validator.FieldLevel) bool {
	return checkRegexp("[A-Z]", fl.Field().String())
}

// 数値が含まれるかどうか
func IncludeNumeric(fl validator.FieldLevel) bool {
	return checkRegexp("[0-9]", fl.Field().String())
}

// 特殊記号が含まれるかどうか
func IncludeSymbol(fl validator.FieldLevel) bool {
	availableChar := checkRegexp(`^[0-9a-zA-Z\-^$*.@]+$`, fl.Field().String())
	checkIsSymbol := checkRegexp(`[\-^$*.@]`, fl.Field().String())

	return availableChar && checkIsSymbol
}

// 正規表現共通関数
func checkRegexp(reg, str string) bool {
	r := regexp.MustCompile(reg).Match([]byte(str))
	return r
}

func CustomValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("custom-low", IncludeLowercase)
	v.RegisterValidation("custom-upp", IncludeUppercase)
	v.RegisterValidation("custom-num", IncludeNumeric)
	v.RegisterValidation("custom-symbol", IncludeSymbol)

	return v
}
