package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego"
	"GetDomainWhois/whoisonline"
	"fmt"
	"time"
)

type DomainWhois struct {
	Id             int64
	Domain         string `orm:"index"` //orm 默认为255
	Ip             string
	TopWhoisServer string
	SecWhoisServer string
	RegName        string    `orm:"index"`
	RegEmail       string    `orm:"index"`
	RegPhone       string    `orm:"index"`
	Details        string    `orm:"size(10000)"`
	QueryTime      time.Time `orm:"index"`
}

func init() {
	orm.RegisterModel(new(DomainWhois)) //注册模型
}

func GetAllWhois() ([]*DomainWhois, error) {
	o := orm.NewOrm()
	dw := make([]*DomainWhois, 0)
	qs := o.QueryTable("domain_whois")
	_, err := qs.All(&dw)
	return dw, err

}

func QueryDomain(queryDomain string) DomainWhois {
	o := orm.NewOrm()
	domain := DomainWhois{Domain: queryDomain}
	err := o.Read(&domain, "Domain") //指定字段进行查询
	if err == orm.ErrNoRows {
		fmt.Println("数据库未有该项,正在在线查询")
		// return domain
		return QueryOnline(queryDomain)
	} else {
		return domain
	}
}

func QueryOnline(queryDomain string) DomainWhois {
	domain := &whoisonline.Domain{}
	domain.ReturnWhois(queryDomain)
	// whois.Domain=domain.DomainName
	// whois.Details=domain.Details
	whois := &DomainWhois{
		Domain:         domain.DomainName,
		TopWhoisServer: domain.TopWhoisSrv,
		SecWhoisServer: domain.SecWhoisSrv,
		RegName:        domain.RegName,
		RegEmail:       domain.RegEmail,
		RegPhone:       domain.RegPhone,
		Details:        domain.Details,
		QueryTime:      time.Now(),
	}
	fmt.Println(whois.Details)
	return *whois
}
