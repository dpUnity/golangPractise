package practice

import (
	"fmt"
	"os"
	"strings"
)

func StartSlicePractice() {
	addSingleItemToSlice()
	addMultipleItemsToSlice()
	addSliceToSlice()
}

func addSingleItemToSlice() {
	fmt.Println("-----------新增單一元素到切片-----------")
	sliceOne := []string{}
	if len(os.Args) > 0 {
		for i := 0; i < len(os.Args); i++ {
			sliceOne = append(sliceOne, os.Args[i])
		}
	}
	fmt.Println("建立了切片：" + strings.Join(sliceOne, " "))
}

func addMultipleItemsToSlice() {
	fmt.Println("-----------新增多個元素到切片-----------")
	sliceOne := []int{}
	sliceOne = append(sliceOne, 1, 2, 3, 4, 5)
	fmt.Println("建立了切片：", sliceOne)
}

func addSliceToSlice() {
	fmt.Println("-----------新增切片到切片-----------")
	sliceOne := []int{1, 2, 3}
	sliceTwo := []int{}

	sliceTwo = append(sliceTwo, 9, 9)
	sliceTwo = append(sliceTwo, sliceOne...)
	fmt.Println("建立了切片：", sliceTwo)
}
