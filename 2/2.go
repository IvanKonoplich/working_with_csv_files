package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

//Составить список учебной группы, включающий Х человек. Для каждого студента указать: фамилию и имя,
//дату рождения (год, месяц и число), также у каждого студента есть зачетка, в которой указаны:
//предмет (от трех до пяти), дата экзамена, ФИО преподавателя.
//чтение из терминала или файла или бд
//Вывести в виде таблицы все предметы и количество студентов, получивших оценки по этим предметам

type Class struct { // Our example struct, you can use "-" to ignore a field
	ClassName string `csv:"предмет"`
	Exam      int    `csv:"оценка"`
}

func main() {
	classesFile, err := os.OpenFile("classes.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer classesFile.Close()

	classes := []*Class{}

	if err := gocsv.UnmarshalFile(classesFile, &classes); err != nil { // Load clients from file
		panic(err)
	}
	result := make(map[string]int, len(classes))
	for _, class := range classes {
		if class.Exam > 0 {
			result[class.ClassName]++
		}
	}
	for key, value := range result {
		fmt.Printf("%s:%d\n", key, value)
	}
}
