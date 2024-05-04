package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func worker(mode int32, arr [][]int32, c chan int32) {
	var sum int32
	if mode == 0 {
		for i := 0; i < len(arr); i++ {
			sum += arr[i][i]
		}
	} else {
		for i := 0; i < len(arr); i++ {
			sum += arr[i][len(arr)-1-i]
		}
	}
	c <- sum
}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var res int32
	c := make(chan int32)
	go worker(0, arr, c)
	go worker(1, arr, c)
	primaryDiagonal, secondaryDiagonal := <-c, <-c

	res = primaryDiagonal - secondaryDiagonal
	fmt.Println("sdfajsdfk", primaryDiagonal, secondaryDiagonal, res)
	if res < 0 {
		return -res
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
