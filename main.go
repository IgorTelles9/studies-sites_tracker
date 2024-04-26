package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const watches = 3
const delay = 10

func main() {
	showIntro()
	for {
		showOptions()
		command := getCommand()
		switch command {
		case 1:
			watch()
		case 2:
			printLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Command not found...")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Igor"
	var version float32 = 1.1
	fmt.Println("Hello,", name)
	fmt.Println("Program version:", version)
}

func showOptions() {
	fmt.Println("Choose an option")
	fmt.Println("1- Watch")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func getCommand() int {
	var command int
	fmt.Scanf("%d", &command)
	return command
}

func watch() {
	addrs := getAddressesFromFile()
	for i := 0; i < watches; i++ {
		fmt.Println("------------------------------------------------------------------------")
		fmt.Println("Watching... (", i+1, "/", watches, ")")
		for _, addr := range addrs {
			checkAddress(addr)
		}
		fmt.Println("Checking again in", delay, "seconds...")
		time.Sleep(delay * time.Second)
		fmt.Println("------------------------------------------------------------------------")
	}
}

func getAddressesFromFile() []string {
	var addresses []string
	file, err := os.Open("hello/addresses.txt")
	if err != nil {
		fmt.Println("[ERROR] Error ope ning file", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		addresses = append(addresses, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return addresses

}

func checkAddress(addr string) {
	res, err := http.Get(addr)
	if err != nil {
		fmt.Println("[ERROR] Error", err, "accessing address", addr)
	}

	if res.StatusCode == 200 {
		fmt.Println("Address", addr, "is online!")
		logStatus(addr, true)
	} else {
		fmt.Println("Address", addr, "is down with the code ", res.StatusCode)
		logStatus(addr, false)
	}
}

func logStatus(addr string, status bool) {
	file, err := os.OpenFile("hello/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("[ERROR] Error logging status. Address:", addr, "Status:", status, "Error:", err)
	}
	file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " - " + addr + " - online:" + strconv.FormatBool(status) + "\n")
}

func printLogs() {
	fmt.Println("Showing logs...")
	file, err := os.ReadFile("hello/log.txt")
	if err != nil {
		fmt.Println("[ERROR] Error opening log files ", err)
	}
	fmt.Println(string(file))

}
