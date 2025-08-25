package airportrobot

type Greeter interface {
    LanguageName() string
    Greet(string) string
}

func SayHello(visitorName string, g Greeter) string {
    return "I can speak " + g.LanguageName() + ": " + g.Greet(visitorName) 
}

type Italian struct {}

type Portuguese struct {}

func (s Italian) LanguageName() string {
    return "Italian"
}

func (s Italian) Greet(name string) string {
    return "Ciao " + name + "!"
}

func (s Portuguese) LanguageName() string {
    return "Portuguese"
}

func (s Portuguese) Greet(name string) string {
    return "Olá " + name + "!"
}

