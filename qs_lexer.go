// Code generated by golex. DO NOT EDIT.

package inzure

import (
	"fmt"
	"strconv"
	"unicode"
)

type qsLexer struct {
	src  string
	buf  []byte
	c    int
	at   int
	size int
	errs []error

	result QueryString
}

func newLexer(src string) *qsLexer {
	lex := &qsLexer{
		src:  src,
		at:   0,
		size: len(src),
		buf:  make([]byte, 0, 16),
	}
	lex.next()
	return lex
}

func (l *qsLexer) err(s string, arg ...interface{}) {
	eMsg := fmt.Sprintf("\n%s at char %d in ", fmt.Sprintf(s, arg...), l.at-1)
	pString := ""
	for i := 0; i < len(eMsg); i++ {
		pString += " "
	}
	for i := 0; i < l.at-2; i++ {
		pString += "~"
	}
	pString += "^"
	err := fmt.Errorf("%s%s\n%s", eMsg, l.src, pString)
	if l.errs == nil {
		l.errs = make([]error, 0, 1)
	}
	l.errs = append(l.errs, err)
}

func (l *qsLexer) Error(s string) {
	l.err(s)
}

func (l *qsLexer) next() int {
	if l.c != 0 {
		l.buf = append(l.buf, byte(l.c))
	}
	l.c = 0
	if l.at < l.size {
		l.c = int(l.src[l.at])
		l.at++
	}
	return l.c
}

func (l *qsLexer) Lex(sym *yySymType) int {
	c := l.c

yystate0:

	l.buf = l.buf[:0]

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = l.next()
yystart1:
	switch {
	default:
		goto yyrule19
	case c == '!':
		goto yystate4
	case c == '"':
		goto yystate12
	case c == '&':
		goto yystate16
	case c == '(':
		goto yystate18
	case c == ')':
		goto yystate19
	case c == '*' || c == 'A' || c >= 'C' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate20
	case c == ',':
		goto yystate21
	case c == '-':
		goto yystate22
	case c == '.':
		goto yystate24
	case c == '/':
		goto yystate26
	case c == '<' || c == '>':
		goto yystate27
	case c == '=':
		goto yystate28
	case c == 'B':
		goto yystate29
	case c == 'L':
		goto yystate57
	case c == '[':
		goto yystate61
	case c == '\t' || c == ' ':
		goto yystate3
	case c == '\x00':
		goto yystate2
	case c == ']':
		goto yystate62
	case c == 'f':
		goto yystate63
	case c == 'l':
		goto yystate68
	case c == 't':
		goto yystate71
	case c == '|':
		goto yystate75
	case c >= '0' && c <= '9':
		goto yystate23
	}

yystate2:
	c = l.next()
	goto yyrule1

yystate3:
	c = l.next()
	switch {
	default:
		goto yyrule2
	case c == '\t' || c == ' ':
		goto yystate3
	}

yystate4:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate5
	case c == 'L':
		goto yystate6
	case c == 'l':
		goto yystate9
	}

yystate5:
	c = l.next()
	goto yyrule18

yystate6:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'I':
		goto yystate7
	}

yystate7:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'K':
		goto yystate8
	}

yystate8:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'E':
		goto yystate5
	}

yystate9:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'i':
		goto yystate10
	}

yystate10:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'k':
		goto yystate11
	}

yystate11:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate5
	}

yystate12:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate13
	case c == '\\':
		goto yystate14
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate12
	}

yystate13:
	c = l.next()
	goto yyrule16

yystate14:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate15
	case c == '\\':
		goto yystate14
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate12
	}

yystate15:
	c = l.next()
	switch {
	default:
		goto yyrule16
	case c == '"':
		goto yystate13
	case c == '\\':
		goto yystate14
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate12
	}

yystate16:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '&':
		goto yystate17
	}

yystate17:
	c = l.next()
	goto yyrule8

yystate18:
	c = l.next()
	goto yyrule11

yystate19:
	c = l.next()
	goto yyrule12

yystate20:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate21:
	c = l.next()
	goto yyrule6

yystate22:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	case c >= '0' && c <= '9':
		goto yystate23
	}

yystate23:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == '*' || c == '-' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	case c >= '0' && c <= '9':
		goto yystate23
	}

yystate24:
	c = l.next()
	switch {
	default:
		goto yyrule5
	case c >= 'A' && c <= 'Z':
		goto yystate25
	}

yystate25:
	c = l.next()
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate25
	}

yystate26:
	c = l.next()
	goto yyrule13

yystate27:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c == '=':
		goto yystate5
	}

yystate28:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate5
	}

yystate29:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate20
	case c == 'o':
		goto yystate30
	}

yystate30:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate20
	case c == 'o':
		goto yystate31
	}

yystate31:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate20
	case c == 'l':
		goto yystate32
	}

yystate32:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'M' || c >= 'O' && c <= 'S' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	case c == 'F':
		goto yystate33
	case c == 'N':
		goto yystate38
	case c == 'T':
		goto yystate49
	case c == 'U':
		goto yystate51
	}

yystate33:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate20
	case c == 'a':
		goto yystate34
	}

yystate34:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate20
	case c == 'l':
		goto yystate35
	}

yystate35:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate20
	case c == 's':
		goto yystate36
	}

yystate36:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate20
	case c == 'e':
		goto yystate37
	}

yystate37:
	c = l.next()
	switch {
	default:
		goto yyrule3
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate38:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate20
	case c == 'o':
		goto yystate39
	}

yystate39:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate20
	case c == 't':
		goto yystate40
	}

yystate40:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	case c == 'A':
		goto yystate41
	}

yystate41:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate20
	case c == 'p':
		goto yystate42
	}

yystate42:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate20
	case c == 'p':
		goto yystate43
	}

yystate43:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate20
	case c == 'l':
		goto yystate44
	}

yystate44:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate20
	case c == 'i':
		goto yystate45
	}

yystate45:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate20
	case c == 'c':
		goto yystate46
	}

yystate46:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate20
	case c == 'a':
		goto yystate47
	}

yystate47:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate20
	case c == 'b':
		goto yystate48
	}

yystate48:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate20
	case c == 'l':
		goto yystate36
	}

yystate49:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate20
	case c == 'r':
		goto yystate50
	}

yystate50:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate20
	case c == 'u':
		goto yystate36
	}

yystate51:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate20
	case c == 'n':
		goto yystate52
	}

yystate52:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate20
	case c == 'k':
		goto yystate53
	}

yystate53:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate20
	case c == 'n':
		goto yystate54
	}

yystate54:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate20
	case c == 'o':
		goto yystate55
	}

yystate55:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate20
	case c == 'w':
		goto yystate56
	}

yystate56:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate20
	case c == 'n':
		goto yystate37
	}

yystate57:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	case c == 'I':
		goto yystate58
	}

yystate58:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	case c == 'K':
		goto yystate59
	}

yystate59:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	case c == 'E':
		goto yystate60
	}

yystate60:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate61:
	c = l.next()
	goto yyrule9

yystate62:
	c = l.next()
	goto yyrule10

yystate63:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate20
	case c == 'a':
		goto yystate64
	}

yystate64:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate20
	case c == 'l':
		goto yystate65
	}

yystate65:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate20
	case c == 's':
		goto yystate66
	}

yystate66:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate20
	case c == 'e':
		goto yystate67
	}

yystate67:
	c = l.next()
	switch {
	default:
		goto yyrule15
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate68:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate20
	case c == 'i':
		goto yystate69
	}

yystate69:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate20
	case c == 'k':
		goto yystate70
	}

yystate70:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate20
	case c == 'e':
		goto yystate60
	}

yystate71:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate20
	case c == 'r':
		goto yystate72
	}

yystate72:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate20
	case c == 'u':
		goto yystate73
	}

yystate73:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate20
	case c == 'e':
		goto yystate74
	}

yystate74:
	c = l.next()
	switch {
	default:
		goto yyrule14
	case c == '*' || c == '-' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate20
	}

yystate75:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '|':
		goto yystate76
	}

yystate76:
	c = l.next()
	goto yyrule7

yyrule1: // \0
	{
		return 0
	}
yyrule2: // [ \t]+

	goto yystate0
yyrule3: // {unknown_bool}
	{
		sym.ub = ubFromString(string(l.buf))
		return UNKNOWN_BOOL
		goto yystate0
	}
yyrule4: // {field}
	{
		sym.s = string(l.buf[1:])
		return FIELD
		goto yystate0
	}
yyrule5: // "."
	{
		return int('.')
	}
yyrule6: // ","
	{
		return int(',')
	}
yyrule7: // "||"
	{
		return OR
	}
yyrule8: // "&&"
	{
		return AND
	}
yyrule9: // "["
	{
		return OBRA
	}
yyrule10: // "]"
	{
		return CBRA
	}
yyrule11: // "("
	{
		return OPAR
	}
yyrule12: // ")"
	{
		return CPAR
	}
yyrule13: // "/"
	{
		return int('/')
	}
yyrule14: // "true"
	{
		sym.b = true
		return BOOL
		goto yystate0
	}
yyrule15: // "false"
	{
		sym.b = false
		return BOOL
		goto yystate0
	}
yyrule16: // {str}
	{
		return l.str(sym)
	}
yyrule17: // {number}
	{
		return l.int(sym)
	}
yyrule18: // {op}
	{
		sym.s = string(l.buf)
		return l.op(sym)
		goto yystate0
	}
yyrule19: // {chars}
	{
		sym.s = string(l.buf)
		return CHARS
		goto yystate0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	return int(unicode.ReplacementChar)
}

func (l *qsLexer) field(sym *yySymType) int {
	return FIELD
}

func (l *qsLexer) str(sym *yySymType) int {
	s := string(l.buf)
	s, err := strconv.Unquote(s)
	if err != nil {
		// TODO Error
		panic(err)
		return int(unicode.ReplacementChar)
	}
	sym.s = s
	return STR
}

func (l *qsLexer) unknownBool(sym *yySymType) int {
	sym.ub = ubFromString(string(l.buf))
	return UNKNOWN_BOOL
}

func (l *qsLexer) int(sym *yySymType) int {
	n, err := strconv.ParseInt(string(l.buf), 0, 64)
	if err != nil {
		// TODO Error
		panic(err)
		return int(unicode.ReplacementChar)
	}
	sym.i = n
	return NUMBER
}

func (l *qsLexer) op(sym *yySymType) int {
	ops := string(l.buf)
	var op QSOpT
	switch ops {
	case "!=":
		op = QSOpNe
	case "==":
		op = QSOpEq
	case ">=":
		op = QSOpGte
	case "<=":
		op = QSOpLte
	case ">":
		op = QSOpGt
	case "<":
		op = QSOpLt
	case "like", "LIKE":
		op = QSOpLike
	case "!like", "!LIKE":
		op = QSOpNotLike
	}
	sym.op = op
	return OP
}