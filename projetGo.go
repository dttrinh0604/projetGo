package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// function to read matrix from file name
func readMatrix(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, value := range strings.Split(line, " ") {
			num, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

//function to perform multiplication of 2 matrix given by readMatrix

func multiply(a, b [][]int, c chan [][]int) {
	rowsA := len(a)
	colsA := len(a[0])
	colsB := len(b[0])

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			sum := 0
			for k := 0; k < colsA; k++ {
				sum += a[i][k] * b[k][j]
			}
			result[i][j] = sum
		}
	}
	c <- result
}

func main() {
	a := readMatrix("matriceA.txt")
	b := readMatrix("matriceB.txt")
	c := make(chan [][]int)

	go multiply(a, b, c)
	result := <-c

	fmt.Println("Result: ", result)
}
