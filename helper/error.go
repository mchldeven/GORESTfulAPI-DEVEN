package helper

func PanicIfError(err erorr) {
	if err != nil {
		panic
	}
}