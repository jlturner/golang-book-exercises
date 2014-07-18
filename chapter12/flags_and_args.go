package main
import("fmt";"flag";"math/rand";"time")

func main() {
	maxRandom := flag.Int("max", 6, "the max value")
	flag.Parse()

	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(*maxRandom))
	fmt.Println(flag.Args())
}
