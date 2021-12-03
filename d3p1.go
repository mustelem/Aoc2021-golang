package main

import (
	"strconv"
)

func getZerosOnes(arr [arraySize]string) ([signalLength]int, [signalLength]int) {
	var zeros [signalLength]int
	var ones [signalLength]int

	for i := 0; i < arraySize; i++ {
		if arr[i] == "" {
			break
		}
		for j := 0; j < signalLength; j++ {
			if string(arr[i][j]) == "0" {
				zeros[j]++
			} else {
				ones[j]++
			}
		}
	}

	return zeros, ones
}

func getDec(bin string) int64 {

	retval, err := strconv.ParseInt(bin, 2, 64)

	if err != nil {
		panic(err)
	}

	return retval
}

func solveD3P1() int64 {
	var gammaBin, epsilonBin string

	var arr [arraySize]string = readFile("diagnostic.file")
	zeros, ones := getZerosOnes(arr)
	for i := 0; i < signalLength; i++ {
		if zeros[i] > ones[i] {
			gammaBin += "0"
			epsilonBin += "1"
		} else {
			gammaBin += "1"
			epsilonBin += "0"
		}
	}

	return getDec(gammaBin) * getDec(epsilonBin)
}
