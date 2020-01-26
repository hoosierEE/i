package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	testCases := []struct {
		sig string
		b   string
		e   string
	}{
		{"I:II", "x+y", "20002001 6a"},
		//{"I:I", "x/r:i;r", "00?"},
		//{"I:I", "x/r+:i;r", "00?"},
		{"I:II", "(3+x)*y", "4103 2000 6a 2001 6c"},
		{"I:I", "1+x", "410120006a"},
		{"F:FF", "(x*y)", "20002001 a2"},
		{"F:FF", "x-y", "20002001 a1"},
		{"F:FF", "3.*x+y", "44 0000000000000840 20002001 a0 a2"},
		{"I:I", "x:1+x;x*2", "4101 2000 6a 2100 2000 4102 6c"},
	}
	for n, tc := range testCases {
		f := newfn(tc.sig, tc.b)
		e := f.parse(nil, nil)
		b := string(hex(e.bytes()))
		s := trim(tc.e)
		if b != s {
			t.Fatalf("%d: expected/got:\n%s\n%s", n+1, s, b)
		}
		fmt.Println(b)
	}
}
func TestRun(t *testing.T) {
	g := s(hex(run(strings.NewReader("add:I:II{x+y}/cnt\n/\n/sum:I:I{x/r+:i;r}\n/")).wasm()))
	e := "0061736d0100000001070160027f7f017f0302010005030100010707010361646400000a0b010901027f200020016a0b"
	if e != g {
		t.Fatalf("expected/got\n%s\n%s\n", e, g)
	}
}

func hex(a []c) []c {
	var r bytes.Buffer
	for _, b := range a {
		hi, lo := hxb(b)
		r.WriteByte(hi)
		r.WriteByte(lo)
	}
	return r.Bytes()
}
func newfn(sig string, body string) fn {
	var buf bytes.Buffer
	buf.WriteString(body)
	buf.WriteByte('}')
	v := strings.Split(sig, ":")
	if len(v) != 2 {
		panic("signature")
	}
	f := fn{src: [2]int{1, 0}, Buffer: buf}
	f.t = typs[v[0][0]]
	for _, c := range v[1] {
		f.locl = append(f.locl, typs[byte(c)])
	}
	f.args = len(v[1])
	return f
}
func trim(s string) string { return strings.Replace(s, " ", "", -1) }
