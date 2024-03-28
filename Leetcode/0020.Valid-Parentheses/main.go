package validparentheses

type Stack struct {
	runes []rune
}

func NewStack() *Stack {
	return &Stack{runes: []rune{}}
}

func (s *Stack) Push(str rune) {
	s.runes = append(s.runes, str)
}

func (s *Stack) Pop() rune {
	str := s.runes[len(s.runes)-1]
	s.runes = s.runes[:len(s.runes)-1]
	return str
}

// 時間複雜 O(n), 空間複雜 O(n)
func IsValid(s string) bool {
	runeStack := NewStack()
	for _, v := range s {
		// fmt.Println(string(v))
		if v == '(' || v == '[' || v == '{' {
			runeStack.Push(v)
		} else if (v == ')' && len(runeStack.runes) > 0 && runeStack.runes[len(runeStack.runes)-1] == '(') || (v == ']' && len(runeStack.runes) > 0 && runeStack.runes[len(runeStack.runes)-1] == '[') || (v == '}' && len(runeStack.runes) > 0 && runeStack.runes[len(runeStack.runes)-1] == '{') {
			runeStack.Pop()
		} else {
			return false
		}
	}
	return len(runeStack.runes) == 0
}
