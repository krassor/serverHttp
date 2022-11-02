package main

import (
	"bufio"
	//"context"
	"fmt"

	//"net"
	"os"
	"strings"
	"time"

	grpcm "github.com/krassor/serverHttp/internal/transport/grpc"
	httpm "github.com/krassor/serverHttp/internal/transport/rest"
)

//var DATA = make(map[string]Coin)

//var DATAFILE = "/tmp/dataFile.gob"

func main() {

	//arguments := os.Args
	// if len(arguments) == 1 {
	// 	fmt.Println("using default http port: ", PORT)
	// 	fmt.Println("using default grpc port: ", portGrpc)
	// } else {
	// 	PORT = ":" + arguments[1]
	// }

	go func() {
		if err := httpm.ServerHttpStart("8001"); err != nil {
			fmt.Println(err)
		}

	}()

	go func() {
		if err := grpcm.ServerGrpcStart("8080"); err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(1 * time.Second)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(string(text))) == "stop" {
			fmt.Println("Program exiting...")
			return
		}
	}
	//fmt.Println("End program")
}
