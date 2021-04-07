package utils

// HandleError : Handle errors.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
