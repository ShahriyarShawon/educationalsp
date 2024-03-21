package main

import (
	"bufio"
	"educationalsp/lsp"
	"educationalsp/rpc"
	"encoding/json"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/shahriyar/Documents/programming/go/educationalsp/log.txt")
	logger.Println("Hello")
	logger.Println("World")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an err %s", err)
			continue
		}

		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with method: %s\n", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest

		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey we couldnt parse this %s", err)
		}

		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)
	}
}

func getLogger(filename string) *log.Logger {
	_, e := os.Stat(filename)
	if e != nil {
		os.Remove(filename)
	}
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0766)
	if err != nil {
		panic("hey not good file")
	}

	return log.New(logfile, "[educationalsp]", log.Ldate|log.Lshortfile)
}
