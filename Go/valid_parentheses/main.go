package main

import "fmt"

type stack struct {
	s   []rune
	top int
}

func (s *stack) push(r rune) {
	s.top++
	if s.top == len(s.s) {
		s.s = append(s.s, make([]rune, 10)...)
	}
	s.s[s.top] = r
}

func (s *stack) pop() (rune, bool) {
	if s.top == -1 {
		return 0, false
	}
	top := s.s[s.top]
	s.top--
	return top, true
}

func NewStack() *stack {
	return &stack{s: make([]rune, 10), top: -1}
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	st := NewStack()
	for i := 0; i < len(s); i++ {
		r := rune(s[i])

		if r == '(' || r == '[' || r == '{' {
			st.push(r)
			continue
		}
		switch r {
		case ')':
			poped, ok := st.pop()
			if !ok || poped != '(' {
				return false
			}
			break
		case '}':
			poped, ok := st.pop()
			if !ok || poped != '{' {
				return false
			}
			break
		case ']':
			poped, ok := st.pop()
			if !ok || poped != '[' {
				return false
			}
			break
		}
	}
	return st.top == -1
}

func main() {
	fmt.Println(isValid("){"))
}
