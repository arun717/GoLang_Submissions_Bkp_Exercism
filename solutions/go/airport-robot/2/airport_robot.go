package airportrobot
import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(string) string
}

func SayHello(visitorName string, g Greeter) string {
    return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(visitorName))
}

type Italian struct {}

type Portuguese struct {}

func (s Italian) LanguageName() string {
    return "Italian"
}

func (s Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %s!", name)
}

func (s Portuguese) LanguageName() string {
    return "Portuguese"
}

func (s Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %s!", name)
}