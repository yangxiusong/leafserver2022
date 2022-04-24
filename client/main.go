package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"github.com/name5566/leaf/log"
)

var readLen int
var bufReadLen int
var buf []byte

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	var wg sync.WaitGroup
	helloTest(conn, &wg)
	loginTest(conn, &wg)
	addTest(conn, &wg)
	go readMsg(conn, &wg)
	wg.Wait()
}

func readMsg(conn net.Conn, wg *sync.WaitGroup) {
	reader := bufio.NewReader(conn)
	for {
		peek, err := reader.Peek(2)
		if err != nil {
			// log.Debug("peek err:%s", err)
			continue
		}
		log.Debug("peek:%v\n", peek)
		buffer := bytes.NewBuffer(peek)
		var length uint16
		err = binary.Read(buffer, binary.BigEndian, &length)
		if err != nil {
			log.Debug("read err:%+v", err)
			continue
		}
		log.Debug("length:%d", length)
		if int32(reader.Buffered()) < int32(int(length))+2 {
			log.Debug("bufM:%+v", err)
			continue
		}
		data := make([]byte, length+2)
		_, err = reader.Read(data)
		if err != nil {
			log.Debug("read2 err:%+v", err)
			continue
		}

		// fmt.Printf("receive data:%+v \n", data[2:])
		var jdata interface{}
		err = json.Unmarshal(data[2:], &jdata)
		if err != nil {
			log.Debug("Unmarshal err:%+v", err)
			continue
		}
		fmt.Printf("receive jdata ==>  %+v\n", jdata)
		wg.Done()

	}
}

func helloTest(conn net.Conn, wg *sync.WaitGroup) {
	data := []byte(`{
		"Hello": {
			"Name": "leaf"
		}
	}`)

	m := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	wg.Add(1)
	conn.Write(m)
}

func loginTest(conn net.Conn, wg *sync.WaitGroup) {
	data := []byte(`{
		"Login": {
			"Name": "username",
			"Pwd": "123456"
		}
	}`)

	m := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	wg.Add(1)
	conn.Write(m)

}

func addTest(conn net.Conn, wg *sync.WaitGroup) {
	data := []byte(`{
		"Add": {
			"A": 100,
			"B": 200
		}
	}`)

	m := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)
	wg.Add(1)
	conn.Write(m)

}
