package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

var host = "127.0.0.1"
var port = "18080"

func main() {
	fmt.Println("started...")
	addr, err := net.ResolveUDPAddr("udp", host+":"+port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	fmt.Println(n, remoteAddr)
	fmt.Println(string(data))
	args, _ := getArgs()
	bStr := "{\"instances\":" + args[0] + ", \"connNum\":" + args[1] + ", \"ovsId\": 0 }"
	//b := []byte("{\"instances\":20, \"connNum\":100, \"ovsId\": 0 }")
	fmt.Println(bStr)
	b := []byte(bStr)
	conn.WriteToUDP(b, remoteAddr)
}

func getArgs() (args []string, err error) {

	file, err := os.Open("ARGS")
	if err != nil {
		fmt.Println("open ARGS failed.", err)
		return
	}

	defer file.Close()

	fd, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read ARGS failed.", err)
		return
	}
	strARGS := string(fd)

	args = strings.Split(strARGS, ",")
	return
}
