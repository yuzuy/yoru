package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/yuzuy/yoru/evaluator"
	"github.com/yuzuy/yoru/lexer"
	"github.com/yuzuy/yoru/object"
	"github.com/yuzuy/yoru/parser"
	"github.com/yuzuy/yoru/repl"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		fmt.Printf("Hello! This is the Yoru programming language!\n")
		fmt.Println("Feel free to type in commands")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	if ext := filepath.Ext(filename); ext != ".yoru" {
		log.Printf("invalid file extension %q\n", ext)
		return
	}
	f, err := os.Open(filename)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err.Error())
		return
	}

	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			log.Println(err.Error())
		}
		return
	}
	env := object.NewEnvironment()
	if result := evaluator.Eval(program, env); result.Type() == object.ErrorObj {
		log.Println(result.Inspect())
	}
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("yoru: ")
}
