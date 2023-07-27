package utils

import (
	"errors"
	"fmt"
	"strconv"
)

func Sorter(array []string) []string {
	key := array[0]
	switch key {
	case "-string":
		return sortStrings(array[1:])
	case "-int":
		sortedArr, err := sortInts(array[1:])
		if err != nil {
			fmt.Println("Error:", err)
		}
		return sortedArr
	case "-mix":
		return sortMixedArray(array[1:])
	default:
		return array
	}
}

func sortStrings(strs []string) []string {
	length := len(strs)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if strs[i] > strs[j] {
				strs[i], strs[j] = strs[j], strs[i]
			}
		}
	}
	return strs
}

func sortInts(ints []string) ([]string, error) {
	length := len(ints)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			a, err1 := strconv.Atoi(ints[i])
			if err1 != nil {
				return nil, errors.New("could not convert string to integer")
			}
			b, err2 := strconv.Atoi(ints[j])
			if err2 != nil {
				return nil, errors.New("could not convert string to integer")
			}
			if a > b {
				ints[i], ints[j] = ints[j], ints[i]
			}
		}
	}
	return ints, nil
}

func sortMixedArray(arr []string) []string {
	ints := make([]string, 0)
	strs := make([]string, 0)
	for _, str := range arr {
		_, err := strconv.Atoi(str)
		if err == nil {
			ints = append(ints, str)
		} else {
			strs = append(strs, str)
		}
	}

	sortedInts, _ := sortInts(ints)
	sortedStrings := sortStrings(strs)

	return append(sortedInts, sortedStrings...)
}
