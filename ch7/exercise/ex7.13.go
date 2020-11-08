//@author: hdsfade
//@date: 2020-11-07-16:31
package eval

import "fmt"
//A Var identifies a variable
type Var string

//A literal is a numeric constant
type literal float64

//A unary represents a unary operator expression
type unary struct {
	op rune
	x  Expr
}

//A binary represents a binary operator expression
type binary struct {
	op   rune
	x, y Expr
}

//A call represents a function call expression
type call struct {
	fn   string
	args []Expr
}

//var->value
type Env map[Var]float64
//An Expr is an arithmetic  expression
type Expr interface {
	//Eval returns the value of this Expr in the enviroment env
	Eval(env Env) float64
	//Check reports errors in this Expr and adds its Vars to the set
	Check(vars map[Var]bool) error
	//String return a expression which is easy to read.
	Strings() string
}

func (u unary) String() string {
	switch u.op {
	case '+':
		return "+" + u.x.Strings()
	case '-':
		return "-" + u.x.Strings()
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) String() string {
	switch b.op {
	case '+':
		return b.x.Strings() + "+" + b.y.Strings()
	case '-':
		return b.x.Strings() + "-" + b.y.Strings()
	case '*':
		return b.x.Strings() + "*" + b.y.Strings()
	case '/':
		return b.x.Strings() + "/" + b.y.Strings()
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (v Var) String() string {
	return string(v)
}
func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}
func (c call) String() string {
	switch c.fn {
	case "pow":
		return "pow("+c.args[0].Strings()+","+c.args[1].Strings()+")"
	case "sin":
		return "sin("+c.args[0].Strings()+")"
	case "sqrt":
		return "sqrt("+c.args[0].Strings()+")"
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}
