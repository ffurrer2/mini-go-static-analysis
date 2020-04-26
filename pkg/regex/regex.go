package regex

import "go/ast"

type Visitor interface {
	Visit(node Stmt) (w Visitor)
}

func WalkRegex(v Visitor, node Stmt) {
	if v = v.Visit(node); v == nil {
		return
	}
}

type Root struct {
	File     *ast.File
	Main     *Func
	ForkList []*Fork
	FuncList []*Func
}

type Fork struct {
	Smnt *ast.GoStmt
	Func *Func
}

type Func struct {
	FuncDecl *ast.FuncDecl // FuncDecl of function or nil
	Regex    ExprStmt
}

func NewRoot(f *ast.File) (r *Root) {
	r = new(Root)
	r.File = f
	r.Main = nil
	r.ForkList = []*Fork{}
	r.FuncList = []*Func{}
	return r
}

func NewFork(s *ast.GoStmt, fun *Func) (f *Fork) {
	f = new(Fork)
	f.Smnt = s
	f.Func = fun
	return f
}

func NewFunc(fd *ast.FuncDecl) (f *Func) {
	f = new(Func)
	f.FuncDecl = fd
	f.Regex = new(ConcatExpr)
	return f
}

func (r *Root) SetMain(m *Func) {
	r.Main = m
}

func (r *Root) AppendFork(f *Fork) {
	r.ForkList = append(r.ForkList, f)
}

func (r *Root) AppendFunc(f *Func) {
	r.FuncList = append(r.FuncList, f)
}

// ################################

type Stmt interface {
	stmt()
}

type SndStmt struct {
	Node *ast.SendStmt
}

type RcvStmt struct {
	Node *ast.UnaryExpr
}

type CloseStmt struct {
	Node *ast.CallExpr
}

type CallStmt struct {
	Node *ast.CallExpr
}

type EpsStmt struct {
}

// Method implementations for Stmt interface

func (*SndStmt) stmt()   {}
func (*RcvStmt) stmt()   {}
func (*CloseStmt) stmt() {}
func (*CallStmt) stmt()  {}
func (*EpsStmt) stmt()   {}

func (*ConcatExpr) stmt() {}
func (*AltExpr) stmt()    {}
func (*StarExpr) stmt()   {}

func NewSndStmt(n *ast.SendStmt) (s *SndStmt) {
	s = new(SndStmt)
	s.Node = n
	return s
}

func NewRcvStmt(n *ast.UnaryExpr) (s *RcvStmt) {
	s = new(RcvStmt)
	s.Node = n
	return s
}

func NewCloseStmt(n *ast.CallExpr) (s *CloseStmt) {
	s = new(CloseStmt)
	s.Node = n
	return s
}

func NewCallStmt(n *ast.CallExpr) (s *CallStmt) {
	s = new(CallStmt)
	s.Node = n
	return s
}

func NewEpsStmt() *EpsStmt {
	return new(EpsStmt)
}

// ################################

type ExprStmt interface {
	Stmt
	exprStmt()
	AppendStmt(s Stmt)
}

// Concatenation
type ConcatExpr struct {
	List []Stmt
}

// Alternation
type AltExpr struct {
	List []Stmt
}

// Kleene star
type StarExpr struct {
	Node Stmt
}

func (*ConcatExpr) exprStmt() {}
func (*AltExpr) exprStmt()    {}
func (*StarExpr) exprStmt()   {}

func (e *ConcatExpr) AppendStmt(s Stmt) {
	e.List = append(e.List, s)
}

func (e *AltExpr) AppendStmt(s Stmt) {
	e.List = append(e.List, s)
}

func (e *StarExpr) AppendStmt(s Stmt) {

	if e.Node == nil {
		e.Node = s
	} else {
		concatExpr := new(ConcatExpr)
		concatExpr.List = []Stmt{}
		concatExpr.AppendStmt(e.Node)
		e.Node = concatExpr
	}
}
