package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"time"
)

//参考链接：https://blog.csdn.net/simplelovecs/article/details/51146960
type ICMP struct {
	Type uint8 //icmp类型
	Code uint8 //ICMP的类型,标识生成的错误报文
	CheckSum uint16 //报文头校验值
	Identifier uint16 //标识icmp
	SequenceNum uint16 //序列号
}

//生成icmp报文头
func getICMP(seq uint16) ICMP {
	icmp := ICMP{
		Type:        8,
		Code:        0,
		CheckSum:    0,
		Identifier:  0,
		SequenceNum: seq,
	}

	var buffer bytes.Buffer
	//在网络中传输的数据需要是大端字节序的
	binary.Write(&buffer,binary.BigEndian,icmp)
	icmp.CheckSum = CheckSum(buffer.Bytes())
	buffer.Reset()
	return icmp
}

//发送请求，获取请求
func sendICMPRequest(icmp ICMP,destAddr *net.IPAddr) (duration int,err error) {
	//创建连接
	conn,err := net.DialIP("ip4:icmp",nil,destAddr)
	if err != nil{
		log.Println("net.DialIP err:",err)
		return -1,err
	}
	defer conn.Close()

	//构建icmp报文
	var buffer bytes.Buffer
	binary.Write(&buffer,binary.BigEndian,icmp)

	//向连接通道中发送报文
	if _,err := conn.Write(buffer.Bytes());err != nil{
		return -1,err
	}

	start := time.Now()
	//设置超时时间
	conn.SetReadDeadline(time.Now().Add(time.Second))
	recv := make([]byte,1024)
	//读取返回报文,如果超时，那么 err = "i/o timeout",这里只是一个简单超时判断
	receiveCnt,err := conn.Read(recv)
	if err != nil{
		return -1,err
	}
	end := time.Now()
	//精度
	duration = int(end.Sub(start).Nanoseconds() / 1e6)
	log.Printf("%d bytes from %s: seq = %d time = %dms\n",receiveCnt,destAddr.String(),icmp.SequenceNum,duration)
	return duration,err
}

//校验值
func CheckSum(data []byte) uint16 {
	var (
		sum uint32
		length int = len(data)
		index int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)
	return uint16(^sum)
}

