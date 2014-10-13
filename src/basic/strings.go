package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 文字列を含むかどうかの判定
	Contains()
	ContainsAny()
	ContainsRune()

	// 文字列のカウント
	Count()

	// 文字列のイコール判定
	EqualFold()

	// 文字列のフィールド
	Fields()
	FieldsFunc()

	// 文字列のインデックス
	Index()
	IndexAny()
	IndexByte()
	IndexRune()
	IndexFunc()
	LastIndex()
	LastIndexAny()
	LastIndexFunc()

	// 文字列の連結
	Join()

	// 文字列のマップ
	Map()

	// 文字列の繰り返し
	Repeat()

	// 文字列の置換
	Replace()

	// 文字列の分割
	Split()
	SplitN()
	SplitAfter()
	SplitAfterN()

	// 文字列のタイトル
	Tilte()
	ToTitle()
	ToTitleSpecial()

	// 文字の除去
	Trim()
	TrimLeft()
	TrimLeftFunc()
	TrimRight()
	TrimRightFunc()
	TrimSpace()
	TrimPrefix()
	TrimSuffix()

	// 文字列の変換
	ToUpper()
	ToUpperSpecial()
	ToLower()
	ToLowerSpecial()
}

func startLog(s string) {
	fmt.Println("********** start " + s + " **********")
}

func endLog(s string) {
	fmt.Println("********** end " + s + " **********")
}

func Contains() {
	startLog("Contains")
	fmt.Println(strings.Contains("seafood", "foo"))	// true
	fmt.Println(strings.Contains("seafood", "bar"))	// false
	fmt.Println(strings.Contains("seafood", ""))	// true
	fmt.Println(strings.Contains("", ""))			// true
	endLog("Contains")
}

func ContainsAny() {
	startLog("ContainsAny")
	fmt.Println(strings.ContainsAny("team", "i"))			// false
	fmt.Println(strings.ContainsAny("failure", "u & i"))	// true
	fmt.Println(strings.ContainsAny("foo", ""))				// false
	fmt.Println(strings.ContainsAny("", ""))				// false
	endLog("ContainsAny")
}

func ContainsRune() {
	startLog("ContainsRune")
	fmt.Println(strings.ContainsRune("hello world", 0x01))	// false
	endLog("ContainsRune")
}

func Count() {
	startLog("Count")
	fmt.Println(strings.Count("cheese", "e"))	// 3
	fmt.Println(strings.Count("five", ""))		// 5 before & after each rune
	endLog("Count")
}

func EqualFold() {
	startLog("EqualFold")
	fmt.Println(strings.EqualFold("Go", "Go"))	// true
	fmt.Println(strings.EqualFold("Go", "go"))	// true
	endLog("EqualFold")
}

func Fields() {
	startLog("Fields")
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz  "))	// Fields are: ["foo" "bar" "baz"]
	endLog("Fields")
}

func FieldsFunc() {
	startLog("FieldsFunc")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.FieldsFunc("  foo1;bar2,baz3...", f))
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,baz3...", f))	// Fields are: ["foo1" "bar2" "baz3"]
	endLog("FieldsFunc")
}

func Index() {
	startLog("Index")
	fmt.Println(strings.Index("chicken", "ken"))	// 4
	fmt.Println(strings.Index("chicken", "dmr"))	// -1
	endLog("Index")
}

func IndexAny() {
	startLog("IndexAny")
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))	// 2
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))	// -1
	endLog("IndexAny")
}

func IndexByte() {
	startLog("IndexByte")
	fmt.Println(strings.IndexByte("chicken", 'k'))	// 4
	fmt.Println(strings.IndexByte("chicken", 'a'))	// -1
	endLog("IndexByte")
}

func IndexRune() {
	startLog("IndexRune")
	fmt.Println(strings.IndexRune("chicken", 'k'))	// 4
	fmt.Println(strings.IndexRune("chicken", 'd'))	// -1
	endLog("IndexRune")
}

func IndexFunc() {
	startLog("IndexFunc")
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))		// 7
	fmt.Println(strings.IndexFunc("Hello, world", f))	// -1
	endLog("IndexFunc")
}

func Join() {
	startLog("Join")
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))	// foo, bar, baz
	endLog("Join")
}

func LastIndex() {
	startLog("LastIndex")
	fmt.Println(strings.Index("go gopher", "go"))			// 0
	fmt.Println(strings.LastIndex("go gopher", "go"))		// 3
	fmt.Println(strings.LastIndex("go gopher", "rodent"))	// -1
	endLog("LastIndex")
}

func LastIndexAny() {
	startLog("LastIndexAny")
	fmt.Println(strings.Index("go gopher", "go"))				// 0
	fmt.Println(strings.LastIndexAny("go gopher", "go"))		// 4
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))	// 8
	endLog("LastIndexAny")
}

func LastIndexFunc() {
	startLog("LastIndexFunc")
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.LastIndexFunc("Hello, 世界", f))		// 10
	fmt.Println(strings.LastIndexFunc("Hello, world", f))	// -1
	endLog("LastIndexFunc")
}

func Map() {
	startLog("Map")
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))	// 'Gjnf oevyyvt naq gur fyvgul tbcure...
	endLog("Map")
}

func Repeat() {
	startLog("Repeat")
	fmt.Println("ba" + strings.Repeat("na", 2))		// banana
	endLog("Repeat")
}

func Replace() {
	startLog("Replace")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))		// oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))	// moo moo moo
	endLog("Replace")
}

func Split() {
	startLog("Split")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))		// ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))	// ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))		// [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))		// [""]
	endLog("Split")
}

func SplitN() {
	startLog("SplitN")
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
	endLog("SplitN")
}

func SplitAfter() {
	startLog("SplitAfter")
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))		// ["a," "b," "c"]
	endLog("SplitAfter")
}

func SplitAfterN() {
	startLog("SplitAfterN")
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))	// ["a," "b,c"]
	endLog("SplitAfterN")
}

func Tilte() {
	startLog("Tilte")
	fmt.Println(strings.Title("her royal highness"))	// Her Royal Highness
	endLog("Tilte")
}

func ToTitle() {
	startLog("ToTitle")
	fmt.Println(strings.ToTitle("loud noises"))			// LOUD NOISES
	fmt.Println(strings.ToTitle("хлеб"))				// ХЛЕБ
	endLog("ToTitle")
}

func ToTitleSpecial() {
	startLog("ToTitleSpecial")
	var SC unicode.SpecialCase
	fmt.Println(strings.ToTitleSpecial(SC, "loud noises"))	// LOUD NOISES
	endLog("ToTitleSpecial")
}

func Trim() {
	startLog("Trim")
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))			// ["Achtung! Achtung"]
	endLog("Trim")
}

func TrimLeft() {
	startLog("TrimLeft")
	fmt.Printf("[%q]\n", strings.TrimLeft(" !!! Achtung! Achtung! !!! ", "! "))		// ["Achtung! Achtung! !!! "]
	endLog("TrimLeft")
}

func TrimLeftFunc() {
	startLog("TrimLeftFunc")
	f := func(r rune) bool {
		return r <= 'A'
	}
	fmt.Printf("[%q]\n", strings.TrimLeftFunc(" !!! Achtung! Achtung! !!! ", f))	// ["chtung! Achtung! !!! "]
	endLog("TrimLeftFunc")
}

func TrimRight() {
	startLog("TrimRight")
	fmt.Printf("[%q]\n", strings.TrimRight(" !!! Achtung! Achtung! !!! ", "! "))	// [" !!! Achtung! Achtung"]
	endLog("TrimRight")
}

func TrimRightFunc() {
	startLog("TrimRightFunc")
	f := func(r rune) bool {
		return r <= 'g'
	}
	fmt.Printf("[%q]\n", strings.TrimRightFunc(" !!! Achtung! Achtung! !!! ", f))	// [" !!! Achtung! Achtun"]
	endLog("TrimRightFunc")
}

func TrimSpace() {
	startLog("TrimSpace")
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))					// a lone gopher
	endLog("TrimSpace")
}

func TrimPrefix() {
	startLog("TrimPrefix")
	var s = "Goodbye,, world!"
	s = strings.TrimPrefix(s, "Goodbye,")
	s = strings.TrimPrefix(s, "Howdy,")
	fmt.Print("Hello" + s + "\n")			// Hello, world!
	endLog("TrimPrefix")
}

func TrimSuffix() {
	startLog("TrimSuffix")
	var s = "Hello, goodbye, etc!"
	s = strings.TrimSuffix(s, "goodbye, etc!")
	s = strings.TrimSuffix(s, "planet")
	fmt.Print(s, "world!\n")				// Hello, world!
	endLog("TrimSuffix")
}

func ToUpper() {
	startLog("ToUpper")
	fmt.Println(strings.ToUpper("Gopher"))				// GOPHER
	endLog("ToUpper")
}

func ToUpperSpecial() {
	startLog("ToUpperSpecial")
	var SC unicode.SpecialCase
	fmt.Println(strings.ToUpperSpecial(SC, "Gopher"))	// GOPHER
	endLog("ToUpperSpecial")
}

func ToLower() {
	startLog("ToLower")
	fmt.Println(strings.ToLower("Gopher"))				// gopher
	endLog("ToLower")
}

func ToLowerSpecial() {
	startLog("ToLowerSpecial")
	var SC unicode.SpecialCase
	fmt.Println(strings.ToLowerSpecial(SC, "Gopher"))	// gopher
	endLog("ToLowerSpecial")
}
