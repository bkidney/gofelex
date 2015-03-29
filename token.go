package gofelex

// Define Tokens
// TODO: Refactor out into interface
type Token int
const (
  // Special Tokens
  ILLEGAL Token = iota
  EOF
  WS

  // Literals
  IDENT

  // Operator Types 
  LOGICAL
  TEMPORAL
  CONDITION
  FLOW
)
