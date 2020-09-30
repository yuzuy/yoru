package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		u, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", u.Name)
		fmt.Println("Feel free to type in commands")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	if ext := filepath.Ext(filename); ext != ".monkey" {
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
	log.SetPrefix("monkey: ")
}
