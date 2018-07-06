package check

import (
	"regexp"
	"strconv"
)

/*
 @检查号码是否正确
 @号码正确格式为：11个数字组成
*/
func CheckPhone(phone string) (bool, error) {
	r, _ := regexp.Compile("[0-9]+")
	if !r.MatchString(phone) {
		return false, nil
	}
	if len(phone) != 11 {
		return false, nil
	}
	return true, nil
}

/*
 @检查餐馆名是否正确
 @餐馆名正确格式为：汉字、英文字母、数字组成，不允许特殊字符
*/
func CheckName(name string) (bool, error) {
	r, _ := regexp.Compile("[^\u4e00-\u9fa5a-zA-Z0-9]+")
	if r.MatchString(name) {
		return false, nil
	}
	return true, nil
}

//Checkservertime: "08:00-20:00"
/*
 @检查时间是否正确
 @时间正确格式为："08:00-20:00"
*/
func Checkservertime(time string) (bool, error) {
	if len(time) != 11 {
		return false, nil
	}
	hour1, _ := strconv.Atoi(time[0:2])
	minute1, _ := strconv.Atoi(time[3:5])
	hour2, _ := strconv.Atoi(time[6:8])
	minute2, _ := strconv.Atoi(time[9:11])
	if hour1 >= 24 || hour2 >= 24 || minute1 >= 60 || minute2 >= 60 {
		return false, nil
	}
	return true, nil
}
