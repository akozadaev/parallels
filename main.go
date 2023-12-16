package main

import (
	"fmt"
)

func printMessage(message string) {
	fmt.Println(message)
}

//func main() {
//	go printMessage("Привет, параллелизм! (Ожидаем это)")
//	fmt.Println("А может напечататься  ЭТО!.")
//	var a string
//	fmt.Scan(&a)
//}

func sumSlice(slice []int, result *int) {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	*result = sum
}


//func main() {
//	slice1 := []int{1, 2, 3, 4} //10
//	slice2 := []int{6, 7, 8, 9, 10} //40
//	sharedResult := 0
//	go sumSlice(slice1, &sharedResult)
//	go sumSlice(slice2, &sharedResult)
//	time.Sleep(1 * time.Second)
//	fmt.Println("Result:", sharedResult)
//}

func sumSliceN(slice []int, ch chan int) {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	ch <- sum
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{6, 7, 8, 9, 10}
	ch := make(chan int)
	go sumSliceN(slice1, ch)
	go sumSliceN(slice2, ch)
	result1 := <-ch
	result2 := <-ch
	close(ch)
	fmt.Println("Результат:", result1+result2)
}

// Небуферизованный канал ch := make(chan int) // Буферизованный канал с емкостью 5 bufCh := make(chan int, 5)
