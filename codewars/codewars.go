package codewars

import (
	"fmt"
	"strconv"
	"strings"
)

func ReverseString(str string) string {
	sa := strings.Split(str, " ")
	sb := []string{}
	for _, s := range sa {
		var res string
		for _, r := range s {
			res = string(r) + res
		}
		sb = append(sb, res)
	}
	str = strings.Join(sb, " ")
	return str
}

func GetSum(a, b int) int {
	if a == b {
		return a
	}
	s := 0
	if a < b {
		for i := a; i <= b; i++ {
			s += i
		}
	} else {
		for i := b; i <= a; i++ {
			s += i
		}
	}

	return s
}

/*
Implementieren Sie eine Funktion, die zwei IPv4-Adressen empfängt und die Anzahl der Adressen zwischen ihnen zurückgibt (einschließlich der ersten, ausschließlich der letzten).

Alle Eingaben sind gültige IPv4-Adressen in Form von Zeichenfolgen. Die letzte Adresse ist immer größer als die erste.
{"1.", "10.0.0.0", "10.0.0.50", 50},
{"2.", "20.0.0.10", "20.0.1.0", 246},
20.0.0.10
20.0.1.0
*/
func IpsBetween(start, end string) int {

	sa := strings.Split(start, ".")
	se := strings.Split(end, ".")
	sia := convertStrToInt(sa)
	sie := convertStrToInt(se)
	fmt.Printf("%v\n", sia)
	fmt.Printf("%v\n", sie)
	var sum int
	for i := 0; i < 4; i++ {
		if sia[i] == sie[i] {
			continue
		}
		if sia[i] < sie[i] {
			sie[i] = sie[i] + 255*(3-i)
			fmt.Println(sie[i])
		}
		sum += (sia[i] - sie[i])
	}
	fmt.Println(sum * -1)
	return sum
}

func convertStrToInt(arr []string) []int {
	ia := []int{}
	for _, a := range arr {
		i, _ := strconv.ParseInt(a, 10, 32)
		ia = append(ia, int(i))
	}
	return ia
}

func NextBigger(n int) int {
	res := 0

	if !testEqualDidit(n) {
		return -1
	}
	res = reveseDigit(n)
	return res

}

func reveseDigit(digits int) int {
	str := fmt.Sprintf("%d", digits)

	pos := struct {
		actual   int
		next     int
		postiton int
	}{0, 0, 0}

	for i := 0; i < len(str); i++ {
		if pos.actual < int(str[i]) {
			pos.actual = int(str[i])
		}
		if (i + 1) < len(str) {
			pos.next = int(str[i+1])
		}
		if pos.actual < pos.next {
			pos.actual = pos.next
		}
	}
	fmt.Printf("%v\n", pos)
	rest := []string{}
	for i := 0; i < len(str); i++ {
		rest = append(rest, string(str[i]))
	}

	a := rest[pos.postiton-1]
	b := rest[pos.postiton]
	rest[pos.postiton] = a
	rest[pos.postiton-1] = b
	r, _ := strconv.ParseInt(strings.Join(rest, ""), 10, 32)
	return int(r)
}

func testEqualDidit(n int) bool {

	str := fmt.Sprintf("%d", n)
	test := false
	
	var first int
	for i := 0; i < len(str); i++ {
		rint, _ := strconv.ParseInt(string(str[i]), 10, 32)
		if i == 0 {
			first = int(rint)
			continue
		}
		test = first < int(rint)
	}

	return test
}
