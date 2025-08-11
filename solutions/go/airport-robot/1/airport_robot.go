package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return "I can speak " + g.LanguageName() + ": " + g.Greet(name)
}

type Italian struct {
}

func (it Italian) LanguageName() string {
	return "Italian"
}

func (it Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

type Portuguese struct {
}

func (pt Portuguese) LanguageName() string {
	return "Portuguese"
}

func (pt Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}
