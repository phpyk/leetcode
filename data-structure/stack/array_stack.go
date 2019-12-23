package stack

import (
	"errors"
)
//数组实现的顺序栈
type ArrayStack struct {
	items []interface{} //栈中数据
	count int //栈中元素个数
	n int //栈的大小
}

func (s *ArrayStack) Init(l int) *ArrayStack {
	s.count = 0
	s.n = l
	return s
}

func (s *ArrayStack) Push(value string) (bool,error) {
	if s.count == s.n {
		return false,errors.New("not enough space")
	}
	s.items[s.count] = value
	s.count++
	return true,nil
}

func (s *ArrayStack) Pop() (string,error) {
	l := len(s.items)
	if s.count == 0 || l == 0 {
		return "",errors.New("empty stack")
	}
	val := s.items[l-1].(string)
	s.count--
	s.items = s.items[:l-2]
	return val,nil
}