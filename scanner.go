package gofelex

import (
  "bufio"
  "bytes"
  "io"
  "strings"
)


// Main object is a Scanner with its read and unread function

type Scanner struct {
  r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
  return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
  ch, _, err := s.r.ReadRune()
  if err != nil {
    return eof
  }
  return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }


// Scan() returns the next token and literal value
func (s *Scanner) Scan() (tok Token, lit string) {

  ch := s.read()

  // Consume whitespace and characters in blocks
  if isWhitespace(ch) {
    s.unread()
    return s.scanWhitespace()
  } else if isLetter(ch) {
    s.unread()
    return s.scanIdent()
  } else if isOperator(ch) {
    s.unread()
    return s.scanOperator()
  } else {
    return ILLEGAL, string(ch)
  }

}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {

  var buf bytes.Buffer
  buf.WriteRune(s.read())

  for {
    if ch := s.read(); ch == eof {
      break
    } else if !isWhitespace(ch) {
      s.unread()
      break
    } else {
      buf.WriteRune(ch)
    }
  }

  return WS, buf.String()
}

func (s *Scanner) scanIdent() (tok Token, lit string) {

  var buf bytes.Buffer
  buf.WriteRune(s.read())

  for {
    if ch := s.read(); ch == eof {
      break
    } else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
      s.unread()
      break
    } else {
      buf.WriteRune(ch)
    }
  }

  switch strings.ToUpper(buf.String()) {
  case "WITHIN":
    return WITHIN, buf.String()
  case "PRECEDES":
    return PRECEDES, buf.String()
  case "AND":
    return AND, buf.String()
  case "OR":
    return OR, buf.String()
  case "IN":
    return IN, buf.String()
  case "FLOWSTO*":
    return FLOWSTO, buf.String()
  }

  return IDENT, buf.String()
}

func (s *Scanner) scanOperator() (tok Token, lit string) {
  ch := s.read()

  switch ch {
  case eof:
    return EOF, ""
  default:
    return ILLEGAL, string(ch)
  }

}

// Utility functions

func isWhitespace(ch rune) bool {
  return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
  return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
  return (ch >= '0' && ch <= '9')
}



func isOperator(ch rune) bool {
 return (ch == eof)
}

var eof = rune(0)


