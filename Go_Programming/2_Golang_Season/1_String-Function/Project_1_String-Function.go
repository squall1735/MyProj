package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p("Contains:  ", s.Contains("test", "es")) // มีในชุด string
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("fooo", "o", "0", 2)) // ถ้าพบ 2 ก็แทน 2 ตัวเลย
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p("Len:       ", len("hello _"))
}
