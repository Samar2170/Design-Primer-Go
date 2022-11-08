package structural

type Action func(int) int

func LogAction(action Action) Action {
	return func(n int) int {
		println("Calling action")
		return action(n)
	}
}

func Double(n int) int {
	return n * 2
}

func main() {
	action := LogAction(Double)
	println(action(3))
}
