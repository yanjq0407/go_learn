package main

import "fmt"

func main() {
	// 示例切片
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slice)

	// 删除索引为 2 的元素（即元素 3）
	slice = removeElement(slice, 2)
	fmt.Println("Slice after removal:", slice)

	// 示例切片（字符串类型）
	stringSlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println("Original string slice:", stringSlice)

	// 删除索引为 2 的元素（即元素 "c"）
	stringSlice = removeElement(stringSlice, 2)
	fmt.Println("String slice after removal:", stringSlice)
}

// removeElement 泛型函数从切片中删除索引为 index 的元素并缩容
func removeElement[T any](slice []T, i int) []T {
	// 创建一个新的切片，容量为当前切片长度减一
	newSlice := make([]T, 0, len(slice)-1)

	// 将要删除元素之前的元素复制到新切片
	newSlice = append(newSlice, slice[:i]...)

	// 将要删除元素之后的元素复制到新切片
	newSlice = append(newSlice, slice[i+1:]...)

	return newSlice
}
