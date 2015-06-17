package models

import (
	"GetDomainWhois/extractdetails"
	"GetDomainWhois/gettld"
	// "GetDomainWhois/models"sss
	// "fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
	// "unicode/utf8"
)

type Domain struct {
	RawUrl      string //网址
	DomainName  string //可查询域名
	TopDomain   string //顶级域名
	topTld      string //要查询的域名
	Ip          string
	TopWhoisSrv string    //顶级域名服务器
	SecWhoisSrv string    //二级域名服务器
	RegName     string    //注册人姓名
	RegPhone    string    //注册人电话
	RegEmail    string    //注册人邮箱
	Details     string    //详细信息
	CreatedTime time.Time //创建时间
}

//
func (domain *Domain) init() {
	// tldTop, topDomain, domainName := getTld(domain.RawUrl)
	tldTop, topDomain, domainName := gettld.Extract_tld(domain.RawUrl)
	//以下代码为区分顶级域名服务器不同的情况
	var topWhoisSrv string
	if tldTop == "za" {
		topWhoisSrv = GetSrv(topDomain)
	} else {
		topWhoisSrv = GetSrv(tldTop)
	}

	if len(topWhoisSrv) == 0 {
		// fmt.Println("")
		os.Exit(1)
	}
	domain.topTld = tldTop
	domain.TopWhoisSrv = topWhoisSrv
	domain.TopDomain = topDomain
	domain.DomainName = domainName

}

//获取域名whois信息
func (domain *Domain) domainWhois() {
	re, _ := regexp.Compile("[\x00-\x7F]*")
	result, _ := GetDomainWhois(domain.TopWhoisSrv, domain.DomainName)
	var ip string
	// fmt.Println(domain.DomainName)
	ips, err := net.LookupHost(domain.DomainName)
	if err == nil {
		ip = strings.Join(ips, ",")
	}

	regname, regphone, regemail, newResult := extractdetails.ExtractWhoisInfo(result, domain.TopWhoisSrv, domain.DomainName)
	result1 := re.FindAllString(result, -1)
	result = strings.Join(result1, " ")
	regname1 := re.FindAllString(regname, -1)
	regname = strings.Join(regname1, "")
	domain.Details = strings.TrimSpace(result + newResult)
	domain.RegName = regname
	domain.RegEmail = regemail
	domain.RegPhone = regphone
	domain.Ip = ip

}

func (domain *Domain) ReturnWhois(rawurl string) {

	domain.RawUrl = rawurl
	domain.init()
	// fmt.Println(domain)
	domain.domainWhois()
}
