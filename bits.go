package basex

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// FormatBits computes the string representation of u in the given base, for 2 <= base <= 62.
// If neg is negative, u is treated as negative int64 value.
func FormatBits(u uint64, base int, negative bool) string {
	return FormatBitsFromCharset(charset, u, base, negative)
}

// FormatBitsFromCharset computes the string representation of u in the given base, for 2 <= base <= len(charset).
// If neg is negative, u is treated as negative int64 value.
func FormatBitsFromCharset(charset string, u uint64, base int, negative bool) string {
	if base < 2 || base > len(charset) {
		panic("basex: illegal base")
	}

	var a [64 + 1]byte // +1 for sign of 64bit value in base 2
	i := len(a)

	if negative {
		u = -u
	}

	b := uint64(base)
	for u >= b {
		i--
		q := u / b
		a[i] = charset[uint(u-q*b)]
		u = q
	}
	i--
	a[i] = charset[uint(u)]

	if negative {
		i--
		a[i] = '-'
	}

	return string(a[i:])
}
