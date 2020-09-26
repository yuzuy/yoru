package object

import (
	"bytes"
	"fmt"
	"strings"

	"monkey/ast"
)

type Type string

const (
	NullObj        = "NULL"
	ErrorObj       = "ERROR"
	ReturnValueObj = "RETURN_VALUE"
	IntObj         = "INTEGER"
	BoolObj        = "BOOLEAN"
	FunctionObj    = "FUNCTION"
	StringObj      = "STRING"

	BuiltInObj = "BUILD-IN"
)

type Object interface {
	Type() Type
	Inspect() string
}

type Null struct{}

func (n *Null) Type() Type      { return NullObj }
func (n *Null) Inspect() string { return "null" }

type Error struct {
	Message string
}

func (e *Error) Type() Type      { return ErrorObj }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() Type      { return ReturnValueObj }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type      { return IntObj }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type      { return BoolObj }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type { return FunctionObj }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() Type      { return StringObj }
func (s *String) Inspect() string { return s.Value }

type BuiltIn struct {
	Fn BuiltInFunction
}

func (b *BuiltIn) Type() Type      { return BuiltInObj }
func (b *BuiltIn) Inspect() string { return "built-in function" }

type BuiltInFunction func(args ...Object) Object
