package main

func getLastFour(a [arraySize]string, idx int) (int64, int64, int64, int64) {
	return getInt(a[idx]), getInt(a[idx-1]), getInt(a[idx-2]), getInt(a[idx-3])
}

func solveD1P2() int64 {
	var arr [arraySize]string = readFile("depth.file")

	var count int64 = 0
	for i := 3; i < arraySize; i++ {
		if arr[i] == "" {
			break
		}
		last, prev, nextprev, oldest := getLastFour(arr, i)
		if last+prev+nextprev > prev+nextprev+oldest {
			count++
		}
	}

	return count
}
