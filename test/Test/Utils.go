package Test_Utils

func ThrowErr(msg string) func(interface{}) interface{} {
	return func(interface{}) interface{} {
		panic(msg)
	}
}
