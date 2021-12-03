package main

import "fmt"

func solveD2P2() int64 {
	var arr [1000000]string = readFile("directions.file")

	var xPos int64 = 0
	var zPos int64 = 0
	var aim int64 = 0
	for i := 0; i < 1000000; i++ {
		if arr[i] == "" {
			break
		}
		command, distance := getCommand(arr[i])

		switch command {
		case "forward":
			xPos += distance
			zPos += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		default:
			fmt.Println("Unknown command found", command, "on line", i)
		}
	}

	return xPos * zPos
}
