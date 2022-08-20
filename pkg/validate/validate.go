package validate

import "regexp"

func IsEmail(s string) bool {
	reg := regexp.MustCompile(`^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	return reg.MatchString(s)
}

func IsPhone(s string) bool {
	reg := regexp.MustCompile(`^((13[0-9])|(14[0-9])|(15[0-9])|(16[0-9])|(17[0-9])|(18[0-9])|(19[0-9]))\d{8}$`)
	return reg.MatchString(s)
}

func IsUrl(s string) bool {
	reg := regexp.MustCompile(`((http|https)://)(www.)?[a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)`)
	return reg.MatchString(s)
}
