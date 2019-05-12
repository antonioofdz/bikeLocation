package handlers

import (
	"regexp"
)

func checkValidGuidRegExpFn(uuid string) bool {
	r := regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}")
	return r.MatchString(uuid)
}

/*
	Minimo 8 caracteres
	Maximo 15
	Al menos una letra mayúscula
	Al menos una letra minucula
	Al menos un dígito
	No espacios en blanco
	Al menos 1 caracter especial
*/
func checkValidPassword(s string) bool {
	lower_cond := regexp.MustCompile(`[a-z]`)
	upper_cond := regexp.MustCompile(`[A-Z]`)
	digit_cond := regexp.MustCompile(`[0-9]`)
	//whole_cond := regexp.MustCompile(`^[0-9a-zA-Z!@#$%^&*()]*$`)
	return lower_cond.MatchString(s) && upper_cond.MatchString(s) && digit_cond.MatchString(s) && len(s) >= 8
}
