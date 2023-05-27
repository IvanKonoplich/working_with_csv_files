package main

import (
	"fmt"
	"log"
)

//Напишите функцию, которая начинает, записывает с определенной позиции,
//в первый массив, все значения второго (без использования стандартной библиотеки)

func main() {
	fmt.Println("введите колличество элементов первого массива")
	var num int
	fmt.Scan(&num)
	fmt.Println("введите элементы первого массива")
	var elem string
	first := make([]string, 0, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&elem)
		first = append(first, elem)
	}
	fmt.Println("введите колличество элементов второго массива")
	fmt.Scan(&num)
	fmt.Println("введите элементы второго массива")
	second := make([]string, 0, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&elem)
		second = append(second, elem)
	}
	fmt.Println("введите позицию(индекс начиная с 0) первого массива, с которой будут вставляться значения ")
	var point int
	fmt.Scan(&point)
	if point >= num {
		log.Fatal("позиция больше длины массива")
	}
	if point > num {
		log.Fatal("позиция меньше длины массива")
	}
	fmt.Println(insert(first, second, point))
}

func insert(f, s []string, point int) []string {
	result := make([]string, 0, len(f)+len(s))
	result = append(result, f[:point]...)
	result = append(result, s...)
	result = append(result, f[point:]...)
	return result
}

//какие то изменения в ветке dev
