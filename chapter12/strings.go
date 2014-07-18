package main
import(
	"fmt"
	"strings"
)

func main() {
	test := "test"
	fmt.Println(
		strings.Contains(test, "es"),
		strings.Count(test, "t"),
		strings.HasPrefix(test, "te"),
		strings.HasSuffix(test, "st"),
		strings.Index("test", "e"),
		strings.Join([]string{"a","b"}, "-"),
		strings.Repeat("lol",5),
		strings.Replace("aaaaa","a","b",2),
		strings.Split("Hi there everybody", " "),
		strings.ToLower("James"),
		strings.ToUpper("James"),
	)
}
