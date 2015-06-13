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
	re, _ := regexp.Compile("Registrant Name:.*|Organization Name.*|Person.*")
	regName = re.FindString(details)
	if len(regName) != 0 {
		regName = strings.TrimSpace(strings.Split(regName, ":")[1]) //获得注册人姓名并且去掉姓名前后的空白
	}
	//获得注册人电话
	re, _ = regexp.Compile("Registrant Phone Number:.*|Phone.*")
	regPhone = re.FindString(details)
	if len(regPhone) != 0 {
		regPhone = strings.TrimSpace(strings.Split(regPhone, ":")[1]) //获得电话并且去掉姓名前后的空白
	}
	//获得注册人邮箱
	re, _ = regexp.Compile("Registrant Email:.*|Registrant E-mail:.*")
	regEmail = re.FindString(details)
	if len(regEmail) != 0 {
		regEmail = strings.TrimSpace(strings.Split(regEmail, ":")[1]) //获得注册人姓名并且去掉姓名前后的空白
	}

	// fmt.Println(regName)
	// fmt.Println(regEmail)
	// fmt.Println(regPhone)
	return regName, regPhone, regEmail
}

//顶级域名tr提取函数
func trManage(details string) (regName, regPhone, regEmail string) {
	re, _ := regexp.Compile("Organization Name.*|Person.*")
	regName = re.FindString(details)
	if len(regName) != 0 {
		regName = strings.TrimSpace(strings.Split(regName, ":")[1]) //获得注册人姓名并且去掉姓名前后的空白
	}
	//获得注册人电话
	re, _ = regexp.Compile("Phone.*")
	regPhone = re.FindString(details)
	if len(regPhone) != 0 {
		regPhone = strings.TrimSpace(strings.Split(regPhone, ":")[1]) //获得电话并且去掉姓名前后的空白
	}
	//获得注册人邮箱
	re, _ = regexp.Compile(".*@.*")
	regEmail = re.FindString(details)
	if len(regEmail) != 0 {
		regEmail = strings.TrimSpace(regEmail) //获得注册人姓名并且去掉姓名前后的空白
	}

	// fmt.Println(regName)
	// fmt.Println(regEmail)
	// fmt.Println(regPhone)
	return regName, regPhone, regEmail

}

//选择函数
func ExtractWhoisInfo(details, topServer string) (regName, regPhone, regEmail string) {

	switch topServer {
	case "whois.nic.us", "whois.nic.co", "whois.nic.website", "whois.nic.xxx", "whois.nic.press", "whois.nic.mn", "whois.nic.me", "whois.meregistry.net":
	case "whois.publicinterestregistry.net", "whois.pir.org", "whois.pandi.or.id", "whois2.afilias-grs.net", "whois.afilias.info", "whois.nic.uno":
		regName, regPhone, regEmail = generalManage(details)
		return
	case "whois.nic.uk":
		generalManage(details)
		return
	case "whois.nic.tr":
		regName, regPhone, regEmail = trManage(details)
		return

	default:
		fmt.Println("meiyou")
		// return "", "", ""
	}
	return "", "", ""

}
