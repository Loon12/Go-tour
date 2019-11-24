package main

import (
	"fmt"
	"math/rand"
	"sort"
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
	Plus5end(&a)
	printSlice(a)

	//3 Добавить в начало слайса число 5
	Plus5begin(&a)
	printSlice(a)

	//4 Взять последнее число слайса, вернуть его пользователю, а из слайса этот элемент удалить
	delEnd(&a)
	printSlice(a)

	//5 Взять первое число слайса, вернуть его пользователю, а из слайса этот элемент удалить
	delBegin(&a)
	printSlice(a)

	//6 Взять i-е число слайса, вернуть его пользователю, а из слайса этот элемент удалить. Число i передает пользователь в функцию
	del(2, &a)
	printSlice(a)

	//7 Объединить два слайса и вернуть новый со всеми элементами первого и второго
	b := make([]int, 10)
	for i := range b {
		b[i] = rand.Intn(10)
	}
	printSlice(b)
	merge(&a, &b)
	printSlice(a)

	//8 Из первого слайса удалить все числа, которые есть во втором
	//a := delel(a, b) если надо из а удалить и зменить значения элементов слайса
	delel(&b, &a)
	printSlice(b)

	//9 Сдвинуть все элементы слайса на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним.
	moveLeft(&a)
	printSlice(a)

	//10 Тоже, но сдвиг на заданное пользователем i
	moveLeftI(&a, 2)
	printSlice(a)

	//11 Тоже, что 9, но сдвиг вправо
	moveRight(&a)
	printSlice(a)

	//12 Тоже, но сдвиг на i
	moveRightI(&a, 2)
	printSlice(a)

	//13 Вернуть пользователю копию переданного слайса
	//теперь в tmp можно менять значения и a[]int не повредится
	tmp := make([]int, len(a))
	copy(tmp, a)
	//printSlice(tmp)

	//14 В слайсе поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...
	neighbouR(&a)
	printSlice(a)

	//15 Упорядочить слайс в порядке: прямом
	sorT(&a)
	printSlice(a)

	//обратном
	sorTBack(&a)
	printSlice(a)

	//лексикографическом
	s := []string{"asdad", "gafgsaDF", "SDFGSFGSFG", "dfsdfad", "ddfssf", "QWQWQW", "1234", "12335"}
	sorABC(s)
	fmt.Println(s)

}

//функция вывода слайса
func printSlice(x []int) {
	fmt.Printf("%v   len=%d cap=%d\n", x, len(x), cap(x))
}

//Plus5end функция добавления числа 5 в конец слайса
func Plus5end(x *[]int) {
	*x = append(*x, 5)
}

//Plus5begin функция добавляет 5 в начало слайса
func Plus5begin(x *[]int) {
	f := []int{5}
	f = append(f, *x...)
	*x = f
}

//функция возврата и удаления последнего элемента слайса
func delEnd(a *[]int) {
	fmt.Println((*a)[len(*a)-1])
	*a = (*a)[:len(*a)-1]
}

//функция возвората и удаления первого элемента слайса
func delBegin(a *[]int) {
	fmt.Println((*a)[0])
	*a = (*a)[1:]
}

//функция объединения слайсов
func merge(a *[]int, b *[]int) {
	*a = append(*a, *b...)
}

/*функция удаления i-го элемента
я сделал удаленеие i-го элемента, где a[1] - это первый элемент среза и т.д.
для удаления i-го элемента, где a[0] - это первый элемент необходимо
k := y[x:]
y = y[:x-1]*/
func del(x int, y *[]int) {
	*y = append((*y)[:x], (*y)[x+1:]...)
}

//функция сдвига на 1 элемент влево
func moveLeft(a *[]int) {
	*a = append(*a, (*a)[0])
	*a = (*a)[1:]

	//можно вот так бегать.
	/*z := a[0]
	  for i := range a {
	      if i == len(a)-1 {
	          a[i] = z
	      } else {
	          a[i] = a[i+1]
	      }
	  }*/
}

//фунския сдвига на i элементов влево
func moveLeftI(a *[]int, d int) {
	a1 := (*a)[:d]
	*a = append(*a, a1...)
	*a = (*a)[d:]
}

//функция сдвига на 1 элемент вправо
func moveRight(a *[]int) {
	a1 := (*a)[len(*a)-1:]
	*a = (*a)[:len(*a)-1]
	*a = append(a1, (*a)...)
}

//функция сдвига на i элементов вправо
func moveRightI(a *[]int, d int) {
	a1 := (*a)[len(*a)-d:]
	*a = (*a)[:len(*a)-d]
	*a = append(a1, (*a)...)
}

//функция меняющая местами соседние элементы массива
func neighbouR(a *[]int) {
	for i := 0; i < len(*a)-1; i += 2 {
		(*a)[i], (*a)[i+1] = (*a)[i+1], (*a)[i]
	}
}

//пузырьковая прямая сортировка
func sorT(a *[]int) {
	sort.Ints(*a)
}

//пузырьковая обратная сортировка
func sorTBack(a *[]int) {
	sort.Ints(*a)
	for i := 0; i < len(*a)/2; i++ {
		(*a)[i], (*a)[len(*a)-i-1] = (*a)[len(*a)-i-1], (*a)[i]
	}
}

//сортировка в лексикографическом порядке
func sorABC(a []string) {
	sort.Strings(a)
}

func delel(a *[]int, b *[]int) {
	for i := 0; i < len(*a); i++ {
		for j := 0; j < len(*b); j++ {
			if (*a)[i] == (*b)[j] {
				if i == len(*a)-1 {
					*a = (*a)[:i]
					break
				} else {
					*a = append((*a)[:i], (*a)[i+1:]...)
					j = -1
				}
			}
		}
	}
}
