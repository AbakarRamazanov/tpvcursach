package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var maxRandomValue = 10

func getMatrix(w int, h int) [][]int {
	matrix := make([][]int, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]int, w)
		for j := 0; j < w; j++ {
			matrix[i][j] = rand.Intn(maxRandomValue)
		}
	}
	//fmt.Println(matrix)
	return matrix
}

func getVector(len int) []int {
	vector := make ([]int, len)
	for i := 0; i < len; i++ {
		vector[i] = rand.Intn(maxRandomValue)
	}
	//fmt.Println(vector)
	return vector
}

func getMatrixAndVector(w int, h int) ([][]int, []int) {
	return getMatrix(w,h), getVector(h)
}

func SerialMatrixVectorMultiplication(matrix [][]int, vector []int) []int {
	result := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[i] += matrix[i][j]*vector[j]
		}
	}
	return result
}

func VectorNumberMultiplication(vector []int, number int, result *int, wg *sync.WaitGroup) {
	for i := 0; i < len(vector); i++ {
		*result += vector[i]*number
	}
	wg.Done()
}

func VectorVectorMultiplication(vector1 []int, vector2 []int, result *int, wg *sync.WaitGroup) {
	for i := 0; i < len(vector1); i++ {
		*result += vector1[i]*vector2[i]
	}
	wg.Done()
}

func ParallelStringMatrixVectorMultiplication(matrix [][]int, vector []int) []int{
	result := make([]int, len(matrix))
	var wg sync.WaitGroup
	i := 0
	for i = 0; i < len(matrix); i++ {
		wg.Add(1)
		go VectorVectorMultiplication(matrix[i], vector, &result[i], &wg)
	}
	wg.Wait()
	return result
}

	func NumberNumberMultiplication(vectorNumber int, matrixNumber int, result *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	*result += vectorNumber * matrixNumber
	mutex.Unlock()
	wg.Done()
}

func ParallelColumnMatrixVectorMultiplication(matrix [][]int, vector []int) []int{
	result := make([]int, len(matrix))
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < len(matrix); i++ {
		wg.Add(len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			go NumberNumberMultiplication(vector[i], matrix[i][j], &result[i], &wg, &mutex)
		}
	}
	wg.Wait()
	return result
}

var result2 []int
var matrix2 [][]int
var vector2 []int
var mutex2 []sync.Mutex
var wg2 sync.WaitGroup

func ParallelColumnMatrixVectorMultiplication2(){
	for i := 0; i < len(matrix2); i++ {
		wg2.Add(len(matrix2[i]))
		for j := 0; j < len(matrix2[i]); j++ {
			go NumberNumberMultiplication2(i,j)
		}
	}
	wg2.Wait()
}

func NumberNumberMultiplication2(i int, j int) {
	mutex2[i].Lock()
	result2[i] += vector2[i] * matrix2[i][j]
	mutex2[i].Unlock()
	wg2.Done()
}

func Equal3(a, b , c[]int) bool {
	if len(a) != len(b) && len(a) != len(c) {
		return false
	}
	for i, v := range a {
		if v != b[i] || v != c[i] {
			return false
		}
	}
	return true
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

func main() {
	var w, h int
	//fmt.Print("Enter w and h: ")
	//fmt.Scanf("%d", &w)
	//fmt.Scanf("%d", &h)
	//fmt.Println("You input ", w, " ", h)
	w, h = 1000, 1000
	mutex2 = make([]sync.Mutex, w)
	result2 = make([]int, w)
	for i := 0; i < 3; i++ {
		matrix2 = getMatrix(w, h)
		vector2 = getVector(h)
		matrix := matrix2
		vector := vector2


		ts0 := time.Now()
		resultSerial := SerialMatrixVectorMultiplication(matrix, vector)
		ts1 := time.Now()
		fmt.Println("SerialMatrixVectorMultiplication time: ", ts1.Sub(ts0))

		tps0 := time.Now()
		resultParallelString := ParallelStringMatrixVectorMultiplication(matrix, vector)
		tps1 := time.Now()
		fmt.Println("ParallelStringMatrixVectorMultiplication time: ", tps1.Sub(tps0))

		//tpc0 := time.Now()
		//resultParallelColumn := ParallelColumnMatrixVectorMultiplication(matrix, vector)
		//tpc1 := time.Now()
		//fmt.Println("ParallelColumnMatrixVectorMultiplication time: ", tpc1.Sub(tpc0))
		//
		//tpc02 := time.Now()
		//ParallelColumnMatrixVectorMultiplication2()
		//tpc12 := time.Now()
		//fmt.Println("ParallelColumnMatrixVectorMultiplication2 time: ", tpc12.Sub(tpc02))



		fmt.Println("equal2? ", Equal2(resultParallelString, resultSerial))
		//fmt.Println("equal3? ", Equal3(resultParallelString, resultSerial, resultParallelColumn))
	}
}