//该包用来获取输入网址的顶级域名(二级域名)和查询域名(domainName)
package gettld

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strings"
)

//打开文件
func openfile(filename string) string {
	//打开tld文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		// fmt.Println("...")
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)

}

//从rawurl中提取host
func extract_host(rawurl string) (host string) {
	rawurl = strings.ToLower(rawurl)
	flag, _ := regexp.MatchString("http://.*", rawurl)
	if flag != true {
		rawurl = "http://" + rawurl
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	host = u.Host
	return

}

//从rawurl中提取顶级域名服务商,tldTop用来查询，topDomain用来找注册商服务器，domainName用来查询的域名
func Extract_tld(rawurl string) (tldTop, topDomain, domainName string) {
	filename := "./res/data.txt"
	data := openfile(filename)
	host := extract_host(rawurl)
	host_slice := strings.Split(host, ".")
	length := len(host_slice)
	tldTop = host_slice[length-1]
	// fmt.Println(tldTop)
	//得到顶级域名和主机名
	for i := 0; i < length; i++ {
		expr := ""
		for j := i; j < length; j++ {
			expr = expr + "." + host_slice[j]
		}

		reg, _ := regexp.CompilePOSIX("^" + expr[1:])
		// reg, _ := regexp.CompilePOSIX("^" + expr[1:] + "$")
		if reg.MatchString(data) {
			if i != 0 {
				domainName = host_slice[i-1] + expr
			} else {
				domainName = expr
			}
			topDomain = expr[1:]
			return
		}
	}
	return
}
