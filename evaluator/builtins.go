package evaluator

import "monkey/object"

var builtIns = map[string]*object.BuiltIn{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported. got %s", arg.Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d", len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("the first argument to `push` must be ARRAY. got=%s", args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)
			newElements := make([]object.Object, length, length+len(args[1:]))

			copy(newElements, array.Elements)
			newElements = append(newElements, args[1:]...)

			return &object.Array{Elements: newElements}
		},
	},
}
