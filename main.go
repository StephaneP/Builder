package main

import (
	"Builder/cmd"
	"Builder/logger"
	"Builder/utils"
	"fmt"
	"os"
	"time"
)

func main() {
	timestamp := time.Now().Format("20060102-150405")

	localLogFile := fmt.Sprintf("local-%s.log", timestamp)

	log, err := logger.NewDefaultLogger(localLogFile, "global.log")
	if err != nil {
		fmt.Printf("Error initializing logger: %s\n", err)
		os.Exit(1)
	}
	defer log.Close()

	if len(os.Args) > 1 {
		utils.Help()
		builderCommand := os.Args[1]
		if builderCommand == "init" {
			cmd.Init()
		} else if builderCommand == "config" {
			cmd.Config()
		} else {
			cmd.Builder()
		}
	} else {
		cmd.Builder()
	}
	fmt.Println("Build Complete 🔨")
}
