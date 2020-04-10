package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var(

)

//对-的处理
//对/的处理
func getIps(ipfile string) (ips []string) {
	f, err := os.Open(ipfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	LinsNum:=CountFileLine(ipfile)


	for lines:=0;lines<=LinsNum;lines++ {
		line, _ := rd.ReadString('\n') //以'\n'为结束符读入一行
		line = strings.TrimSpace(line)
		if len(line)>5 && len(strings.Split(line, "."))==4 {
			fmt.Println(strings.Split(line,".")[3])
			if strings.Split(line,".")[3] == "*"{
				for i:=0;i<256;i++{
					iStr := strconv.Itoa(i)
					ip := strings.Split(line, "*")[0] + iStr
					ips = append(ips,ip)
				}
			}else{
				ips=append(ips,line)
			}

		}
	}

	return ips
}

func CountFileLine(name string) (count int) {
	data, _ := ioutil.ReadFile(name)
	count = 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	return count
}

//func ()  {
//
//}