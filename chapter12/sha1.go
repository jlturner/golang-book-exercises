package main
import (
	"fmt"
	"crypto/sha1"
)

func main() {
	h := sha1.New()
	h.Write([]byte("hide"))
	bs := h.Sum([]byte{}) // can you provide salt here?
	fmt.Println(bs)
}
	
	
