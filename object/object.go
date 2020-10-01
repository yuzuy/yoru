package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
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
	ArrayObj       = "ARRAY"
	HashObj        = "HASH"

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
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Value: h.Sum64()}
}

type Array struct {
	Elements []Object
}

func (a *Array) Type() Type { return ArrayObj }
func (a *Array) Inspect() string {
	var out bytes.Buffer

	var elements []string
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type Hashable interface {
	HashKey() HashKey
}

type HashKey struct {
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() Type { return HashObj }
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	var pairs []string
	for _, v := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", v.Key.Inspect(), v.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type BuiltIn struct {
	Fn BuiltInFunction
}

func (b *BuiltIn) Type() Type      { return BuiltInObj }
func (b *BuiltIn) Inspect() string { return "built-in function" }

type BuiltInFunction func(args ...Object) Object
