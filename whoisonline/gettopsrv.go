//该代码获取顶级域名的whois服务器名称
package whoisonline

import (
	"math/rand"
	"time"
)

var tlds = map[string][]string{
	"us": []string{"whois.nic.us"},
	"tr": []string{"whois.nic.tr"},
	// "com":              []string{"whois.verisign-grs.com", "whois.crsnic.net"},
	// "net":              []string{"whois.verisign-grs.com", "whois.crsnic.net"},
	// "org":              []string{"whois.pir.org", "whois.publicinterestregistry.net"},
	// "info":             []string{"whois.afilias.info", "whois.afilias.net"},
	// "biz":              []string{"whois.neulevel.biz"},
	// "uk":               []string{"whois.nic.uk"},
	// "furnitureclub.uk": []string{"whois.nic.uk"},
	// "ca":               []string{"whois.cira.ca"},
	// "tel":              []string{"whois.nic.tel"},
	// "ie":               []string{"whois.iedr.ie", "whois.domainregistry.ie"},
	// "it":               []string{"whois.nic.it"},
	// "com.cn":           []string{"jsdlkfj"},
}

//根据顶级toptld域名，得到对应whois信息服务器
func GetWhoisSrv(top string) string {
	value, exists := tlds[top]
	if exists {
		r := rand.New(rand.NewSource(time.Now().UnixNano())) //产生随机种子
		length := len(value)
		i := r.Intn(length)
		return value[i]
	}
	return ""

}
