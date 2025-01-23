package result

type Result struct {
	status bool
	value  interface{}
}

func Success(value interface{}) *Result {
	return &Result{
		status: true,
		value:  value,
	}
}

func Failure(value interface{}) *Result {
	return &Result{
		status: false,
		value:  value,
	}
}

func (r *Result) IsSuccess() bool {
	return r.status
}

func (r *Result) IsFailure() bool {
	return !r.status
}

func (r *Result) GetValue() interface{} {
	return r.value
}
