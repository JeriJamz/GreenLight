package greetings

import("fmt")

func Hello(name string) string{

	message := fmt.Sprintf("Hi %v, it is nice to meet you", name)
	return message

}

