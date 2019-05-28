package w

type b = byte
type k = uint32
type i = int32
type f = float64

//                i  f   z  s  c  g  l  d
var lns = [9]k{0, 4, 8, 16, 4, 1, 4, 4, 8}
var e k = 0xFFFFFFFF

var m []b

func mk(t b, n int) k { // make type t of len n (-1:atom)
	sz := lns[t]
	if n >= 0 {
		sz *= k(n)
	}
	sz += 8 // size needed including header
	bs := k(16)
	bt := 0
	for i := 4; i < 30; i++ { // calculate bucket bt from size sz (clz)
		if sz <= bs {
			bt = i
			break
		}
		bs <<= 1
	}
	if bt == 0 {
		return e
	}
	fb, a := 0, k(0)
	for i := bt; i < 30; i++ { // find next free bucket >= bt
		if 1<<b(i) > len(m) {
			grw()
		}
		if get(k(4*i)) != 0 {
			fb, a = i, get(k(4*i))
			break
		}
	}
	if fb == 0 {
		return e
	}
	for i := fb; i > bt; i-- {
		put(k(4*i), get(a+8))
		l := k(1) << b(i-1)
		put(k(a+l), k(i-1))
		put(k(a), k(i-1))
		put(k(a+8), a+l)
		put(k(4*i-4), a+l)
	}
	if n < 0 {
		m[int(a+1)] = t
	} else {
		m[int(a)] |= t << 6
	}
	put(a+4, 1) // refcount
	return a
}
func typ(a k) (b, int) { // type and length at addr
	i := int(a)
	t := m[i] >> 6
	if t == 0 {
		return m[int(i+1)], -1
	}
	return t, int(m[i+1]) | int(m[i+2]<<8) | int(m[i+3]<<16)
}
func rst() { // reset memory
	m = make([]b, 1<<16)
	p := k(len(m))
	for i := 15; i > 7; i-- {
		p >>= 1
		m[p] = b(i)
		put(k(4*i), p)
	}
	m[0] = 7
	put(4, 1<<16)
	// TODO: pointer to k-tree at 8
	m[k(1<<7)] = 7
	put(k(4*9), 0)   // no free bucket 9
	put(1<<9, k(73)) // 73: 1<<6|9 (type i, bucket 9), length is ignored
	for i := range lns {
		put(k(4*i+8)+1<<9, k(lns[i]))
	}
	put(4+1<<9, 1) // rc
}
func put(a, x k) {
	i := int(a)
	m[i] = b(x)
	m[i+1] = b(x >> 8)
	m[i+2] = b(x >> 16)
	m[i+3] = b(x >> 24)
}
func get(a k) k { i := int(a); return k(m[i]) | k(m[i+1])<<8 | k(m[i+2])<<16 | k(m[i+3])<<24 }
func grw()      { m = append(m, make([]b, len(m))...) }
func Leb128(v uint32) []byte {
	if v < 0x80 {
		return []byte{byte(v)}
	}
	var b []byte
	for {
		c := uint8(v & 0x7f)
		v >>= 7
		if v != 0 {
			c |= 0x80
		}
		b = append(b, c)
		if c&0x80 == 0 {
			break
		}
	}
	return b
}
