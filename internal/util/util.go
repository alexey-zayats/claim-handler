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
			9*inn[6] + 4*inn[7] + 6*inn[8] +
			8*inn[9]

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

		if check1 != sum1 || check2 != sum2 {
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

/*
	    А    Б    В    Г    Д    Е    Ё    Ж    З    И    Й    К    Л    М    Н    О    П    Р    С    Т    У    Ф    Х    Ц    Ч    Ш    Щ    Ь    Ы    Ъ    Э    Ю    Я"
	[1040 1041 1042 1043 1044 1045 1025 1046 1047 1048 1049 1050 1051 1052 1053 1054 1055 1056 1057 1058 1059 1060 1061 1062 1063 1064 1065 1068 1067 1066 1069 1070 1071]
	  a  b  c   d   e   f   g   h   i   j   k   l   m   n   o   p   q   r   s   t   u   v   w   x   w  z
	[97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 119 122]
	  A  B  C  D  E  F  G  H  I  J  K  L  M  N  O  P  Q  R  S  T  U  V  W  X  Y  Z
	[65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90]

A,a - А
B(только прописная) - В
Cc - С
Ee - Е
Kk - К
M  - М
H  - Н
Oo - О
Pp - Р
T  - Т
Xx - Х
y (англ. строчная)- У

*/
var table = map[int32]int32{
	97:  1040, // a -> А
	65:  1040, // A -> А
	66:  1042, // B -> В
	99:  1057, // c -> C
	67:  1057, // C -> C
	101: 1045, // e -> Е
	69:  1045, // E -> Е
	107: 1050, // k -> К
	75:  1050, // K -> К
	77:  1052, // M - М
	72:  1053, // H -> Н
	111: 1054, // o -> О
	79:  1054, // O -> О
	112: 1056, // p - Р
	80:  1056, // P - Р
	84:  1058, // T - Т
	120: 1061, // x - Х
	88:  1061, // X - Х
	89:  1059, // Y - У
}

// NormalizeCarNumber ...
func NormalizeCarNumber(number string) string {
	rnum := []rune(number)

	for i := range rnum {
		if i < 7 && rnum[i] > 57 && rnum[i] < 1040 {
			rnum[i] = table[rnum[i]]
		}
	}

	return string(rnum)
}
