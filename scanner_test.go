package gofelex_test

import (
  "strings"
  "testing"
  "github.com/bkidney/gofelex"
)

type testset struct {
    s string
    tok gofelex.Token
    lit string
}

// Ensure the scanner can scan tokens correctly
func TestScanner_Scan(t *testing.T) {
  var tests = []testset {
    // Special tokens (EOF, ILLEGAL, WS)
    {s: ``, tok: gofelex.EOF },
    {s: `#`, tok: gofelex.ILLEGAL, lit: `#`},
    {s: ` `, tok: gofelex.WS, lit: ` ` },
    {s: `   `, tok: gofelex.WS, lit: `   `},
    {s: "\t", tok: gofelex.WS, lit: "\t"},
    {s: "\n", tok: gofelex.WS, lit: "\n"},

    // Identifiers 
    //{s: , tok: , lit: },
    //{s: , tok: , lit: },
    //{s: , tok: , lit: },

    // Keywords
    {s: "Precedes", tok: gofelex.PRECEDES, lit: "Precedes"},
    {s: "Within", tok: gofelex.WITHIN, lit: "Within"},
    {s: "and", tok: gofelex.AND, lit: "and"},
    {s: "or", tok: gofelex.OR, lit: "or"},
    {s: "in", tok: gofelex.IN, lit: "in"},
    //{s: `FlowsTo*`, tok: gofelex.FLOWSTO, lit: `FlowsTo*`},
  }

  for i, tt := range tests {
    s := gofelex.NewScanner(strings.NewReader(tt.s))
    tok, lit := s.Scan()
    if tt.tok != tok {
      t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
    } else if tt.lit != lit {
      t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
    }
  }

}
