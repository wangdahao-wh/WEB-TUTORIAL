package main

import "fmt"

func main() {
    height := []int{1,2,3,4,5}
	fmt.Println(stableMountains(height, 2))
}

func stableMountains(height []int, threshold int) []int {
    len := len(height)
    res := make([]int, 0)
    var num int = 0
    for i := 1; i < len; i++ {
        if (height[i-1] > threshold) {
            res = append(res, i)
            num++
        }
    }
	fmt.Println(num)
    return res
}
