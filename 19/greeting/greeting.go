package greeting

func Hello(name, lang string) string {

	if IsEmpty(name) {
		name = "Aaron Parandian"
	}

	switch lang {
	case "en":
		return "Hello, " + name
	case "es":
		return "Hola, " + name
	case "fr":
		return "Bonjour, " + name
	default:
		return "Hello, " + name
	}
}

func IsEmpty(s string) bool {
	return len(s) == 0
}
