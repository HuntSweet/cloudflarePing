package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//import "net/http"

//domain_id 域名ID，必选
//record_id 记录ID，必选
//sub_domain 主机记录，默认@，如 www，可选
//record_type 记录类型，通过API记录类型获得，大写英文，比如：A，必选
//record_line 记录线路，通过API记录线路获得，中文，比如：默认，必选
//value 记录值, 如 IP:200.200.200.200, CNAME: cname.dnspod.com., MX: mail.dnspod.com.，必选
//mx {1-20} MX优先级, 当记录类型是 MX 时有效，范围1-20, mx记录必选
//ttl {1-604800} TTL，范围1-604800，不同等级域名最小值不同，可选
//status [“enable”, “disable”]，记录状态，默认为”enable”，如果传入”disable”，解析不会生效，也不会验证负载均衡的限制，可选
//weight 权重信息，0到100的整数，可选。仅企业 VIP 域名可用，0 表示关闭，留空或者不传该参数，表示不设置权重信息
//注意事项：
//如果1小时之内，提交了超过5次没有任何变动的记录修改请求，该记录会被系统锁定1小时，不允许再次修改。
//示例:
//
//curl -X POST https://dnsapi.cn/Record.Modify -d 'login_token=LOGIN_TOKEN&format=json&domain_id=2317346&record_id=16894439&sub_domain=www&value=3.2.2.2&record_type=A&record_line=默认'
//返回参考：
//
//JSON:
//
//{
//"status": {
//"code":"1",
//"message":"Action completed successful",
//"created_at":"2015-01-18 16:53:23"
//},
//"record": {
//"id":16894439,
//"name":"@",
//"value":"3.2.2.2",
//"status":"enable"
//}
//}
//字段说明:
//id: 记录ID, 即为 record_id
//name: 子域名
//value”: 记录值
//status”: 记录状态

//record_id: "560520430"
//domain: "uowo.tk"
//sub_domain: "@"
//record_type: "A"
//record_line: "移动"
//ttl: "600"
//weight: "3"
//value: "198.41.214.134"
//mx: 5
//api: "Record.CheckImpact"


//都要大写
type Resp struct {
	Status Status `json:"status"`
}

type Status struct {
	//返回类型为unicode
	Message interface{} `json:"message"`
	Code string `json:"code"`

}
func changeRec() bool {

	client := http.Client{}

	//postdata为，不能post json格式，format格式是json
	resp,err := client.PostForm("https://dnsapi.cn/Record.Modify",url.Values{
		"login_token":{login_token},
		"domain_id":{domain_id},
		"record_line":{record_line},
		"record_type":{record_type},
		"record_id":{record_id},
		"value":{best},
		"format":{format},
		"weight":{weight},
	})
	if err != nil{
		log.Println("post data err:",err)
		return false
	}
	defer resp.Body.Close()

	recv,_ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(recv))
	mes := Resp{}
	err = json.Unmarshal(recv,&mes)
	if err != nil{
		log.Println("json unmarshal err:",err)
		return false
	}
	//fmt.Println(mes.Status.Code,mes.Status.Message)

	if mes.Status.Code == "1"{
		return true
	}

	return false
}