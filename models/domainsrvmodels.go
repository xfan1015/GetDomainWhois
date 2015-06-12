//该代码为whois服务器代码
package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego"
	// "GetDomainWhois/whoisonline"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//定义whois服务器信息
type DomainSrv struct {
	Id          int64
	Name        string `orm:"index;null"` //orm 默认为255
	Tld         string `orm:"index;null"`
	RegNameAvl  string `orm:"size(10);null"` //注册人姓名是否可查
	RegEmailAvl string `orm:"size(10);null"` //注册人邮箱是否可查
	RegPhoneAvl string `orm:"size(10);null"` //注册人电话是否可查
}

//初始化模型
func init() {
	orm.RegisterModel(new(DomainSrv)) //注册模型
}

//获取顶级域名的whois服务器名称
func GetSrv(tld string) string {
	o := orm.NewOrm()
	srv := DomainSrv{Tld: tld}
	err := o.Read(&srv, "Tld") //指定字段进行查询
	if err == orm.ErrNoRows {
		fmt.Println("没有该顶级域名的服务器")
		return ""
	}
	srvName := srv.Name
	if strings.Contains(srvName, ",") { //包含两个whois服务器时候，随机选取一个
		r := rand.New(rand.NewSource(time.Now().UnixNano())) //产生随机种子
		srvNameSlice := strings.Split(srvName, ",")
		length := len(srvNameSlice)
		i := r.Intn(length)
		return srvNameSlice[i]
	} else { //只含有一个服务器
		return srvName
	}
}

func GetAllSrv() ([]*DomainSrv, int, error) {
	o := orm.NewOrm()
	dw := make([]*DomainSrv, 0)
	qs := o.QueryTable("domain_srv")
	_, err := qs.All(&dw)
	// _, err := qs.OrderBy("-id").Limit(20).All(&dw) //返回limite个数的最近查询数据，通过倒序id
	count := len(dw)
	return dw, count, err

}
