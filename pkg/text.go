package pkg

import "regexp"

// ValidateEmail 检查邮箱格式
func ValidateEmail(email string) bool {
	//re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return re.MatchString(email)
}

// ValidatePhone 检查中国的手机号格式
func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}
