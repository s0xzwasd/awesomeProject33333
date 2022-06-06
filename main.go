// You can edit this code!
// Click here and start typing.
package main

type BiConsumer[T any, S any] func(first T, second S)

func main() {
	var consumer BiConsumer[string, string]

	consumer("one", "two") // goland will report: Too many arguments for a conversion to 'consumer'
}
