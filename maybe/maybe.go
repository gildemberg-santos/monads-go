package maybe

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
