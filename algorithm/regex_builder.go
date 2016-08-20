package algorithm

import (
	"fmt"
	"go/ast"
	"github.com/ffurrer2/mini-go-static-analysis/regex"
	"go/token"
)

func Walk(v ast.Visitor, node ast.Node) {
	if v = v.Visit(node); v == nil {
		return
	}
}

type RegexBuilder struct {
	Root       *regex.Root
	ActualNode regex.ExprStmt
	Stack      Stack
}

// Helper functions for common node lists. They may be empty.

func walkIdentList(v ast.Visitor, list []*ast.Ident) {
	for _, x := range list {
		Walk(v, x)
	}
}

func walkExprList(v ast.Visitor, list []ast.Expr) {
	for _, x := range list {
		Walk(v, x)
	}
}

func walkStmtList(v ast.Visitor, list []ast.Stmt) {
	for _, x := range list {
		Walk(v, x)
	}
}

func walkDeclList(v ast.Visitor, list []ast.Decl) {
	for _, x := range list {
		Walk(v, x)
	}
}

func (v *RegexBuilder) Visit(node ast.Node) (w ast.Visitor) {

	if node == nil {
		return nil
	}

	// walk children
	// (the order of the cases matches the order
	// of the corresponding node types in ast.go)
	switch t := node.(type) {
	// Comments and fields
	case *ast.Comment:
	// nothing to do

	case *ast.CommentGroup:
		for _, c := range t.List {
			Walk(v, c)
		}

	case *ast.Field:
		if t.Doc != nil {
			Walk(v, t.Doc)
		}
		walkIdentList(v, t.Names)
		Walk(v, t.Type)
		if t.Tag != nil {
			Walk(v, t.Tag)
		}
		if t.Comment != nil {
			Walk(v, t.Comment)
		}

	case *ast.FieldList:
		for _, f := range t.List {
			Walk(v, f)
		}

	// Expressions
	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
	// nothing to do

	case *ast.Ellipsis:
		if t.Elt != nil {
			Walk(v, t.Elt)
		}

	case *ast.FuncLit:
		Walk(v, t.Type)
		Walk(v, t.Body)

	case *ast.CompositeLit:
		if t.Type != nil {
			Walk(v, t.Type)
		}
		walkExprList(v, t.Elts)

	case *ast.ParenExpr:
		Walk(v, t.X)

	case *ast.SelectorExpr:
		Walk(v, t.X)
		Walk(v, t.Sel)

	case *ast.IndexExpr:
		Walk(v, t.X)
		Walk(v, t.Index)

	case *ast.SliceExpr:
		Walk(v, t.X)
		if t.Low != nil {
			Walk(v, t.Low)
		}
		if t.High != nil {
			Walk(v, t.High)
		}
		if t.Max != nil {
			Walk(v, t.Max)
		}

	case *ast.TypeAssertExpr:
		Walk(v, t.X)
		if t.Type != nil {
			Walk(v, t.Type)
		}

	case *ast.CallExpr:

		v.ActualNode.AppendStmt(regex.NewCallStmt(t))

		Walk(v, t.Fun)
		walkExprList(v, t.Args)

	case *ast.StarExpr:
		Walk(v, t.X)

	case *ast.UnaryExpr:

		Walk(v, t.X)

		if t.Op == token.ARROW {
			v.ActualNode.AppendStmt(regex.NewRcvStmt(t))
		}

	case *ast.BinaryExpr:
		Walk(v, t.X)
		Walk(v, t.Y)

	case *ast.KeyValueExpr:
		Walk(v, t.Key)
		Walk(v, t.Value)

	// Types
	case *ast.ArrayType:
		if t.Len != nil {
			Walk(v, t.Len)
		}
		Walk(v, t.Elt)

	case *ast.StructType:
		Walk(v, t.Fields)

	case *ast.FuncType:
		if t.Params != nil {
			Walk(v, t.Params)
		}
		if t.Results != nil {
			Walk(v, t.Results)
		}

	case *ast.InterfaceType:
		Walk(v, t.Methods)

	case *ast.MapType:
		Walk(v, t.Key)
		Walk(v, t.Value)

	case *ast.ChanType:
		Walk(v, t.Value)

	// Statements
	case *ast.BadStmt:
	// nothing to do

	case *ast.DeclStmt:
		Walk(v, t.Decl)

	case *ast.EmptyStmt:
	// nothing to do

	case *ast.LabeledStmt:
		Walk(v, t.Label)
		Walk(v, t.Stmt)

	case *ast.ExprStmt:
		Walk(v, t.X)

	case *ast.SendStmt:

		Walk(v, t.Chan)
		Walk(v, t.Value)

		v.ActualNode.AppendStmt(regex.NewSndStmt(t))

	case *ast.IncDecStmt:
		Walk(v, t.X)

	case *ast.AssignStmt:
		walkExprList(v, t.Lhs)
		walkExprList(v, t.Rhs)

	case *ast.GoStmt:

		// Save ActualNode
		v.Stack.Push(v.ActualNode)

		newFork := regex.NewFork(t, regex.NewFunc(nil))
		v.Root.AppendFork(newFork)
		v.ActualNode = newFork.Func.Regex
		Walk(v, t.Call)

		// Restore ActualNode
		v.ActualNode = v.Stack.Pop().(regex.ExprStmt)

	case *ast.DeferStmt:
		Walk(v, t.Call)

	case *ast.ReturnStmt:
		walkExprList(v, t.Results)

	case *ast.BranchStmt:
		if t.Label != nil {
			Walk(v, t.Label)
		}

	case *ast.BlockStmt:
		walkStmtList(v, t.List)

	case *ast.IfStmt:

		if t.Init != nil {
			Walk(v, t.Init)
		}

		Walk(v, t.Cond)

		// Save ActualNode
		v.Stack.Push(v.ActualNode)

		altExpr := new(regex.AltExpr)
		v.ActualNode.AppendStmt(altExpr)
		v.ActualNode = altExpr

		Walk(v, t.Body)
		if t.Else != nil {
			Walk(v, t.Else)
		} else {
			altExpr.AppendStmt(regex.NewEpsStmt())
		}

		// Restore ActualNode
		v.ActualNode = v.Stack.Pop().(regex.ExprStmt)

	case *ast.CaseClause:

		// Save ActualNode
		v.Stack.Push(v.ActualNode)

		newConcatExpr := new(regex.ConcatExpr)
		v.ActualNode.AppendStmt(newConcatExpr)
		v.ActualNode = newConcatExpr

		walkExprList(v, t.List)
		walkStmtList(v, t.Body)

		// Restore ActualNode
		v.ActualNode = v.Stack.Pop().(regex.ExprStmt)


	case *ast.SwitchStmt:

		if t.Init != nil {
			Walk(v, t.Init)
		}

		if t.Tag != nil {
			Walk(v, t.Tag)
		}

		// Save ActualNode
		v.Stack.Push(v.ActualNode)

		altExpr := new(regex.AltExpr)
		v.ActualNode.AppendStmt(altExpr)
		v.ActualNode = altExpr

		Walk(v, t.Body)

		// Restore ActualNode
		v.ActualNode = v.Stack.Pop().(regex.ExprStmt)

	case *ast.TypeSwitchStmt:
		if t.Init != nil {
			Walk(v, t.Init)
		}
		Walk(v, t.Assign)
		Walk(v, t.Body)

	case *ast.CommClause:

		// Save ActualNode
		v.Stack.Push(v.ActualNode)

		newConcatExpr := new(regex.ConcatExpr)
		v.ActualNode.AppendStmt(newConcatExpr)
		v.ActualNode = newConcatExpr

		if t.Comm != nil {
			Walk(v, t.Comm)
		}
		walkStmtList(v, t.Body)

		// Restore ActualNode
		v.ActualNode = v.Stack.Pop().(regex.ExprStmt)

	case *ast.SelectStmt:

		// Save ActualNode
		v.Stack.Push(v.ActualNode)

		altExpr := new(regex.AltExpr)
		v.ActualNode.AppendStmt(altExpr)
		v.ActualNode = altExpr

		Walk(v, t.Body)

		// Restore ActualNode
		v.ActualNode = v.Stack.Pop().(regex.ExprStmt)

	case *ast.ForStmt:

		if t.Init != nil {
			Walk(v, t.Init)
		}

		if t.Cond != nil {
			Walk(v, t.Cond)
		}

		// Save ActualNode
		v.Stack.Push(v.ActualNode)
		altExpr := new(regex.AltExpr)
		v.ActualNode.AppendStmt(altExpr)
		v.ActualNode = altExpr

		// Save ActualNode
		v.Stack.Push(v.ActualNode)
		newConcatExpr := new(regex.ConcatExpr)
		v.ActualNode.AppendStmt(newConcatExpr)
		v.ActualNode = newConcatExpr

		Walk(v, t.Body)

		// Save ActualNode
		v.Stack.Push(v.ActualNode)
		newStarExpr := new(regex.StarExpr)
		v.ActualNode.AppendStmt(newStarExpr)
		v.ActualNode = newStarExpr

		newConcatExpr2 := new(regex.ConcatExpr)
		v.ActualNode.AppendStmt(newConcatExpr2)
		v.ActualNode = newConcatExpr2

		if t.Post != nil {
			Walk(v, t.Post)
		}
		if t.Cond != nil {
			Walk(v, t.Cond)
		}

		Walk(v, t.Body)

		// Restore ActualNode
		v.ActualNode = v.Stack.Pop().(regex.ExprStmt)

		if t.Post != nil {
			Walk(v, t.Post)
		}
		if t.Cond != nil {
			Walk(v, t.Cond)
		}

		// Restore ActualNode
		v.ActualNode = v.Stack.Pop().(regex.ExprStmt)

		altExpr.AppendStmt(regex.NewEpsStmt())

		// Restore ActualNode
		v.ActualNode = v.Stack.Pop().(regex.ExprStmt)

	case *ast.RangeStmt:
		if t.Key != nil {
			Walk(v, t.Key)
		}
		if t.Value != nil {
			Walk(v, t.Value)
		}
		Walk(v, t.X)
		Walk(v, t.Body)

	// Declarations
	case *ast.ImportSpec:
		if t.Doc != nil {
			Walk(v, t.Doc)
		}
		if t.Name != nil {
			Walk(v, t.Name)
		}
		Walk(v, t.Path)
		if t.Comment != nil {
			Walk(v, t.Comment)
		}

	case *ast.ValueSpec:
		if t.Doc != nil {
			Walk(v, t.Doc)
		}
		walkIdentList(v, t.Names)
		if t.Type != nil {
			Walk(v, t.Type)
		}
		walkExprList(v, t.Values)
		if t.Comment != nil {
			Walk(v, t.Comment)
		}

	case *ast.TypeSpec:
		if t.Doc != nil {
			Walk(v, t.Doc)
		}
		Walk(v, t.Name)
		Walk(v, t.Type)
		if t.Comment != nil {
			Walk(v, t.Comment)
		}

	case *ast.BadDecl:
	// nothing to do

	case *ast.GenDecl:
		if t.Doc != nil {
			Walk(v, t.Doc)
		}
		for _, s := range t.Specs {
			Walk(v, s)
		}

	case *ast.FuncDecl:

		// Create and add a Func with *ast.FuncDecl t to Root
		newFunc := regex.NewFunc(t)
		v.Root.AppendFunc(newFunc)

		if t.Name.Name == "main" {
			v.Root.SetMain(newFunc)
		}
		v.ActualNode = newFunc.Regex

		if t.Doc != nil {
			Walk(v, t.Doc)
		}
		if t.Recv != nil {
			Walk(v, t.Recv)
		}
		Walk(v, t.Name)
		Walk(v, t.Type)
		if t.Body != nil {
			Walk(v, t.Body)
		}

	// Files and packages
	case *ast.File:

		//Create and set root
		v.Root = regex.NewRoot(t)

		if t.Doc != nil {
			Walk(v, t.Doc)
		}
		Walk(v, t.Name)
		walkDeclList(v, t.Decls)
	// don't walk t.Comments - they have been
	// visited already through the individual
	// nodes

	case *ast.Package:
		for _, f := range t.Files {
			Walk(v, f)
		}

	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", t))
	}

	return v
}
