package app

type Token string

const (
	UNKNOWN_TOKEN     Token = "UNKNOWN_TOKEN"
	OUTPUT            Token = "OUTPUT"            // .
	INPUT             Token = "INPUT"             // ,
	INCREMENT         Token = "INCREMENT"         // +
	DECREMENT         Token = "DECREMENT"         // -
	INCREMENT_POINTER Token = "INCREMENT_POINTER" // >
	DECREMENT_POINTER Token = "DECREMENT_POINTER" // <
	BEGIN             Token = "BEGIN"             // [
	END               Token = "END"               // ]
)

type Scanner struct {
	CurrentLine     string
	CurrentPosition int
}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) AssignLine(newLine string) []Token {
	s.CurrentPosition = 0
	s.CurrentLine = newLine

	return s.parseTokens()
}

func (s *Scanner) parseTokens() []Token {
	tokens := make([]Token, 0)

	var char rune

	for s.CurrentPosition, char = range s.CurrentLine {
		token := s.nextChar(char)
		if token == UNKNOWN_TOKEN {
			// TODO: LOG UNKNOWN_TOKEN
			continue
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func (s *Scanner) nextChar(char rune) Token {
	switch char {
	case '.':
		return OUTPUT
	case ',':
		return INPUT
	case '+':
		return INCREMENT
	case '-':
		return DECREMENT
	case '>':
		return INCREMENT_POINTER
	case '<':
		return DECREMENT_POINTER
	case '[':
		return BEGIN
	case ']':
		return END
	default:
		return UNKNOWN_TOKEN
	}
}
