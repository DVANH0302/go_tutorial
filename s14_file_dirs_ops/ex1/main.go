package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("config.env")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
		log.Fatal("Error reading file info:", err)
	}

	mode := fileInfo.Mode()
	fmt.Println("Current permission:", mode.Perm())

	isOtherReadable := mode&0004 != 0
	isOtherWritable := mode&0002 != 0

	if isOtherReadable || isOtherWritable {
		fmt.Println("Unsafe permission detected, fixing...")

		if err := os.Chmod("config.env", 0770); err != nil {
			log.Fatal("Failed to change permission:", err)
		}

		// Re-stat to get updated permissions
		fileInfo, err = os.Stat("config.env")
		if err != nil {
			log.Fatal("Error reading updated file info:", err)
		}
		fmt.Println("Updated permission:", fileInfo.Mode().Perm())
	} else {
		fmt.Println("Permission is correct:", mode.Perm())
	}
}
