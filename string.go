package k

import (
	. "github.com/ktye/wg/module"
)

func Kst(x K) (r K) {
	xt := tp(x)
	if xt < 16 {
		r = Str(x)
		if xt == ct {
			r = emb(34, 34, r)
		} else if xt == st {
			r = ucat(Ku(96), r)
		}
	} else {
		xn := nn(x)
		if xn == 0 {
			dx(x)
			return kst0(xt - 16)
		}
		if xt == Lt {
			x = Ech(28, x) // Kst
		} else if xt != ct {
			x = Str(x)
		}
		switch xt - 16 {
		case bt:
			r = cat1(Ech(4, x), Kc('b'))
		case ct:
			r = emb(34, 34, x)
		case it:
			r = join(Kc(' '), x)
		case st:
			r = ucat(Kc(96), join(Kc('`'), x))
		case ft:
			r = join(Kc(' '), x)
		case zt:
			r = join(Kc(' '), x)
		case lt:
			r = emb(40, 41, join(Kc(';'), x))
		default:
			trap(Nyi)
		}
		if xn == 1 {
			r = ucat(Ku(44), r)
		}
	}
	return r
}
func kst0(t T) (r K) {
	switch t {
	case bt:
		r = 1647321904 // 0#0b
	case ct:
		r = 8738 // ""
	case it:
		r = 12321 // !0
	case st:
		r = 6300464 // 0#`
	case ft:
		r = 774906672 // 0#0.
	case zt:
		r = 1630544688 // 0#0a
	case lt:
		r = 10536 // ()
	default:
		trap(Nyi)
	}
	return Ku(uint64(r))
}
func Str(x K) (r K) {
	xt := tp(x)
	if xt > 16 {
		return Ech(17, x)
	}
	xp := int32(x)
	if xt > dt {
		switch xt - cf {
		case 0: // cf
			rx(x)
			r = flat(Str(K(xp) | K(Lt)<<59))
		case 1: // df
			r = Str(x0(xp))
			p := x1(xp)
			if p%2 != 0 {
				p = cat1(Str(20+p), Kc(':'))
			} else {
				p = Str(21 + p)
			}
			r = ucat(r, p)
		case 2: //pf
			r = ucat(Str(x0(xp)), emb('[', ']', join(Kc(';'), Str(x1(xp)))))
		case 3: // lf
			r = x3(xp)
		}
		dx(x)
		return r
	} else {
		switch xt {
		case 0:
			r = Ku(uint64(I8(227 + xp)))
		case bt:
			r = Ku(uint64(25136) + uint64(I8(xp))<<8) // 0b 1b
		case ct:
			r = Ku(uint64(I8(xp)))
		case it:
			r = si(xp)
		case st:
			r = cs(x)
		case ft:
			r = sf(F64(xp))
		case zt:
			trap(Nyi)
		default:
			trap(Err)
		}
	}
	dx(x)
	return r
}
func emb(a, b int32, x K) (r K) { return cat1(Cat(Kc(a), x), Kc(b)) }
func join(x K, y K) (r K) {
	yn := nn(y)
	yp := int32(y)
	r = mk(Ct, 0)
	for i := int32(0); i < yn; i++ {
		if i > 0 {
			r = cat1(r, x)
		}
		r = ucat(r, x0(yp))
		yp += 8
	}
	return r
}
func si(x int32) (r K) {
	if x == 0 {
		return Ku(uint64('0'))
	} else if x < 0 {
		return ucat(Ku(uint64('-')), si(-x))
	}
	r = mk(Ct, 0)
	for x != 0 {
		r = cat1(r, Kc('0'+x%10))
		x /= 10
	}
	return Rev(r)
}
func sf(x float64) (r K) {
	trap(Nyi)
	return 0
}
