package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func exit() {
	clear()
	fmt.Println("Goodbye")
	os.Exit(0)
}

func clear() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Portfolio Manager v" + version + " by " + author)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
}

func input(variable interface{},prompt ...string) {
	if prompt != nil {
		fmt.Println(prompt[0])
	}
	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	handleError(err)
	switch variable.(type) {
	case *string:
		*variable.(*string) = line[:len(line)-1]
	case *int:
		if len(line) == 1 {
			*variable.(*int) = 0
		} else {
			*variable.(*int), err = strconv.Atoi(line[:len(line)-1])
			if(err != nil) {
				*variable.(*int) = 0
			}
		}
	}
}


func ensureDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		// Create config directory
		err = os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Println("Error creating directory: " + err.Error())
			os.Exit(1)
		}
	}
}

