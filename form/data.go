package form

import (
	"fmt"
	"net/url"
	"regexp"
	"unicode/utf8"
)


var PhoneRX = regexp.MustCompile("(^\\+[0-9]{2}|^\\+[0-9]{2}\\(0\\)|^\\(\\+[0-9]{2}\\)\\(0\\)|^00[0-9]{2}|^0)([0-9]{9}$|[0-9\\-\\s]{10}$)")


var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")


type Input struct {
	Values  url.Values
	VErrors ValidationErrors
	CSRF    string
}


func (inVal *Input) MinLength(field string, d int) {
	value := inVal.Values.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		inVal.VErrors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", d))
	}

}
func (inVal *Input) ExactLength(field string, d int) {
	value := inVal.Values.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) != d {
		inVal.VErrors.Add(field, fmt.Sprintf("This field  should should contain exactly  %d characters)", d))
	}
}


func (inVal *Input) Required(fields ...string) {
	for _, f := range fields {
		value := inVal.Values.Get(f)
		if value == "" {
			inVal.VErrors.Add(f, "This field is required field")
		}
	}
}


func (inVal *Input) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := inVal.Values.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		inVal.VErrors.Add(field, "The value entered is invalid")
	}
}


func (inVal *Input) PasswordMatches(password string, confPassword string) {
	pwd := inVal.Values.Get(password)
	confPwd := inVal.Values.Get(confPassword)

	if pwd == "" || confPwd == "" {
		return
	}

	if pwd != confPwd {
		inVal.VErrors.Add(password, "The Password and Confim Password values did not match")
		inVal.VErrors.Add(confPassword, "The Password and Confim Password values did not match")
	}
}

// Valid checks if any form input validation has failed or not
func (inVal *Input) Valid() bool {
	return len(inVal.VErrors) == 0
}
