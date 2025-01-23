package result

type Result struct {
	isSuccess bool
	value     interface{}
}

func (r *Result) IsSuccess() bool {
	return r.isSuccess
}

func (r *Result) IsFailure() bool {
	return !r.isSuccess
}

func (r *Result) GetSuccess() interface{} {
	if r.isSuccess {
		return r.value
	}
	return nil
}

func (r *Result) GetFailure() interface{} {
	if !r.isSuccess {
		return r.value
	}
	return nil
}

func NewSuccess(values ...interface{}) *Result {
	return &Result{
		isSuccess: true,
		value:     valueOf(values...),
	}
}

func NewFailure(values ...interface{}) *Result {
	return &Result{
		isSuccess: false,
		value:     valueOf(values...),
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
