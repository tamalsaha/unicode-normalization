package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "unicode"

    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
)

func main() {

    isMn := func(r rune) bool {
        return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
    }
    t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

    r := strings.NewReader("Sjöström")
    x := transform.NewReader(r, t)
    b, err := ioutil.ReadAll(x)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(b))

}
