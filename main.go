package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Blast...")
	removePackageLock()
	removeNodeModules()
	fmt.Println("Done")
}

func removePackageLock() {
	const filename = "./package-lock.json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("no package-lock.json file found")
		return
	}

	os.Remove(filename)
	fmt.Println("removed package-lock.json")
}

func removeNodeModules() {
	const filename = "./node_modules"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("no node_modules folder found")
		return
	}

	os.RemoveAll(filename)
	fmt.Println("removed node_modules folder")
}
