package util

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

// TrimNumber ...
func TrimNumber(car string) string {
	re := regexp.MustCompile(`\s+`)

	car = re.ReplaceAllString(car, "")
	car = strings.ToUpper(car)

	re1 := regexp.MustCompile(`(?i:rus?)$`)
	re2 := regexp.MustCompile(`(?i:рус?)$`)

	car = re1.ReplaceAllString(car, "")
	car = re2.ReplaceAllString(car, "")

	return car
}

// TrimSpaces ...
func TrimSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// DigitsCount ...
func DigitsCount(number int64) int64 {
	count := int64(0)
	for number != 0 {
		number /= 10
		count++
	}
	return count

}

// SliceToInt ...
func SliceToInt(data []int64, to int) int64 {

	dlen := len(data)

	if to > dlen {
		to = dlen
	} else if to < 0 {
		to = 0
	}

	f := dlen - to
	d := data[to]

	for i := (to - 1); i >= 0; i-- {
		pow := dlen - (i + f)
		d += int64(math.Pow10(pow)) * data[i]
	}

	return d
}

// IntToSlice ...
func IntToSlice(n int64, sequence []int64) []int64 {
	if n != 0 {
		i := n % 10
		sequence = append([]int64{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

// CheckINN ...
func CheckINN(d int64) bool {
	innLen := DigitsCount(d)
	inn := IntToSlice(d, []int64{})

	if innLen != 10 && innLen != 12 {
		return false
	}

	if innLen == 10 {

		a := 2*inn[0] + 4*inn[1] + 10*inn[2] + 3*inn[3] +
			5*inn[4] + 9*inn[5] + 4*inn[6] + 6*inn[7] + 8*inn[8]

		sum := a % 11
		if sum > 9 {
			sum = sum % 10
		}
		check := inn[9]
		if check != sum {
			return false
		}

	} else if innLen == 12 {

		a := 7*inn[0] + 2*inn[1] + 4*inn[2] +
			10*inn[3] + 3*inn[4] + 5*inn[5] +
			9*inn[6] + 4*inn[7] + 6*inn[8] + 8*inn[9]

		b := 3*inn[0] + 7*inn[1] + 2*inn[2] +
			4*inn[3] + 10*inn[4] + 3*inn[5] +
			5*inn[6] + 9*inn[7] + 4*inn[8] +
			6*inn[9] + 8*inn[10]

		sum1 := a % 11
		if sum1 > 9 {
			sum1 = sum1 % 10
		}

		sum2 := b % 11
		if sum2 > 9 {
			sum2 = sum1 % 10
		}

		check1 := inn[10]
		check2 := inn[11]

		if check1 != sum1 && check2 != sum2 {
			return false
		}
	}

	return true
}

var (
	flags = map[int64]bool{
		1: true,
		2: true,
		3: false,
		4: false,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	}
)

// CheckOGRN ...
func CheckOGRN(d int64) bool {

	ogrnLen := DigitsCount(d)
	if ogrnLen != 13 && ogrnLen != 15 {
		return false
	}

	ogrn := IntToSlice(d, []int64{})
	flag := ogrn[0]

	if ogrnLen == 13 { // ЕГРЮЛ

		if flags[flag] == false {
			return false
		}

		a := SliceToInt(ogrn, 11)
		check := ogrn[12]
		sum := a % 11
		if sum > 9 {
			sum = sum % 10
		}

		if sum != check {
			return false
		}

	} else if ogrnLen == 15 { // ЕГРИП

		if flags[flag] == true {
			return false
		}

		a := SliceToInt(ogrn, 13)
		check := ogrn[14]

		sum := a % 13
		if sum > 9 {
			sum = sum % 10
		}

		if sum != check {
			return false
		}
	}

	return true
}
