package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches   []string
	waitGroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func fileSearch(root, fileName string) {
	defer waitGroup.Done() // Mark this goroutine as done when it exits.
	fmt.Println("Searching in", root)
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if strings.Contains(file.Name(), fileName) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			waitGroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), fileName)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Name Of File : ")
	str1, _ := reader.ReadString('\n')
	str1 = strings.TrimSpace(str1)

	fmt.Print("Enter the Root Path: ")
	str2, _ := reader.ReadString('\n')
	str2 = strings.TrimSpace(str2)
	fmt.Println(str2, str1)
	waitGroup.Add(1)
	go fileSearch(str2, str1)
	waitGroup.Wait() // Wait for all goroutines to complete.
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
	fmt.Println("Count of Matched file:", len(matches))
}
