package result

import "github.com/gildemberg-santos/monads-go/maybe"

type Result struct {
	isSuccess bool
	value     interface{}
}

func NewSuccess(value interface{}) *Result {
	return &Result{
		isSuccess: true,
		value:     value,
	}
}

func NewFailure(value interface{}) *Result {
	return &Result{
		isSuccess: false,
		value:     value,
	}
}

func (r *Result) IsSuccess() bool {
	return r.isSuccess
}

func (r *Result) IsFailure() bool {
	return !r.isSuccess
}

func (r *Result) GetValue() interface{} {
	return r.value
}

func (r *Result) Bind(f func(interface{}) *Result) *Result {
	if r.IsSuccess() {
		return f(r.value)
	}
	return r
}

func (r *Result) Fmap(f func(interface{}) interface{}) *Result {
	if r.IsSuccess() {
		return NewSuccess(f(r.value))
	}
	return r
}

func (r *Result) ValueOr(defaultValue interface{}) interface{} {
	if r.IsSuccess() {
		return r.value
	}
	return defaultValue
}

func (r *Result) Value() interface{} {
	if r.IsSuccess() {
		return r.value
	}
	panic("value! was called on Failure")
}

func (r *Result) Or(other *Result) *Result {
	if r.IsSuccess() {
		return r
	}
	return other
}

func (r *Result) Failure() interface{} {
	if r.IsFailure() {
		return r.value
	}
	panic("failure was called on Success")
}

func (r *Result) ToMaybe() *maybe.Maybe {
	if r.IsSuccess() {
		return maybe.NewSome(r.value)
	}
	return maybe.NewNone()
}

func (r *Result) Either(successFunc, failureFunc func(interface{}) interface{}) interface{} {
	if r.IsSuccess() {
		return successFunc(r.value)
	}
	return failureFunc(r.value)
}

func (r *Result) AltMap(f func(interface{}) interface{}) *Result {
	if r.IsFailure() {
		return NewFailure(f(r.value))
	}
	return r
}
