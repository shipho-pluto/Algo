package Language

import "errors"

func Grating(lang string) (string, error) {
	switch lang {
	case "English":
		return "Hello", nil
	case "Spanish":
		return "Hola", nil
	case "French":
		return "Sova", nil
	case "German":
		return "Hallo", nil
	default:
		return "", errors.New("unsupported language")
	}
}
