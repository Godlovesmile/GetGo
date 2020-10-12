package strmodule

import (
	"fmt"
	"strconv"
	"strings"
)

// HandleStrData info
func HandleStrData() {
	// ------------ 字符串操作 strings ----------
	// 1. func Contains(s, substr string) bool; 字符串s中是否包含substr，返回bool值
	fmt.Println(strings.Contains("seafood", "foo1"))

	// 2. func Join(a []string, sep string) string; 字符串链接，把slice a通过sep链接起来
	s := []string{"foo", "test"}
	fmt.Println(strings.Join(s, ", "))

	// 3. func Index(s, sep string) int
	fmt.Println(strings.Index("chicken", "ken"))

	// 4. func Repeat(s string, count int) string; 重复s字符串count次，最后返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na", 2))

	// 5. func Replace(s, old, new string, n int) string; 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))

	// 6. func Split(s, sep string) []string; 把s字符串按照sep分割，返回slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	// 7. func Trim(s string, cutset string) string; 在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))

	// 8. func Fields(s string) []string; 去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

	// ------------ 字符串转换 strconv ----------
	// 1. Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	// 2. Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	// 3. Parse 系列函数把字符串转换为其他类型
	testParse()
}

func testParse() {
	a, err := strconv.ParseBool("false")
	checkError(err)
	b, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a, b, c, d, e)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
