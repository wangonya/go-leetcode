// https://leetcode.com/problems/valid-parentheses/

package goleetcode

import (
	"container/list"
	"fmt"
)

// stack implementation using container/list package

type bracketStack struct {
	stack *list.List
}

func (b *bracketStack) Push(value rune) {
	b.stack.PushFront(value)
}

func (b *bracketStack) Pop() error {
	if b.stack.Len() > 0 {
		b.stack.Remove(b.stack.Front())
	}
	return fmt.Errorf("Pop Error: Stack is empty")
}

func (b *bracketStack) Front() (rune, error)  {
	if b.stack.Len() > 0 {
		if val, ok := b.stack.Front().Value.(rune); ok {
			return val, nil
		}
		return 0, fmt.Errorf("Peep Error: Stack datatype is incorrect")
	}
	return 0, fmt.Errorf("Peep Error: Stack is empty")
}

func validParentheses(s string) bool {
	if len(s) < 2 {
		return false
	}
	
	b:=&bracketStack{
		stack: list.New(),
	}

	m := map[rune]rune{
		')': '(',
		']':'[',
		'}':'{',
	}

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			b.Push(v)
		} else {
			match, err := b.Front()
			if err != nil {
				return false
			} else if m[v] != match {
				return false
			}
			b.Pop()
		}
	}

	return b.stack.Len() == 0
}
