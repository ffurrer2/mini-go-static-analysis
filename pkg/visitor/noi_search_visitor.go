// SPDX-License-Identifier: MIT
package visitor

import (
	"go/ast"
	"go/token"
)

type NoiSearchVisitor struct {
	Found bool
}

func (v *NoiSearchVisitor) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil || v.Found {
		return nil
	}

	switch t := node.(type) {

	case *ast.GoStmt:
		v.Found = true
		return nil

	case *ast.SendStmt:
		v.Found = true
		return nil

	case *ast.UnaryExpr:
		if t.Op == token.ARROW {
			v.Found = true
			return nil
		}

	case *ast.CallExpr:
		v.Found = true
		return nil

	case *ast.ChanType:
		v.Found = true
		return nil
	default:
		v.Found = false
		return v
	}
	return v
}
