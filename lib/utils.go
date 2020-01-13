package pages

func santizeName(name string) (sanitized string) {
	var buf []rune
	for _, r := range name {
		switch r {
		case 'a', 'A':
			buf = append(buf, 'a')
		case 'b', 'B':
			buf = append(buf, 'b')
		case 'c', 'C':
			buf = append(buf, 'c')
		case 'd', 'D':
			buf = append(buf, 'd')
		case 'e', 'E':
			buf = append(buf, 'e')
		case 'f', 'F':
			buf = append(buf, 'f')
		case 'g', 'G':
			buf = append(buf, 'g')
		case 'h', 'H':
			buf = append(buf, 'h')
		case 'i', 'I':
			buf = append(buf, 'i')
		case 'j', 'J':
			buf = append(buf, 'j')
		case 'k', 'K':
			buf = append(buf, 'k')
		case 'l', 'L':
			buf = append(buf, 'l')
		case 'm', 'M':
			buf = append(buf, 'm')
		case 'n', 'N':
			buf = append(buf, 'n')
		case 'o', 'O':
			buf = append(buf, 'o')
		case 'p', 'P':
			buf = append(buf, 'p')
		case 'q', 'Q':
			buf = append(buf, 'q')
		case 'r', 'R':
			buf = append(buf, 'r')
		case 's', 'S':
			buf = append(buf, 's')
		case 't', 'T':
			buf = append(buf, 't')
		case 'u', 'U':
			buf = append(buf, 'u')
		case 'v', 'V':
			buf = append(buf, 'v')
		case 'w', 'W':
			buf = append(buf, 'w')
		case 'x', 'X':
			buf = append(buf, 'x')
		case 'y', 'Y':
			buf = append(buf, 'y')
		case 'z', 'Z':
			buf = append(buf, 'z')
		case ' ', '-', '_':
			buf = append(buf, '-')
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			buf = append(buf, r)
		}
	}

	return string(buf)
}
