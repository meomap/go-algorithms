// Package stack provides Stack implementation
package stack

import "testing"

func TestNewStack(t *testing.T) {
	s := New()
	if s.Top != 0 {
		t.Errorf("Expect empty stack but got %q", s)
	}
}

func TestPush(t *testing.T) {
	s := New()

	s.Push(1)
	if s.Top != 1 || len(s.Items) != 1 {
		t.Errorf("Expect item 1 but got %q", s)
	}
	if val, ok := s.Items[0].(int); !ok || val != 1 {
		t.Errorf("Expect pushed 1 into items stack but got %q", s)
	}

	// push more
	s.Push(3)
	s.Push(5)
	if s.Top != 3 || len(s.Items) != 3 {
		t.Errorf("Expect 3 items but got %q", s)
	}
	if val, ok := s.Items[2].(int); !ok || val != 5 {
		t.Errorf("Expect last item is 5 but got %q", s)
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(3)
	s.Push(5)

	item := s.Pop()
	if item != 5 {
		t.Errorf("Expect popped 5 but got %#v", item)
	}
	if s.Top != 2 {
		t.Errorf("Expect stack shrinked to 2 items but got %q", s)
	}

	// pop more
	item = s.Pop()
	if item != 3 {
		t.Errorf("Expect popped 3 but got %#v", item)
	}
	if s.Top != 1 {
		t.Errorf("Expect stack shrinked to 1 items but got %q", s)
	}
}
