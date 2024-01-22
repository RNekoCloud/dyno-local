package main

import (
	"dyno-local/config"
	"dyno-local/utils"
	"fmt"
)

func main() {
	// Setup configuration for DynamoDB
	dyno := config.SetupDynoConfig()

	// Stub Operation
	stub := utils.Dyno{
		Client:    dyno,
		TableName: "notes",
	}

	err := stub.CreateTable()

	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}

	fmt.Println("Successfully create table")
}
