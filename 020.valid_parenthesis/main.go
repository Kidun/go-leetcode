package main

// it's classic task for stack AKA LIFO (Last In First Out)
type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func isValid(s string) bool {
	var st Stack
	var ch string

	//simple sanity check
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		ch = s[i : i+1]
		if ch == "(" || ch == "[" || ch == "{" {
			st.Push(ch)
		}
		if ch == ")" || ch == "]" || ch == "}" {
			ch_compl, ch_exist := st.Pop()
			if ch_exist {
				switch ch {
				case ")":
					{
						if ch_compl != "(" {
							return false
						}
					}
				case "]":
					{
						if ch_compl != "[" {
							return false
						}
					}
				case "}":
					{
						if ch_compl != "{" {
							return false
						}
					}
				}
			} else {
				return false
			}
		}
	}
	// in the end, if we have all matching braces, it must be empty
	return st.IsEmpty()
}

func main() {

}
