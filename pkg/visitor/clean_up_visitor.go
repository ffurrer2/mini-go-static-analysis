// SPDX-License-Identifier: MIT
package visitor

import (
	"fmt"
	"go/ast"
)

func containsNoi(node ast.Node) bool {
	noiSearchVisitor := new(NoiSearchVisitor)
	noiSearchVisitor.Found = false
	ast.Walk(noiSearchVisitor, node)
	return noiSearchVisitor.Found
}

type CleanUpVisitor struct {
	File *ast.File
}

func (v *CleanUpVisitor) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil {
		return nil
	}

	switch t := node.(type) {

	// Comments and fields
	case *ast.Comment:
	// nothing to do

	case *ast.CommentGroup:
	// nothing to do

	case *ast.Field:
		t.Doc = nil

		names := []*ast.Ident{}
		for _, identNode := range t.Names {
			if containsNoi(identNode) {
				names = append(names, identNode)
			}
		}
		t.Names = append([]*ast.Ident(nil), names...)

		if !containsNoi(t.Tag) {
			t.Tag = nil
		}

		t.Comment = nil

	case *ast.FieldList:
		list := []*ast.Field{}
		for _, fieldNode := range t.List {
			if containsNoi(fieldNode) {
				list = append(list, fieldNode)
			}
		}
		t.List = append([]*ast.Field(nil), list...)

	// Expressions
	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
	// nothing to do

	case *ast.Ellipsis:
		if !containsNoi(t.Elt) {
			t.Elt = nil
		}

	case *ast.FuncLit:
	// nothing to do

	case *ast.CompositeLit:

		if !containsNoi(t.Type) {
			t.Type = nil
		}

		elts := []ast.Expr{}
		for _, exprNode := range t.Elts {
			if containsNoi(exprNode) {
				elts = append(elts, exprNode)
			}
		}
		t.Elts = append([]ast.Expr(nil), elts...)

	case *ast.ParenExpr:
	// nothing to do

	case *ast.SelectorExpr:
	// nothing to do

	case *ast.IndexExpr:
	// nothing to do

	case *ast.SliceExpr:
		if !containsNoi(t.Low) {
			t.Low = nil
		}

		if !containsNoi(t.High) {
			t.High = nil
		}

		if !containsNoi(t.Max) {
			t.Max = nil
		}

	case *ast.TypeAssertExpr:
	// nothing to do

	case *ast.CallExpr:
		args := []ast.Expr{}
		for _, exprNode := range t.Args {
			if containsNoi(exprNode) {
				args = append(args, exprNode)
			}
		}
		t.Args = append([]ast.Expr(nil), args...)

	case *ast.StarExpr:
	// nothing to do

	case *ast.UnaryExpr:
	// nothing to do

	case *ast.BinaryExpr:
	// nothing to do

	case *ast.KeyValueExpr:
	// nothing to do

	// Types
	case *ast.ArrayType:
		// nothing to do
		return v

	case *ast.StructType:
		// nothing to do
		return v

	case *ast.FuncType:
	// nothing to do

	case *ast.InterfaceType:
	// nothing to do

	case *ast.MapType:
	// nothing to do

	case *ast.ChanType:
	// nothing to do

	// Statements
	case *ast.BadStmt:
	// nothing to do

	case *ast.DeclStmt:
	// nothing to do

	case *ast.EmptyStmt:
	// nothing to do

	case *ast.LabeledStmt:
	// nothing to do

	case *ast.ExprStmt:
	// nothing to do

	case *ast.SendStmt:
	// nothing to do

	case *ast.IncDecStmt:
	// nothing to do

	case *ast.AssignStmt:
	// nothing to do

	case *ast.GoStmt:
	// nothing to do

	case *ast.DeferStmt:
	// nothing to do

	case *ast.ReturnStmt:
		results := []ast.Expr{}
		for _, stmtNode := range t.Results {
			if containsNoi(stmtNode) {
				results = append(results, stmtNode)
			}
		}
		t.Results = append([]ast.Expr(nil), results...)

	case *ast.BranchStmt:
		if !containsNoi(t.Label) {
			t.Label = nil
		}

	case *ast.BlockStmt:
		list := []ast.Stmt{}
		for _, stmtNode := range t.List {
			if containsNoi(stmtNode) {
				list = append(list, stmtNode)
			}
		}
		t.List = append([]ast.Stmt(nil), list...)

	case *ast.IfStmt:
		if !containsNoi(t.Init) {
			t.Init = nil
		}

		if !containsNoi(t.Else) {
			t.Else = nil
		}

	case *ast.CaseClause:
		list := []ast.Expr{}
		for _, exprNode := range t.List {
			if containsNoi(exprNode) {
				list = append(list, exprNode)
			}
		}
		t.List = append([]ast.Expr(nil), list...)

		body := []ast.Stmt{}
		for _, stmtNode := range t.Body {
			if containsNoi(stmtNode) {
				body = append(body, stmtNode)
			}
		}
		t.Body = append([]ast.Stmt(nil), body...)

	case *ast.SwitchStmt:
		if !containsNoi(t.Init) {
			t.Init = nil
		}

		if !containsNoi(t.Tag) {
			t.Tag = nil
		}

	case *ast.TypeSwitchStmt:
		if !containsNoi(t.Init) {
			t.Init = nil
		}

	case *ast.CommClause:
	// nothing to do

	case *ast.SelectStmt:
	// nothing to do

	case *ast.ForStmt:
		if !containsNoi(t.Init) {
			t.Init = nil
		}

		if !containsNoi(t.Cond) {
			t.Cond = nil
		}

		if !containsNoi(t.Post) {
			t.Post = nil
		}

	case *ast.RangeStmt:
		if !containsNoi(t.Key) {
			t.Key = nil
		}

		if !containsNoi(t.Value) {
			t.Key = nil
		}

	// Declarations
	case *ast.ImportSpec:
		t.Doc = nil

		if !containsNoi(t.Name) {
			t.Name = nil
		}

		t.Comment = nil

	case *ast.ValueSpec:
		t.Doc = nil

		names := []*ast.Ident{}
		for _, identNode := range t.Names {
			if containsNoi(identNode) {
				names = append(names, identNode)
			}
		}
		t.Names = append([]*ast.Ident(nil), names...)

		if !containsNoi(t.Type) {
			t.Type = nil
		}

		values := []ast.Expr{}
		for _, exprNode := range t.Values {
			if containsNoi(exprNode) {
				values = append(values, exprNode)
			}
		}
		t.Values = append([]ast.Expr(nil), values...)

		t.Comment = nil

	case *ast.TypeSpec:
		t.Doc = nil
		t.Comment = nil

	case *ast.BadDecl:
	// nothing to do

	case *ast.GenDecl:
		t.Doc = nil

		specs := []ast.Spec{}
		for _, specsNode := range t.Specs {
			if containsNoi(specsNode) {
				specs = append(specs, specsNode)
			}
		}
		t.Specs = append([]ast.Spec(nil), specs...)

	case *ast.FuncDecl:
		t.Doc = nil

	// Files and packages
	case *ast.File:
		t.Doc = nil
		t.Scope = nil
		t.Unresolved = nil
		t.Comments = nil

		decls := []ast.Decl{}
		for _, declNode := range t.Decls {
			if containsNoi(declNode) {
				decls = append(decls, declNode)
			}
		}
		t.Decls = append([]ast.Decl(nil), decls...)

	case *ast.Package:
	// nothing to do

	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", t))
	}

	return v
}
