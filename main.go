package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"sync"

	"time"
)
//读取文件里的ip地址，如果ip分段，进行批量运算

//将ip添加到队列里面
//go协程从队列里取ip，并执行200次ping操作，计算延迟和成功率
//var wg sync.WaitGroup
//从datachan读取数据
var(

	results []result
	wg sync.WaitGroup
)
func worker(task chan string)  {
	defer wg.Done()
	for {
		ip,ok := <- task
		//如果通道中没有任务那么结束goroutine
		if !ok{
			return
		}
		raddr, err := net.ResolveIPAddr("ip", ip)
		if err != nil{
			log.Println(ip,"ResolveIPAddr err:",err)
		}

		//获取每个ip的ping信息
		re := result{}
		durations := make([]int,0,pingNums)
		for i:=0;i<pingNums;i++{
			if duration,err := sendICMPRequest(getICMP(uint16(i)), raddr);err == nil {
				durations = append(durations, duration)
			}
			//}else{
			//	fmt.Println(err)
			//}
			//设置ping间隔
			time.Sleep(pingInterval)
		}
		re.sucNums = len(durations)
		re.ip = ip
		re.latency = avarage(durations)
		//转化为float类型，保留小数
		re.sucRate = float32(float32(len(durations)) / float32(pingNums))
		results = append(results,re)

	}
}

//保存每个ip的测试时间、延迟、成功率
//排序，选择最优ip
func main()  {

	tasks := make(chan string,len(ips))
	runtime.GOMAXPROCS(4)
	start := time.Now()
	wg.Add(routineNums)
	//fmt.Println(ips)
	//生产者：写入通道
	for _,ip := range ips{
		tasks <- ip
	}

	//消费者:goroutine
	for i:=0;i<routineNums;i++{
		go worker(tasks)
	}

	//必须要先关闭通道
	close(tasks)
	wg.Wait()
	fmt.Println("Ping任务已结束")

	//按照成功率对结果进行排序
	fmt.Println("正在对结果进行排序....")

	results = quickSortBySucnums(results)
	re := getPartion(results)
	//保证每个区间都能得到排序
	for i:=0;i<=len(re);i++{
		if i == 0{
			//fmt.Println(results[:re[i]])
			quickSortByLantency(results[:re[i]])
		}else if i == len(re){
			//fmt.Println(results[re[i-1]:])
			quickSortByLantency(results[re[i-1]:])
		} else {
			//fmt.Println(results[re[i-1]:re[i]])
			quickSortByLantency(results[re[i-1]:re[i]])
		}


	}

	best = results[len(results)-1].ip

	//打印结果
	if len(results) > 10{
		for i:=0;i<10;i++{
			data := results[len(results)-1-i]
			dataStr := fmt.Sprintf("IP: %s ,平均延迟: %d ms,丢失: %d %%",data.ip,data.latency,100-int(data.sucRate*100))
			fmt.Println(dataStr)
		}

	} else {
		fmt.Println(results)
	}

	//更改dns记录
	if dnsChange == "true"{
		b := changeRec()
		if b {
			fmt.Println("更改DNS记录成功")
		}
	}


	end := time.Now().Sub(start)
	fmt.Printf("总耗时：%s",end)
}