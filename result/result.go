package result

type Result struct {
	status bool
	value  interface{}
}

func (r *Result) IsSuccessful() bool {
	return r.status
}

func (r *Result) IsFailure() bool {
	return !r.status
}

func (r *Result) GetValue() interface{} {
	return r.value
}

func NewSuccess(values ...interface{}) *Result {
	return &Result{
		status: true,
		value:  valueOf(values...),
	}
}

func NewFailure(values ...interface{}) *Result {
	return &Result{
		status: false,
		value:  valueOf(values...),
	}
}

func valueOf(values ...interface{}) interface{} {
	if len(values) == 1 {
		return values[0]
	} else if len(values) > 1 {
		return values
	}
	return nil
}
