package main

import (
	"fmt"
	"strconv"
)

func main() {
	// b := bill.Bill{
	// 	Text:     "Hallo",
	// 	Amount:   12.3,
	// 	Ballance: 1235.5896,
	// }
	// fmt.Println(b.ToString())
	// fmt.Println("Hier ist die zweite Meldung")

	//println(codewars.NextBigger(21))
	// test := 59884848459853
	test2 := 1891

	// println(test, ":", z(test))
	println(test2, ":", z(test2))
	//println(codewars.NextBigger(43))
	//println(codewars.NextBigger(34))

	fmt.Println()
}

func z(input int) int {

	sinput := fmt.Sprintf("%d", input)
	arrInput := []int{}
	for i := 0; i < len(sinput); i++ {
		is, _ := strconv.ParseInt(string(sinput[i]), 10, 32)
		arrInput = append(arrInput, int(is))
	}
	fmt.Printf("\nInput Array:\t%v\n", arrInput)

	max := struct {
		max, next, pos int
	}{0, 0, 0}

	for i := 0; i < len(arrInput); i++ {
		if max.max < arrInput[i] || max.next < arrInput[i] {
			max.max = arrInput[i]
			max.pos = i
			if i < len(arrInput) {
				max.next = arrInput[i+1]
				max.pos = i
			}
		}
		fmt.Println(max.pos, i)
	}
	fmt.Printf("\nMaximaler Wert:\t%d\n\n", max)
	if max.pos == 0 {
		return -1
	}
	a := arrInput[max.pos-1]
	b := arrInput[max.pos]

	arrInput[max.pos] = a
	arrInput[max.pos-1] = b

	fmt.Printf("Ouput Array:\t%v\n\n", arrInput)
	var str string
	for _, s := range arrInput {
		str += fmt.Sprint(s)
	}
	res, _ := strconv.ParseInt(str, 10, 32)

	return int(res)
}
