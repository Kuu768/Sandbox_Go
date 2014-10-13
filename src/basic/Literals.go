package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
	"math"
)

func main() {
	// Lieterals, Operators, and Escapes
	text1 := "\"what's that?\", he said"	// Interpreted string literal
	text2 := `"what's that?", he said`		// Raw string literal
	raidacals := "√ \u221A \U0000221a"		// radicals == "√ √ √"
	fmt.Println(text1)
	fmt.Println(text2)
	fmt.Println(raidacals)

	book := "The Spirit Level" +
			" by Richard Wilkinson"						// String concatenation
	book += " and Kate Pickett"							// String append
	fmt.Println(book)
	fmt.Println("Josey" < "Jose", "Josey" == "Jose")	// String comparisons

	æ := ""
	for _, char := range []rune{'æ', 0xE6, 0346, 230, '\xE6', '\u00E6'} {
		fmt.Printf("[0x%X '%c'] ", char, char)
		æ += string(char)
	}
	fmt.Println()

	// Characters and Strings
	phrase := "test"
	fmt.Printf("string: \"%s\"\n", phrase)
	fmt.Println("index rune    char bytes")
	for index, char := range phrase {
		fmt.Printf("%-2d    %U  '%c'  %X\n",
		index, char, char, []byte(string(char)))
	}

	// Indexing and Slicing Strings
	line1 := "line1 line2 line3 line4 line5"
	i := strings.Index(line1, " ")
	firstWord := line1[:i]
	j := strings.LastIndex(line1, " ")
	lastWord := line1[j+1:]
	fmt.Println(firstWord, lastWord)

	line2 := "word1 word2 word3 word4 word5"
	k := strings.IndexFunc(line2, unicode.IsSpace)
	firstWord2 := line2[:k]
	l := strings.LastIndexFunc(line2, unicode.IsSpace)
	_, size := utf8.DecodeRuneInString(line2[l:])
	lastword2 := line2[l+size:]
	fmt.Println(firstWord2, lastword2)

	// String Formatting with the Fmt Package
	type polar struct{radius, θ float64}
	p := polar{8.32, .49}
	fmt.Print(-18.5, 17, "Elephant", -8+.7i, 0x3C7, '\u03C7', "a", "b", p)
	fmt.Println()
	fmt.Println(-18.5, 17, "Elephant", -8+.7i, 0x3C7, '\u03C7', "a", "b", p)

	// Formatting Booleans
	fmt.Printf("%t %t\n", true, false)
	fmt.Printf("%d %d\n", IntForBool(true), IntForBool(false))

	// Formatting Integers
	fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n", 37, 37, 37, 37, 37)
	fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41)
	m := 3931
	fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", m, m, m, m, m, m)
	n := 569
	fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", n, n, n, Pad(n, 6, '*'))

	// Formatting Characters
	fmt.Printf("%d %#04x %U '%c'\n", 0x3A6, 934, '\u03A6', '\U000003A6')

	// Formatting Floating-Point Numbers
	for _, x := range []float64{-.258, 7194.84, -60897162.0218, 1.500089e-8} {
		fmt.Printf("|%20.5e|%20.5f|%s|\n", x, x, Humanize(x, 20, 5, '*', ','))
	}

	for _, x := range []complex128{2 + 3i, 172.6 - 58.3019i, -.827e2 + 9.04831e-3i} {
		fmt.Printf("|%15s|%9.3f|%.2f|%.1e|\n", fmt.Sprintf("%6.2f%+.3fi", real(x), imag(x)), x, x, x)
	}

	// Formatting Strings and Slices
	slogan := "End ♥"
	fmt.Printf("%s\n%q\n%+q\n%#q\n", slogan, slogan, slogan, slogan)

	chars := []rune(slogan)
	fmt.Printf("%x\n%#x\n%#X\n", chars, chars, chars)

	bytes := []byte(slogan)
	fmt.Printf("%s\n%x\n%X\n% X\n%v\n", bytes, bytes, bytes, bytes, bytes)

	str := "dare to be naive"
	fmt.Printf("|%22s|%-22s|%10s|\n", str, str, str)

	idx := strings.Index(str, "n")
	fmt.Printf("|%.10s|%.*s|%-22.10s|%s|\n", str, idx, str, str, str)

	// Formatting for Debugging
	pol := polar{-83.40, 71.60}
	fmt.Printf("|%T|%v|%#v|\n", pol, pol, pol)
	fmt.Printf("|%T|%v|%t|\n", false, false, false)
	fmt.Printf("|%T|%v|%d|\n", 7607, 7607, 7607)
	fmt.Printf("|%T|%v|%f|\n", math.E, math.E, math.E)
	fmt.Printf("|%T|%v|%f|\n", 5+7i, 5+7i, 5+7i)
	str1 := "Realativity"
	fmt.Printf("|%T|\"%v\"|\"%s\"|%q|\n", str1, str1, str1, str1)

	str2 := "AliasSynonym"
	chars2 := []rune(str2)
	bytes2 := []byte(str2)
	fmt.Printf("%T: %v\n%T: %v\n", chars2, chars2, bytes2, bytes2)

	i2 := 5
	fl := -48.3124
	str3 := "Tomas Breton"
	fmt.Printf("|%p → %d|%p → %f|%#p → %s|\n", &i2, i2, &fl, fl, &str3, str3)

	fmt.Println([]float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%v\n", []float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%#v\n", []float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%.5f\n", []float64{math.E, math.Pi, math.Phi})

	fmt.Printf("%q\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%v\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%#v\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%17s\n", []string{"Software patents", "kill", "innovation"})

	fmt.Printf("%v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%#v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	fmt.Printf("%#v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	fmt.Printf("%04b\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})

	// Other String-Related Packages
	names := "Nicole●Noel●Geofferey●Amelie●●Jose"
	fmt.Print("|")
	for _, name := range strings.Split(names, "●") {
		fmt.Printf("%s|", name)
	}
	fmt.Println()

	for _, record := range []string{"Test1*1892*1963", "Test2\t1823\t1892", "Test3|1775|1814"} {
		fmt.Println(strings.FieldsFunc(record, func(char rune) bool {
				switch char {
				case '\t', '*', '|':
					return true
				}
				return false
			}))
	}
}

func IntForBool(b bool) int {
	if b{
		return 1
	}
	return 0
}

func Pad(number, width int, pad rune) string {
	str := fmt.Sprint(number)
	gap := width - utf8.RuneCountInString(str)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + str
	}
	return str
}

func Humanize(amount float64, width, decimals int, pad, separator rune) string {
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:] // Strip ”±"
	fraction := ""
	if decimals > 0 {
		fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:] // Strip "±0"
	}
	sep := string(separator)
	for i := len(whole) -3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]
	}
	if amount < 0.0 {
		whole = "-" + whole
	}
	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + number
	}
	return number
}
