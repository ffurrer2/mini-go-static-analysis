package algorithm

import (
	"bytes"
	"fmt"
	"github.com/ffurrer2/projektarbeit1/regex"
	"strings"
	"go/token"
	"path/filepath"
	"strconv"
)

type RegexPrinter struct {
	Buffer bytes.Buffer
	Fset   *token.FileSet
}

func (v *RegexPrinter) Visit(node regex.Stmt) (w regex.Visitor) {

	if node == nil {
		return nil
	}

	switch t := node.(type) {

	case *regex.SndStmt:

		v.Buffer.WriteString(fmt.Sprintf("!%s", v.posToString(t.Node.Pos())))

	case *regex.RcvStmt:

		v.Buffer.WriteString(fmt.Sprintf("?%s", v.posToString(t.Node.Pos())))

	case *regex.CloseStmt:

		v.Buffer.WriteString(fmt.Sprintf("#%s", v.posToString(t.Node.Pos())))

	case *regex.CallStmt:

		v.Buffer.WriteString(fmt.Sprintf("func():%s", v.posToString(t.Node.Pos())))

	case *regex.EpsStmt:

		v.Buffer.WriteString("ε")

	case *regex.ConcatExpr:

		v.Buffer.WriteString("(")
		for _, node := range t.List {
			regex.WalkRegex(v, node)
			v.Buffer.WriteString("⋅")
		}
		if strings.HasSuffix(v.Buffer.String(), "⋅") {
			v.Buffer.Truncate(v.Buffer.Len() - 3)
		}

		v.Buffer.WriteString(")")

	case *regex.AltExpr:
		v.Buffer.WriteString("(")
		for _, node := range t.List {
			regex.WalkRegex(v, node)
			v.Buffer.WriteString("+")
		}
		if strings.HasSuffix(v.Buffer.String(), "+") {
			v.Buffer.Truncate(v.Buffer.Len() - 1)
		}
		v.Buffer.WriteString(")")

	case *regex.StarExpr:
		v.Buffer.WriteString("(")
		regex.WalkRegex(v, t.Node)
		v.Buffer.WriteString(")*")

	default:
		panic(fmt.Sprintf("visitor.Visit: unexpected node type %T", t))
	}

	return v
}

func (v *RegexPrinter) posToString(pos token.Pos) string {

	position := v.Fset.Position(pos)
	posString := filepath.Base(v.Fset.File(pos).Name()) + ":" + strconv.Itoa(position.Line) + ":" + strconv.Itoa(position.Column)

	return posString
}
