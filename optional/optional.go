package optional

type Optional[T any] struct {
	value *T
}

func Empty[T any]() Optional[T] {
	return Optional[T]{value: nil}
}

func New[T any](value *T) Optional[T] {
	return Optional[T]{
		value: value,
	}
}

func NewNotNil[T any](value *T) Optional[T] {
	if value == nil {
		return Empty[T]()
	}
	return New(value)
}

// Or
/*func NewNotNil[T any](value T) (Optional[T], error) {
	return New(&value), nil
}
*/

func (o *Optional[T]) HasValue() bool {
	return o.value != nil
}

func (o *Optional[T]) IsEmpty() bool {
	return o.value == nil
}

func (o *Optional[T]) OrElse(other T) T {
	if !o.HasValue() {
		return other
	}
	return *o.value
}

func (o *Optional[T]) OrElseError(e error) (*T, error) {
	if o.value == nil {
		return o.value, e
	}
	return o.value, nil
}

func (o *Optional[T]) DoIfPresent(action func(T)) {
	if o.value != nil {
		action(*o.value)
	}
}
