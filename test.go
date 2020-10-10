package test

import "io"

func f(w io.Writer, x []byte, s string) {
	io.WriteString(w, string(x))
	io.WriteString(w, s)
	io.WriteString(w, string(s))
}
