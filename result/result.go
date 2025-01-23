package result

type Result struct {
	Status bool
	Value  interface{}
}

func Success(value interface{}) *Result {
	return &Result{
		Status: true,
		Value:  value,
	}
}

func Failure(value interface{}) *Result {
	return &Result{
		Status: false,
		Value:  value,
	}
}

func (r *Result) IsSuccess() bool {
	return r.Status
}

func (r *Result) IsFailure() bool {
	return !r.Status
}
