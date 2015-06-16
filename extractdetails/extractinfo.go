package extractdetails

import (
	"bytes"
	"fmt"
	"io"
	"net"
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
	re, _ = regexp.Compile("Registrant Phone Number:.*|Phone.*|Registrant Phone:.*")
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
	return regName, regPhone, regEmail

}

//顶级域名as提取函数
func asManage(details string) (regName, regPhone, regEmail string) {
	re, _ := regexp.Compile("Registrar:[\n].*")
	regName = re.FindString(details)
	if len(regName) != 0 {
		if len(strings.Split(regName, ":")) == 3 {
			regName = strings.TrimSpace(strings.Split(regName, ":")[1] + strings.Split(regName, ":")[2]) //获得注册人姓名并且去掉姓名前后的空白
		} else {
			regName = strings.TrimSpace(strings.Split(regName, ":")[1])
		}

	}

	return regName, regPhone, regEmail

}

//顶级com域名处理，分为三种情况来判断
func comManage(details, domainName, topServer string) (regName, regPhone, regEmail, newResult string) {
	//采集第一层数据，若发现注册人信息，说明已经提取到，则结束
	regName, regPhone, regEmail = generalManage(details)
	if len(regEmail) != 0 {
		return
	}
	var secWhois string
	//判断是否有标志位xxx
	if isXxx(details) {
		details, _ = GetDomainWhois(topServer, "="+domainName) //含有标志位，则向顶级服务器重新发送域名，域名前加=
		secWhois = extractSecSrv(details, domainName)          //提取二级注册服务器
	} else { //若不含有则直接提取二级服务器
		secWhois = extractSecSrv(details, domainName)
	}

	if len(secWhois) != 0 {
		details, _ := GetDomainWhois(secWhois, domainName)
		newResult = details
		regName, regPhone, regEmail = generalManage(details)
		return
	}
	return

}

//提取第二级服务器名称
func extractSecSrv(details, domainName string) string {
	re, _ := regexp.Compile("(?i)Whois Server:.*|Domain Name:.*|Registrar WHOIS Server:.*") //(?i)不区分大小写
	whoisSrvs := re.FindAllString(details, -1)
	for i, v := range whoisSrvs {
		if strings.TrimSpace(strings.ToLower(strings.Split(string(v), ":")[1])) == strings.TrimSpace(domainName) {
			fmt.Println(strings.Split(whoisSrvs[i+1], ":")[1])
			return strings.TrimSpace(strings.Split(whoisSrvs[i+1], ":")[1])
		}
	}
	return ""
}

//判断返回结果中，是否含有xxx标志位，若含有xxx则需要在查询的域名前加=
func isXxx(details string) bool {
	re, _ := regexp.Compile("xxx")
	xxx := re.FindString(details)
	if len(xxx) == 0 {
		return false
	}
	return true
}

//选择函数
func ExtractWhoisInfo(details, topServer, domainName string) (regName, regPhone, regEmail, newResult string) {
	newResult = ""
	switch topServer {
	case "whois.nic.us", "whois.nic.co", "whois.nic.website", "whois.nic.xxx", "whois.nic.press", "whois.nic.mn", "whois.nic.me", "whois.meregistry.net":
	case "whois.publicinterestregistry.net", "whois.pir.org", "whois.pandi.or.id", "whois2.afilias-grs.net", "whois.afilias.info", "whois.nic.uno", "whois.registry.pro":
	case "whois.kenic.or.ke", "whois.nic.xyz", "whois.neulevel.biz", "whois.dotmobiregistry.net", "whois.registry.in", "whois.adamsnames.tc", "whois.registrypro.pro":
	case "whois.cat", "whois.donuts.co", "whois.nic.pw", "whois.nic.net.ng", "whois-dub.mm-registry.com", "whois.nic.zm", "whois.nic.club", "whois.uniregistry.net":
		regName, regPhone, regEmail = generalManage(details)
		return
	case "whois.nic.uk":
		generalManage(details)
		return
	case "whois.nic.tr":
		regName, regPhone, regEmail = trManage(details)
		return
	case "whois.nic.as":
		regName, regPhone, regEmail = asManage(details)
		return
	case "whois.crsnic.net", "whois.verisign-grs.com", "jobswhois.verisign-grs.com":
		regName, regPhone, regEmail, newResult = comManage(details, domainName, topServer)
		return
	default:
		fmt.Println("meiyou")

	}
	return "", "", "", ""

}

//以下函数用来获取whois服务器信息，需要优化
//得到域名的whois完整信息
func GetDomainWhois(service, domain string) (string, error) {
	service = service + ":43"
	conn, err := net.Dial("tcp", service)
	if err != nil {
		return "", err
	}
	_, err = conn.Write([]byte(domain + "\r\n"))
	if err != nil {
		return "", err
	}

	result, err := readFully(conn)
	if err != nil {
		return "", err
	}
	// fmt.Println(result)
	return string(result), nil

}

//readFully完整的读取whois信息，并返回完整结果
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close() //关闭连接
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
