package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var maxRandomValue = 10
var width = 6000
var height = 6000

func main() {
	//fmt.Print("Enter w and h: ")
	//fmt.Scanf("%d", &w)
	//fmt.Scanf("%d", &h)
	//fmt.Println("You input ", w, " ", h)
	for i := 0; i < 3; i++ {
		matrix := getVector(width*height)
		vector := getVector(width)

		ts0 := time.Now()
		resultSerial := SerialMatrixVectorMultiplication(matrix, vector, width, height)
		ts1 := time.Now()
		fmt.Println("SerialMatrixVectorMultiplication time: ", ts1.Sub(ts0))

		tps0 := time.Now()
		resultParallelString := ParallelMatrixVectorStringMultiplication(matrix, vector, width, height)
		tps1 := time.Now()
		fmt.Println("ParallelMatrixVectorStringMultiplication time: ", tps1.Sub(tps0))

		tpc0 := time.Now()
		resultParallelColumn := ParallelMatrixVectorColumnMultiplication(matrix, vector, width, height)
		tpc1 := time.Now()
		fmt.Println("ParallelMatrixVectorColumnMultiplication time: ", tpc1.Sub(tpc0))
		//
		//tpc02 := time.Now()
		//ParallelColumnMatrixVectorMultiplication2()
		//tpc12 := time.Now()
		//fmt.Println("ParallelColumnMatrixVectorMultiplication2 time: ", tpc12.Sub(tpc02))

		fmt.Println("equal2? ", Equal2(resultParallelString, resultSerial))
		fmt.Println("equal3? ", Equal3(resultParallelString, resultSerial, resultParallelColumn))
	}
}


func ParallelMatrixVectorColumnMultiplication(matrix []int, vector []int, width int, height int) []int {
	result := make([]int, len(matrix))
	//width := len(matrix[0])
	var wg sync.WaitGroup
	wg.Add(height)
	for i := 0; i < width; i++ {
		go VectorVectorMultiplication2(i, width, height, matrix, vector[i], &result, &wg)
	}
	wg.Wait()
	return result
}

func VectorVectorMultiplication2(columnNumber, width, height int, vector []int, number int, result *[]int, wg *sync.WaitGroup) {
	for i := 0; i < height; i++ {
		(*result)[i] += vector[columnNumber+width*i]*number
	}
	wg.Done()
}



func SerialMatrixVectorMultiplication(matrix []int, vector []int, width int, height int) []int {
	result := make([]int, len(matrix))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			result[i] += matrix[i*height+j]*vector[j]
		}
	}
	return result
}

func ParallelMatrixVectorStringMultiplication(matrix []int, vector []int, width int, height int) []int{
	result := make([]int, len(matrix))
	var wg sync.WaitGroup
	wg.Add(height)
	for i := 0; i < height; i++ {
		go VectorVectorMultiplication(i, width, height, matrix, vector, &result[i], &wg)
	}
	wg.Wait()
	return result
}

func VectorVectorMultiplication(stringNumber, width, height int, vector1 []int, vector2 []int, result *int, wg *sync.WaitGroup) {
	for i := 0; i < width; i++ {
		*result += vector1[stringNumber*height+i]*vector2[i]
	}
	wg.Done()
}

func getVector(len int) []int {
	vector := make ([]int, len)
	for i := 0; i < len; i++ {
		vector[i] = rand.Intn(maxRandomValue)
	}
	//fmt.Println(vector)
	return vector
}


func Equal3(a, b , c[]int) bool {
	return Equal2(a,b) && Equal2(a,c)
	//if len(a) != len(b) && len(a) != len(c) {
	//	return false
	//}
	//for i, v := range a {
	//	if v != b[i] || v != c[i] {
	//		return false
	//	}
	//}
	//return true
}

func Equal2(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}