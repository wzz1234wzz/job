package helper

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"fmt"
)

const (
	constMobileRegexp = "^1[0-9]{10}$"
	constLowerRegexp  = "[a-z]"
	constUpperRegexp  = "[A-Z]"
	constNumberRegexp = "[0-9]"
	// supported special characters:
	// ` ~ ! @ # $ % ^ & * _ + - = ( ) [ ] { } < > \ | ; : ' " , . / ?
	constSpecialCharRegexp = "[`" + `~!@#$%^&*_+\-=()\[\]{}<>\\|;:'",./?]`

	constMinPasswdLength      = 8
	constMaxPasswdLength      = 16
	constMinPasswdCombination = 2
)

// PwdGenerate PwdGenerate
func PwdGenerate(pwd string) (pwdHash string, err error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(h), nil
}

// PwdVerify PwdVerify
func PwdVerify(pwd string, pwdHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(pwd))
	return err
}

// IsPasswordValid tells whether the password is valid:
// 8-16 characters, 2 or more combination of lower case letters, upper case letters, numbers, and special characters
func IsPasswordValid(passwd string) bool {
	if len(passwd) < constMinPasswdLength || len(passwd) > constMaxPasswdLength {
		return false
	}

	lower := regexp.MustCompile(constLowerRegexp)
	upper := regexp.MustCompile(constUpperRegexp)
	num := regexp.MustCompile(constNumberRegexp)
	special := regexp.MustCompile(constSpecialCharRegexp)

	re := regexp.MustCompile(fmt.Sprintf("^(%v|%v|%v|%v)+$", constLowerRegexp, constUpperRegexp, constNumberRegexp, constSpecialCharRegexp))
	if !re.MatchString(passwd) {
		return false
	}

	regexps := []*regexp.Regexp{lower, upper, num, special}
	t := 0
	for _, r := range regexps {
		if r.MatchString(passwd) {
			t++
		}
	}
	return t >= constMinPasswdCombination
}