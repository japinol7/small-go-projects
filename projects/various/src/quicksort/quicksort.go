package main

import (
	"fmt"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Partition[T Ordered](arr []T, low, high int) int {
	lessThanOrEqual := func(a, b T) bool {
		switch v := any(a).(type) {
		case string:
			return v <= any(b).(string)
		case int:
			return v <= any(b).(int)
		case int8:
			return v <= any(b).(int8)
		case int16:
			return v <= any(b).(int16)
		case int32:
			return v <= any(b).(int32)
		case int64:
			return v <= any(b).(int64)
		case uint:
			return v <= any(b).(uint)
		case uint8:
			return v <= any(b).(uint8)
		case uint16:
			return v <= any(b).(uint16)
		case uint32:
			return v <= any(b).(uint32)
		case uint64:
			return v <= any(b).(uint64)
		case uintptr:
			return v <= any(b).(uintptr)
		case float32:
			return v <= any(b).(float32)
		case float64:
			return v <= any(b).(float64)
		default:
			panic("unsupported type")
		}
	}
	index := low - 1
	pivotElement := arr[high]
	for i := low; i < high; i++ {
		if lessThanOrEqual(arr[i], pivotElement) {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

// QuicksortRange sorts a range within the array
func QuicksortRange[T Ordered](arr []T, low, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := Partition(arr, low, high)
		QuicksortRange(arr, low, pivot-1)
		QuicksortRange(arr, pivot+1, high)
	}
}

// Quicksort sorts an array
func Quicksort[T Ordered](arr []T) []T {
	QuicksortRange(arr, 0, len(arr)-1)
	return arr
}

func main() {
	appName := "Quicksort"
	fmt.Println("Start app", appName, "\n")

	// Integer example
	numbers := []int{64, 34, -25, 14, 12, 22, 11, 90}
	fmt.Println("Original integers:", numbers)
	sortedNumbers := Quicksort(numbers)
	fmt.Println("Sorted integers:", sortedNumbers)

	// Float example
	floats := []float64{2.21, 3.44, 1.41, -0.75, 1.71, 0.51, 1.83}
	fmt.Println("\nOriginal floats:", floats)
	sortedFloats := Quicksort(floats)
	fmt.Println("Sorted floats:", sortedFloats)

	// String example
	words := []string{"banana", "eye", "reload", "holster", "arrow", "blue"}
	fmt.Println("\nOriginal strings:", words)
	sortedWords := Quicksort(words)
	fmt.Println("Sorted strings:", sortedWords)

	fmt.Println("\nEnd app", appName)
}
