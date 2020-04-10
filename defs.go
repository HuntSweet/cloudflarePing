package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"time"
)


var (

	//所有ip
	ips  []string

	//ping配置
	//ping次数
	pingNums int
	//ping间隔
	pingInterval = time.Millisecond * 1000
	//协程数目
	routineNums int
	//是否自动更换dns解析
	dnsChange string

	//dns配置
	login_token string
	domain_id string
	record_id string
	record_type  string
	record_line string
	weight string
	format = "json"

	//返回的最优ip
	best string

)

type result struct {
	ip string
	latency int
	sucRate float32
	sucNums int
}

func init()  {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	ipFile := cfg.Section("ping").Key("ipFile").String()

	ips = getIps(ipFile)

	//ping次数
	pingNums,_ = cfg.Section("ping").Key("pingNums").Int()
	//协程数目
	fmt.Println(pingNums)
	routineNums,_ = cfg.Section("ping").Key("routineNums").Int()

	//dns配置
	//是否自动更换dns解析
	dnsChange = cfg.Section("dnspod").Key("isChange").String()
	login_token = cfg.Section("dnspod").Key("login_token").String()
	domain_id = cfg.Section("dnspod").Key("domain_id").String()
	record_id = cfg.Section("dnspod").Key("record_id").String()
	record_type  = cfg.Section("dnspod").Key("record_type").String()
	record_line = cfg.Section("dnspod").Key("record_line").String()
	weight = cfg.Section("dnspod").Key("weight").String()


}