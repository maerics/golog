package golog

func Must(err error) {
	if err != nil {
		Fatal(err.Error())
	}
}

func Must1[T any](t T, err error) T {
	Must(err)
	return t
}

func Must2[T any, U any](t T, u U, err error) (T, U) {
	Must(err)
	return t, u
}

func Must3[T any, U any, V any](t T, u U, v V, err error) (T, U, V) {
	Must(err)
	return t, u, v
}
