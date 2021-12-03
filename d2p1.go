package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getCommand(s string) (command string, length int64) {
	v := strings.Fields(s)

	val, err := strconv.ParseInt(v[1], 10, 64)
	if err != nil {
		fmt.Println(err, "err parsing string", v[1])
	}

	return v[0], val
}

func solveD2P1() int64 {
	var arr [1000000]string = readFile("directions.file")

	var xPos int64 = 0
	var zPos int64 = 0
	for i := 0; i < 1000000; i++ {
		if arr[i] == "" {
			break
		}
		command, distance := getCommand(arr[i])

		switch command {
		case "forward":
			xPos += distance
		case "down":
			zPos += distance
		case "up":
			zPos -= distance
		default:
			fmt.Println("Unknown command found", command, "on line", i)
		}
	}

	return xPos * zPos
}
