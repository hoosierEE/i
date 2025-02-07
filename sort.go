package main

import (
	. "github.com/ktye/wg/module"
)

func Srt(x K) K { // ^x
	r := K(0)
	xt := tp(x)
	if xt < 16 {
		trap(Type)
	}
	if xt == Dt {
		r = x0(x)
		x = r1(x)
		i := rx(Asc(rx(x)))
		return Key(atv(r, i), atv(x, i))
	}
	if nn(x) < 2 {
		return x
	}
	return atv(x, Asc(rx(x)))
}
func Asc(x K) K { // <x  <`file
	if tp(x) == st {
		return readfile(cs(x))
	}
	return grade(x, 343)
}
func Dsc(x K) K { return grade(x, 336) } // >x
func grade(x K, f int32) K { // <x >x
	r := K(0)
	xt := tp(x)
	if xt < 16 {
		trap(Type)
	}
	if xt == Dt {
		r = x0(x)
		return Atx(r, grade(r1(x), f))
	}
	n := nn(x)
	if xt == Tt {
		return kxy(104, x, Ki(I32B(f == 336))) //gdt ngn:{(!#x){x@<y x}/|.+x}
	}
	if n < 2 {
		dx(x)
		return seq(n)
	}
	r = seq(n)
	rp := int32(r)
	xp := int32(x)
	//	if n < 16 {
	//		igrd(rp, xp, n, sz(xt), f+int32(xt))
	//	} else {
	w := mk(It, n)
	wp := int32(w)
	Memorycopy(wp, rp, 4*n)
	msrt(wp, rp, 0, n, xp, sz(xt), f+int32(xt))
	dx(w)
	//	}
	dx(x)
	return r
}

func msrt(x, r, a, b, p, s, f int32) {
	if b-a < 2 {
		return
	}
	c := (a + b) >> 1
	msrt(r, x, a, c, p, s, f)
	msrt(r, x, c, b, p, s, f)
	mrge(x, r, 4*a, 4*b, 4*c, p, s, f)
}
func mrge(x, r, a, b, c, p, s, f int32) {
	q := int32(0)
	i, j := a, c
	for k := a; k < b; k += 4 {
		if i < c && j < b {
			q = Func[f].(f2i)(p+s*I32(x+i), p+s*I32(x+j))
		} else {
			q = 0
		}
		if i >= c || q != 0 {
			SetI32(r+k, I32(x+j))
			j += 4
		} else {
			SetI32(r+k, I32(x+i))
			i += 4
		}
	}
}

func guC(xp, yp int32) int32 { return I32B(I8(xp) < I8(yp)) }
func guI(xp, yp int32) int32 { return I32B(I32(xp) < I32(yp)) }
func guF(xp, yp int32) int32 { return ltf(F64(xp), F64(yp)) }
func guZ(xp, yp int32) int32 { return ltz(F64(xp), F64(xp+8), F64(yp), F64(yp+8)) }
func guL(xp, yp int32) int32 { return ltL(K(I64(xp)), K(I64(yp))) }

func gdC(xp, yp int32) int32 { return I32B(I8(xp) > I8(yp)) }
func gdI(xp, yp int32) int32 { return I32B(I32(xp) > I32(yp)) }
func gdF(xp, yp int32) int32 { return guF(yp, xp) }
func gdZ(xp, yp int32) int32 { return guZ(yp, xp) }
func gdL(xp, yp int32) int32 { return guL(yp, xp) }

func ltL(x, y K) int32 { // sort lists lexically
	r := int32(0)
	xt := tp(x)
	if xt != tp(y) {
		return I32B(xt < tp(y))
	}
	if xt < 16 {
		return int32(Les(rx(x), rx(y)))
	}
	xp, yp := int32(x), int32(y)
	if xt > Lt {
		a, b := K(I64(xp)), K(I64(yp))
		if match(a, b) == 0 {
			return ltL(a, b)
		}
		return ltL(K(I64(xp+8)), K(I64(yp+8)))
	}
	xn, yn := nn(x), nn(y)
	n := mini(xn, yn)
	switch sz(xt) >> 2 {
	case 0:
		r = taoC(xp, yp, n)
	case 1:
		r = taoI(xp, yp, n)
	case 2:
		if xt == Lt {
			r = taoL(xp, yp, n)
		} else {
			r = taoF(xp, yp, n)
		}
	default:
		r = taoZ(xp, yp, n)
	}
	if r == 2 {
		return I32B(xn < yn)
	} else {
		return r
	}
}
func taoC(xp, yp, n int32) int32 {
	e := xp + n
	for xp < e {
		if I8(xp) != I8(yp) {
			return I32B(I8(xp) < I8(yp))
		}
		yp++
		xp++
	}
	return 2
}
func taoI(xp, yp, n int32) int32 {
	e := xp + 4*n
	for xp < e {
		if I32(xp) != I32(yp) {
			return I32B(I32(xp) < I32(yp))
		}
		yp += 4
		xp += 4
	}
	return 2
}
func taoL(xp, yp, n int32) int32 {
	e := xp + 8*n
	for xp < e {
		x, y := K(I64(xp)), K(I64(yp))
		if match(x, y) == 0 {
			return ltL(x, y)
		}
		yp += 8
		xp += 8
	}
	return 2
}
func taoF(xp, yp, n int32) int32 {
	e := xp + 8*n
	for xp < e {
		x, y := F64(xp), F64(yp)
		if eqf(x, y) == 0 {
			return ltf(x, y)
		}
		yp += 8
		xp += 8
	}
	return 2
}
func taoZ(xp, yp, n int32) int32 {
	e := xp + 16*n
	for xp < e {
		xr, xi, yr, yi := F64(xp), F64(xp+8), F64(yp), F64(yp+8)
		if eqz(xr, xi, yr, yi) == 0 {
			return ltz(xr, xi, yr, yi)
		}
		yp += 16
		xp += 16
	}
	return 2
}
