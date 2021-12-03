package main

import "fmt"

const arraySize int = 1000000
const signalLength int = 12

func main() {
	fmt.Println("Solving D3P1")
	fmt.Println(solveD3P1() == 4138664)
	fmt.Println("Solving D2P2")
	fmt.Println(solveD2P2() == 1741971043)
	fmt.Println("Solving D2P1")
	fmt.Println(solveD2P1() == 1746616)
	fmt.Println("Solving D1P2")
	fmt.Println(solveD1P2() == 1362)
	fmt.Println("Solving D1P1")
	fmt.Println(solveD1P1() == 1387)
}
