package whoisonline

import (
	// "fmt"
	"os"
	"time"
)

type Domain struct {
	RawUrl      string //网址
	DomainName  string //可查询域名
	TopDomain   string //顶级域名
	topTld      string
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
	tldTop, topDomain, domainName := getTld(domain.RawUrl)

	topWhoisSrv := GetWhoisSrv(tldTop)
	if len(topWhoisSrv) == 0 {
		os.Exit(1)
	}
	domain.topTld = tldTop
	domain.TopWhoisSrv = topWhoisSrv
	domain.TopDomain = topDomain
	domain.DomainName = domainName

}

//获取顶级域名或者二级域名
func getTld(rawurl string) (string, string, string) {
	tldTop, topDomain, domainName := Extract_tld(rawurl)
	if len(domainName) == 0 {
		os.Exit(1)
	}
	// if !sec {
	// 	return tldTop, topDomain, domainName
	// }
	// return topDomain, topDomain, domainName
	//二级域名查询代码要写在这里
	return tldTop, topDomain, domainName

}

//获取域名whois信息
func (domain *Domain) domainWhois() {
	result, _ := GetDomainWhois(domain.TopWhoisSrv, domain.DomainName)
	// fmt.Println(result)
	// fmt.Println(err)
	domain.Details = result
}

func (domain *Domain) ReturnWhois(rawurl string) {

	domain.RawUrl = rawurl
	domain.init()
	// fmt.Println(domain)
	domain.domainWhois()
}
