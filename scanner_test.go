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
    {s: `DBServerNode:myDB:openSession(_):?sessionID` , tok: gofelex.IDENT, lit: `DBServerNode:myDB:openSession(_):?sessionID`},
    {s: `?node:ip:send(?srcIP,?destIP)`, tok: gofelex.IDENT , lit: `?node:ip:send(?srcIP,?destIP)`},
    {s: `destIP`, tok: gofelex.IDENT, lit: `destIP`},
    {s: `[203.1.112.25]`, tok: gofelex.IDENT, lit: `[203.1.112.25]`},

    // Keywords
    {s: "Precedes", tok: gofelex.TEMPORAL, lit: "Precedes"},
    {s: "Within", tok: gofelex.TEMPORAL, lit: "Within"},
    {s: "and", tok: gofelex.LOGICAL, lit: "and"},
    {s: "or", tok: gofelex.LOGICAL, lit: "or"},
    {s: "in", tok: gofelex.CONDITION, lit: "in"},
    {s: `FlowsTo*`, tok: gofelex.FLOW, lit: `FlowsTo*`},
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
