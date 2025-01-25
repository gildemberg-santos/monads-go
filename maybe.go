package main

type Maybe struct {
	isSome bool
	value  interface{}
}

func NewSome(value interface{}) *Maybe {
	return &Maybe{
		isSome: true,
		value:  value,
	}
}

func NewNone() *Maybe {
	return &Maybe{
		isSome: false,
	}
}

func (m *Maybe) IsSome() bool {
	return m.isSome
}

func (m *Maybe) IsNone() bool {
	return !m.isSome
}

func (m *Maybe) GetValue() interface{} {
	if m.IsSome() {
		return m.value
	}
	return nil
}

func (m *Maybe) Bind(f func(interface{}) *Maybe) *Maybe {
	if m.IsSome() {
		return f(m.value)
	}
	return m
}

func (m *Maybe) Fmap(f func(interface{}) interface{}) *Maybe {
	if m.IsSome() {
		return NewSome(f(m.value))
	}
	return m
}

func (m *Maybe) Maybe(f func(interface{}) interface{}) *Maybe {
	if m.IsSome() {
		result := f(m.value)
		if result == nil {
			return NewNone()
		}
		return NewSome(result)
	}
	return m
}

func (m *Maybe) ValueOr(defaultValue interface{}) interface{} {
	if m.IsSome() {
		return m.value
	}
	return defaultValue
}

func (m *Maybe) Or(other *Maybe) *Maybe {
	if m.IsSome() {
		return m
	}
	return other
}

func (m *Maybe) And(other *Maybe) *Maybe {
	if m.IsSome() && other.IsSome() {
		return NewSome([]interface{}{m.value, other.value})
	}
	return NewNone()
}

func (m *Maybe) Flatten() *Maybe {
	if m.IsSome() {
		if nestedMaybe, ok := m.value.(*Maybe); ok {
			return nestedMaybe
		}
	}
	return m
}
