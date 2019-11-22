package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 10)
	for i := range a {
		a[i] = rand.Intn(10)
	}
	printSlice(a)

	//2 Добавить в конец слайса число 5
	a = append(a, 5)
	printSlice(a)

	//3 Добавить в начало слайса число 5
	f := []int{5}
	f = append(f, a...)
	a = f
	printSlice(a)

	//4 Взять последнее число слайса, вернуть его пользователю, а из слайса этот элемент удалить
	fmt.Println(a[len(a)-1])
	a = a[:len(a)-1]
	printSlice(a)

	//5 Взять первое число слайса, вернуть его пользователю, а из слайса этот элемент удалить
	fmt.Println(a[0])
	a = a[1:]
	printSlice(a)

	//6 Взять i-е число слайса, вернуть его пользователю, а из слайса этот элемент удалить. Число i передает пользователь в функцию
	a = del(2, a)
	printSlice(a)

	//7 Объединить два слайса и вернуть новый со всеми элементами первого и второго
	b := make([]int, 10)
	for i := range b {
		b[i] = rand.Intn(10)
	}
	printSlice(b)
	c := []int{}
	c = append(c, a...)
	c = append(c, b...)
	printSlice(c)

	//8 Из первого слайса удалить все числа, которые есть во втором
	//a := delel(a, b) если надо из а удалить и зменить значения элементов слайса
	g := delel(a, b)
	printSlice(g)

	//9 Сдвинуть все элементы слайса на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним.
	a = moveLeft(a)
	printSlice(a)

	//10 Тоже, но сдвиг на заданное пользователем i
	a = moveLeftI(a, 2)
	printSlice(a)

	//11 Тоже, что 9, но сдвиг вправо
	a = moveRight(a)
	printSlice(a)

	//12 Тоже, но сдвиг на i
	a = moveRightI(a, 2)
	printSlice(a)

	//13 Вернуть пользователю копию переданного слайса
	//теперь в tmp можно менять значения и a[]int не повредится
	tmp := make([]int, len(a))
	copy(tmp, a)
	printSlice(tmp)

	//14 В слайсе поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...
	a = neighbouR(a)
	printSlice(a)

	//15 Упорядочить слайс в порядке: прямом
	a = sorT(a)
	printSlice(a)

	//обратном
	a = sorTBack(a)
	printSlice(a)

	//лексикографическом
	s := []string{"asdad", "gafgsaDF", "SDFGSFGSFG", "dfsdfad", "ddfssf", "QWQWQW"}
	//знаю вот такой вариант
	//sort.Strings(s)
	//от обычной пузырьковой моя отличается только типом
	s = sorABC(s)
	fmt.Println(s)

}

//функция вывода слайса
func printSlice(x []int) {
	fmt.Printf("%v   len=%d cap=%d\n", x, len(x), cap(x))
}

/*функция удаления i-го элемента
я сделал удаленеие i-го элемента, где a[1] - это первый элемент среза и т.д.
для удаления i-го элемента, где a[0] - это первый элемент необходимо
k := y[x:]
y = y[:x-1]*/
func del(x int, y []int) []int {
	k := y[x+1:]
	y = y[:x]
	y = append(y, k...)
	return y
}

//функция сдвига на 1 элемент влево
func moveLeft(a []int) []int {
	a = append(a, a[0])
	a = a[1:]

	//можно вот так бегать.
	/*z := a[0]
	  for i := range a {
	      if i == len(a)-1 {
	          a[i] = z
	      } else {
	          a[i] = a[i+1]
	      }
	  }*/
	return a
}

//фунския сдвига на i элементов влево
func moveLeftI(a []int, d int) []int {
	a1 := a[:d]
	a = append(a, a1...)
	a = a[d:]
	return a
}

//функция сдвига на 1 элемент вправо
func moveRight(a []int) []int {
	a1 := a[len(a)-1:]
	a = a[:len(a)-1]
	a1 = append(a1, a...)
	a = a1
	return a
}

//функция сдвига на i элементов вправо
func moveRightI(a []int, d int) []int {
	a1 := a[len(a)-d:]
	a = a[:len(a)-d]
	a1 = append(a1, a...)
	a = a1
	return a
}

//функция меняющая местами соседние элементы массива
func neighbouR(a []int) []int {
	for i := 0; i < len(a)-1; i += 2 {
		z := a[i]
		a[i] = a[i+1]
		a[i+1] = z
	}
	return a
}

//пузырьковая прямая сортировка
func sorT(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				c := a[i]
				a[i] = a[j]
				a[j] = c
			}
		}
	}
	return a
}

//пузырьковая обратная сортировка
func sorTBack(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] < a[j] {
				c := a[i]
				a[i] = a[j]
				a[j] = c
			}
		}
	}
	return a
}

//сортировка в лексикографическом порядке
func sorABC(a []string) []string {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				c := a[i]
				a[i] = a[j]
				a[j] = c
			}
		}
	}
	return a
}

//удаление из первого слайса элементов, которые есть во втором
func delel(a []int, b []int) []int {
	s := []int{}
	for i := range a {
		l := 0
		for j := range b {
			if a[i] == b[j] {
				l++
			}
		}
		if l == 0 {
			s = append(s, a[i])
			l = 0
		}
	}
	return s
}
