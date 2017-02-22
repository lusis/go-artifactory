package main

import (
	"fmt"
	"os"

	artifactory "artifactory.v491"
)

func main() {
	client := artifactory.NewClientFromEnv()
	p, err := client.CreateUserApiKey()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", p)
		os.Exit(0)
	}
}
