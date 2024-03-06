package reverseString

import "testing"

// helper function to test slice's content
func printNumArray(slice []byte) string {
	if len(slice) <= 0 {
		return "Empty slice"
	}
	res := ""
	for i := 0; i < len(slice); {
		res = res + string(slice[i])
		i++
	}
	return res
}

func Test_reverseString(t *testing.T) {
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	str_a := printNumArray(a)
	reverseString(a)
	str_res := printNumArray(a)
	if str_res != "olleh" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_res + "}")
	}

	a = []byte{'Z'}
	str_a = printNumArray(a)
	reverseString(a)
	str_res = printNumArray(a)
	if str_res != "Z" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_res + "}")
	}

	a = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	str_a = printNumArray(a)
	reverseString(a)
	str_res = printNumArray(a)
	if str_res != "hannaH" {
		t.Error("unexpected value for {" + str_a + "}, {" + str_res + "}")
	}
}
