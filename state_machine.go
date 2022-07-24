package main

import (
	"errors"
	"fmt"
)

type TransitionFunc[T any] func(g T) error

type StateMachine[T any] struct {
	g                   T
	prevState, curState string
	transitions         map[string]TransitionFunc[T]
}

type Transition[T any] struct {
	From, To                  string
	ForwardFunc, BackwardFunc TransitionFunc[T]
}

func NewStateMachine[T any](g T, initialState string) *StateMachine[T] {
	return &StateMachine[T]{
		g:           g,
		curState:    initialState,
		transitions: make(map[string]TransitionFunc[T]),
	}
}

func (st StateMachine[T]) AddTransition(t Transition[T]) {
	if t.From == "" || t.To == "" {
		panic("fields From and To are both required")
	}
	if t.ForwardFunc == nil && t.BackwardFunc == nil {
		panic("neither ForwardFunc nor BackwardFunc is specified")
	}

	if t.ForwardFunc != nil {
		st.addTransition(t.From, t.To, t.ForwardFunc)
	}
	if t.BackwardFunc != nil {
		st.addTransition(t.To, t.From, t.BackwardFunc)
	}
}

func (st *StateMachine[T]) addTransition(from, to string, f TransitionFunc[T]) error {
	k := fmt.Sprintf("%v->%v", from, to)
	if _, ok := st.transitions[k]; ok {
		panic(fmt.Sprintf("transition from %v to %v already exists", from, to))
	}
	st.transitions[k] = f
	return nil
}

func (st *StateMachine[T]) Transit(to string) error {
	k := fmt.Sprintf("%v->%v", st.curState, to)
	f, ok := st.transitions[k]
	if !ok {
		return fmt.Errorf("transition from %v to %v doesn't exist", st.curState, to)
	}
	st.prevState = st.curState
	st.curState = to
	return f(st.g)
}

func (st *StateMachine[T]) Toggle(state string) error {

	if st.curState == state { // if we're already here
		if st.prevState == "" { // and there is no way back
			return errors.New("cannot be back")
		}
		state = st.prevState
	}

	if st.checkBi(st.curState, state) {
		return fmt.Errorf("transitions from %v to %v and vice versa don't exist", st.curState, state)
	}

	return st.Transit(state)
}

func (st *StateMachine[T]) checkBi(s1, s2 string) bool {
	t1 := fmt.Sprintf("%v->%v", s1, s2)
	t2 := fmt.Sprintf("%v->%v", s2, s1)

	_, ok1 := st.transitions[t1]
	_, ok2 := st.transitions[t2]

	return !ok1 || !ok2
}
