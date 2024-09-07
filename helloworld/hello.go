package helloworld

const (
	arabic             = "Arabic"
	french             = "French"
	englishHelloPrefix = "Hello, "
	arabicHelloPrefix  = "اهلا "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case arabic:
		prefix = arabicHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	// fmt.Println(Hello(""))
}