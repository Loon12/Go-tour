//Сделать Фибоначчи с мемоизацией
//я погуглил что такое мемоизация и слеал как понял
//главная идея запоминать уже проделанную работу по вычислению конкретного элемента последовательнсти фибоначи
package main

import "fmt"

var memory = map[int]int{
	1: 1,
	2: 1,
	3: 2,
	4: 3,
}

func main() {

	var k int = 4
	fmt.Println(fibonachi(k))
	fmt.Println(fibonachi(6))

}

//Хз я чет не догнал как все сходу запилить в одну функцию
//Это вспомогательная классическая рекурсивная функция
func fibonachchi(i int) int {
	if i == 1 || i == 2 {
		return 1
	}
	return fibonachchi(i-1) + fibonachchi(i-2)
}

//Ну а это - основа
//есть в памяти нет таких вычислений указанного числа - то вычисляем и записываем в карту; иначе просто выводим известный результат
func fibonachi(k int) int {
	if _, l := memory[k]; l == false {
		memory[k] = fibonachchi(k)
		return memory[k]
	}
	return memory[k]
}
