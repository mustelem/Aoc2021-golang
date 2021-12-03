package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(filename string) [arraySize]string {
	var arr [arraySize]string

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	idx := 0
	for scanner.Scan() {
		arr[idx] = scanner.Text()
		idx++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err, "reading file", file.Name())
	}

	return arr
}

func getInt(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("converting value to int")
	}
	return val
}

func solveD1P1() int64 {
	var arr [arraySize]string = readFile("depth.file")

	var count int64 = 0
	for i := 1; i < arraySize; i++ {
		if arr[i] == "" {
			break
		}

		if getInt(arr[i]) > getInt(arr[i-1]) {
			count++
		}
	}

	return count
}
