package main

import (
  "fmt"
)

type Node struct {
  Value int
  Next  *Node
}

type Stack struct {
  top  *Node
  size int
}

func (s *Stack) isEmpty() bool {
  return s.size == 0
}

func (s *Stack) Push(v int) {
  //*s = append(*s, v)
  node := Node{v, s.top}
  s.top = &node
  s.size++
}

func (s *Stack) Pop() (int, bool) {
  if s.isEmpty() {
    return 0, false
  }
  data := s.top.Value
  tmp := s.top
  s.top = s.top.Next
  tmp.Next = &Node{}
  s.size--
  return data, true
}

func (s *Stack) Peek() (int, bool) {
  if s.isEmpty() {
    return 0, false
  }
  return s.top.Value, true
}

func (s *Stack) Clear() {
  s.top = nil
  s.size = 0
}

func (s *Stack) Contains(v int) bool {
  tmp := s.top
  for tmp != nil {
    if tmp.Value == v {
      return true
    }
    tmp = tmp.Next
  }
  return false
}

func (s *Stack) Increment() {
  a := make([]int, 0)
  tmp := s.top
  for tmp != nil {
    a = append(a, tmp.Value+1)
    tmp = tmp.Next
  }
  s.Clear()
  for i := len(a) - 1; i >= 0; i-- {
    s.Push(a[i])
  }
}

func (s *Stack) Print() {
  tmp := s.top
  for tmp != nil {
    fmt.Print(tmp.Value, " ")
    tmp = tmp.Next
  }
  fmt.Println()
}

func (s *Stack) PrintReverse() {
  tmp := s.top
  a := make([]int, 0)
  for tmp != nil {
    a = append(a, tmp.Value)
    tmp = tmp.Next
  }
  for i := len(a) - 1; i >= 0; i-- {
    fmt.Print(a[i], " ")
  }
  fmt.Println()
}
func main() {
  s := Stack{}
  s.Push(1)
  s.Push(2)
  s.Push(5)
  s.Push(9)
  s.Push(89)
  fmt.Println(s.Pop())
  fmt.Println(s.Peek())
  fmt.Println(s.Contains(9))
  fmt.Println(s.Contains(999))
  s.Print()
  s.Increment()
  s.Print()
  s.PrintReverse()
}