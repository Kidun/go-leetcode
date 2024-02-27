package isAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// implement basic checksum for the string, where index doesn't impact the result
	// xor as an additional check to avoid collisions, the main issue with xor is -> x XOR x == 0
	var ps, pt int
	var x byte
	for i := 0; i < len(s); i++ {
		x ^= s[i] ^ t[i]
		ps += int(s[i] * s[i])
		pt += int(t[i] * t[i])
	}
	return x == 0 && ps == pt
}
