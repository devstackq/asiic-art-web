package main

import "fmt"

//space, true, start index, a,

// func main() {

// 	for _, v := range os.Args[1] {

// 		if v >= 'A' && v <= 'M' || v >= 'a' && v <= 'm' {
// 			z01.PrintRune(v + 13)
// 		} else if v >= 'M' && v <= 'Z' || v >= 'm' && v <= 'z' {
// 			z01.PrintRune(v - 13)
// 		} else {
// 			z01.PrintRune(v)
// 		}
// 	}
// }

//tabmult

//atoi, itoa, func multiple
// func main() {

// 	val := os.Args[1:][0]

// 	num := Atoi(val)

// 	for i := 1; i <= 9; i++ {

// 		res := num * i

// 		idx := Itoa(i)
// 		finalRes := Itoa(res)

// 		for _, v := range idx {
// 			z01.PrintRune(v)
// 		}
// 		z01.PrintRune(' ')
// 		z01.PrintRune('x')
// 		z01.PrintRune(' ')
// 		for _, v := range val {
// 			z01.PrintRune(v)
// 		}

// 		z01.PrintRune(' ')
// 		z01.PrintRune('=')
// 		z01.PrintRune(' ')

// 		for _, v := range finalRes {
// 			z01.PrintRune(v)
// 		}

// 		z01.PrintRune('\n')

// 	}

// }

// func Atoi(s string) int {

// 	res := 0
// 	for _, v := range s {
// 		res = res*10 + int(v-48)
// 	}
// 	return res
// }

// func Itoa(n int) (res string) {

// 	for n != 0 {
// 		mod := n % 10
// 		res = string(mod+48) + res
// 		n = n / 10
// 	}
// 	return res
// }

//doop, fmt, atoi, itoa

// func main() {

// 	a := os.Args[1:][0]
// 	op := os.Args[1:][1]
// 	b := os.Args[1:][2]

// 	n1, _ := strconv.Atoi(a)
// 	n2, _ := strconv.Atoi(b)

// 	switch op {
// 	case "-":
// 		fmt.Println(minus(n1, n2))

// 	case "+":
// 		fmt.Println(plus(n1, n2))

// 	case "*":

// 		fmt.Println(multi(n1, n2))

// 	case "/":

// 		fmt.Println(devide(n1, n2))

// 	case "%":

// 		fmt.Println(modul(n1, n2))

// 	}
// }
// func minus(a, b int) int {
// 	return a - b
// }
// func plus(a, b int) int {
// 	return a + b
// }

// func multi(a, b int) int {
// 	return a * b
// }
// func devide(a, b int) int {
// 	return a / b
// }
// func modul(a, b int) int {
// 	return a % b
// }

// compare prog

// func main() {

// 	args := os.Args[1:]
// 	defer z01.PrintRune('\n')
// 	if len(args) != 2 {
// 		z01.PrintRune('\n')
// 		return
// 	} else {

// 		Compare(args[0], args[1])

// 	}

// }

// func Compare(arg1, arg2 string) {

// 	if arg1 > arg2 {
// 		z01.PrintRune('1')
// 	}
// 	if arg1 < arg2 {
// 		z01.PrintRune('-')
// 		z01.PrintRune('1')
// 	}
// 	if arg1 == arg2 {
// 		z01.PrintRune('0')
// 	}

// }

//alphamirror

// func main() {

// 	if len(os.Args[1:]) < 1 {
// 		z01.PrintRune('\n')
// 		return
// 	} else {

// 		for _, v := range os.Args[1] {
// 			if v >= 'a' && v <= 'z' {
// 				z01.PrintRune('a' + 'z' - rune(v))
// 			} else if v >= 'A' && v <= 'Z' {
// 				z01.PrintRune('A' + 'Z' - rune(v))
// 			} else {
// 				z01.PrintRune(v)
// 			}
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

//last word
// func main() {

// 	//find start index, and end index, and print
// 	args := os.Args[1]
// 	end := 0
// 	startEnd := 0

// 	for i := len(args) - 1; i > 0; i-- {
// 		if args[i] != ' ' {
// 			end = i
// 			break
// 		}
// 	}

// 	for i := len(args) - 1; i > 0; i-- {
// 		if args[i] == ' ' {
// 			startEnd = i
// 			break
// 		}
// 	}
// 	for _, v := range args[startEnd+1 : end+1] {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }

//

//ispowerOF2

// func main() {

// 	args := os.Args[1]

// 	t := "true"
// 	f := "false"
// 	nb, _ := strconv.Atoi(args)

// 	if len(os.Args[1:]) != 1 {
// 		z01.PrintRune('\n')
// 		return
// 	} else {
// 		for nb > 1 {
// 			if nb%2 == 1 {
// 				for _, v := range f {
// 					z01.PrintRune(v)
// 				}
// 				z01.PrintRune('\n')
// 				return
// 			}
// 			nb = nb / 2
// 		}

// 		// if nb == 1 {
// 		// 	res = "true"
// 		// }

// 		for _, v := range t {
// 			z01.PrintRune(v)
// 		}
// 		z01.PrintRune('\n')

// 	}
// }

//Union

// func main() {

//	из 2 слов, по 1 уникальной букве,
//и выводит их результат
// 	if len(os.Args[1:]) <= 1 {
// 		z01.PrintRune('\n')
// 		return
// 	} else {
// 		acc := os.Args[1] + os.Args[2]
// 		var temp []rune
// 		for _, v := range acc {
// 			if isUnique(temp, v) {
// 				temp = append(temp, v)
// 			}
// 		}
// 		for _, v := range temp {
// 			z01.PrintRune(v)
// 		}

// 	}
// }

// func isUnique(args []rune, r rune) bool {
// 	for _, v := range args {
// 		if v == r {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {

// 	nb, _ := strconv.Atoi(os.Args[1])
// 	res := ""
// 	if nb > 0 {
// 		if nb%2 == 0 {
// 			nb = nb / 10
// 			res = "true"
// 		} else {
// 			res = "false"
// 		}
// 	}
// 	fmt.Print(res)
// }

//inter
// func main() {
// 	defer z01.PrintRune(10)
// 	// 1 acc, отбрасывает ненужные буквы, из второго слова, если есть совпадения, собирает их в acc
// 	acc := ""
// 	for _, v := range os.Args[1] {
// 		for _, l := range os.Args[2] {
// 			if l == v {
// 				acc += string(v)
// 				break
// 			}
// 		}
// 	}

// 	temp := []rune{}
// 	//check
// 	for _, v := range acc {
// 		if isUnique(temp, v) {
// 			temp = append(temp, v)
// 		}
// 	}
// 	for _, v := range temp {
// 		z01.PrintRune(v)
// 	}
// }
// func isUnique(arg []rune, r rune) bool {

// 	for _, v := range arg {
// 		if v == r {
// 			return false
// 		}
// 	}
// 	return true
// }

// padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
//robottoorigin

// 4 count == len(L), len(R)

// func main() {
// 	a := os.Args[1]

// 	defer z01.PrintRune('\n')
// 	var l, r, u, d int
// 	res := ""

// 	for _, v := range a {
// 		if v == 'L' {
// 			l++
// 		}
// 		if v == 'D' {
// 			d++
// 		}
// 		if v == 'R' {
// 			r++
// 		}
// 		if v == 'U' {
// 			u++
// 		}

// 	}
// 	//fmt.Print(l, r)
// 	if u == d && l == r {
// 		res = "true"
// 	} else {
// 		res = "false"
// 	}
// 	for _, v := range res {
// 		z01.PrintRune(v)
// 	}

// }

//printchessboard

// func main() {

// 	x := Atoi(os.Args[1])
// 	y := Atoi(os.Args[2])
// 	//fmt.Print(nb, nb2)

// 	for i := 0; i < y; i++ {
// 		for j := 0; j < x; j++ {
// 			if i%2 == 0 { //1
// 				if j%2 == 0 { //1.1
// 					z01.PrintRune('$')
// 				} else { //1.1
// 					z01.PrintRune(' ')
// 				}
// 			} else if i%2 == 1 { //2
// 				if j%2 == 1 { //2.1
// 					z01.PrintRune('$')
// 				} else { //2.2
// 					z01.PrintRune(' ')
// 				}
// 			}
// 		}
// 		z01.PrintRune('\n')
// 		//2 for, odd, even, print value
// 	}
// }

// func Atoi(s string) int {
// 	res := 0
// 	for _, v := range s {
// 		res = res*10 + int(v-48)
// 	}
// 	return res
// }

//swapbits

// func main() {
// 	res := SwapBits(65)
// 	fmt.Print(res)
// }

// func SwapBits(octet byte) byte {
// 	//return (octet / 16) + (octet % 16 * 16)
// 	return octet<<4 + octet>>4
// }

// //twosum

// func main() {
// 	case1 := []int{1, 2, 3, 4}
// 	out := TwoSum(case1, 5)
// 	fmt.Println(out)
// }
// func TwoSum(nums []int, target int) []int {

// 	res := []int{0, 0}

// 	for i, n1 := range nums {
// 		for j, n2 := range nums {
// 			if n1+n2 == target && i != j {
// 				res[0] = i
// 				res[1] = j
// 			}
// 		}
// 	}
// 	return res
// }

// priorprime

func main() {
	fmt.Println(Priorprime(14))

	//find prime nums
	// create array [],int, put value
	// sum + each elems, results
}

func Priorprime(x int) int {
	sum := 0
	for i := x; i > 0; i-- {
		if isPrime(i) {
			sum = sum + i
		}
	}
	return sum
}


func isPrime(n int) bool {

	flag := true

	if n <= 1 {
		flag = false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			flag = false
		}
	}
	return flag
}


Aigerim	 Kydyrbekova	 kydyrbekova +
 Abdubek	 Zholay	 Abdubek +
 Kazybek	 Suieubek	 Kazybeksuieubek +
 Yeldar	 Mauletov	 Yeldarml + 
Timur	 Nurgaliyev	 timurnurgali - MN
Alexey	 Cheremissov	 alnikolaevich - audio call



Elmira	 Aliyeva - justauser123 - ignor
Zhuldyz	 Namazbayeva	 znamazbayeva - ignor

Dias	 Sadykov	 DiasSadykov - not found
karimzhumazhanov - not found

Amina	 Shaikym	 IamhappyXD - left
Yevhenii	 Zhovnuvatyi	 ezhovnuvaty - left



//repeat alpha
// n = int(26 - (122 - v))

//5  -> ascii-web

//6,7 level - task solve