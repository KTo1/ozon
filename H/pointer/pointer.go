package pointer

func FromNotEmpty[T comparable](v T) *T {
	var def T
	if v == def {
		return nil
	}

	return &v
}

func From[T any](v T) *T {
	return &v
}

func To[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}

	return *v
}
