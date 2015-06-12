package extractdetails

import (
	"fmt"
	// "os"
	"regexp"
	"strings"
)

//通用提取函数
func generalManage(details string) (regName, regPhone, regEmail string) {

	//获得注册人姓名
	re, _ := regexp.Compile("Registrant Name:.*")
	regName = re.FindString(details)
	if len(regName) != 0 {
		regName = strings.TrimSpace(strings.Split(regName, ":")[1]) //获得注册人姓名并且去掉姓名前后的空白
	}
	//获得注册人电话
	re, _ = regexp.Compile("Registrant Phone Number:.*")
	regPhone = re.FindString(details)
	if len(regPhone) != 0 {
		regPhone = strings.TrimSpace(strings.Split(regPhone, ":")[1]) //获得电话并且去掉姓名前后的空白
	}
	//获得注册人邮箱
	re, _ = regexp.Compile("Registrant Email:.*")
	regEmail = re.FindString(details)
	if len(regEmail) != 0 {
		regEmail = strings.TrimSpace(strings.Split(regEmail, ":")[1]) //获得注册人姓名并且去掉姓名前后的空白
	}

	// fmt.Println(regName)
	// fmt.Println(regEmail)
	// fmt.Println(regPhone)
	return regName, regPhone, regEmail
}

//选择函数
func ExtractWhoisInfo(details, topServer string) (regName, regPhone, regEmail string) {

	switch topServer {
	case "whois.nic.us", "whois.nic.tr":
		regName, regPhone, regEmail = generalManage(details)
		return
	case "whois.nic.uk":
		generalManage(details)
		return
	default:
		fmt.Println("meiyou")
		// return "", "", ""
	}
	return "", "", ""

}
