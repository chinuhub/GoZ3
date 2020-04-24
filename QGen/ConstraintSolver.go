package main

import (
	"fmt"
	z3 "go-z3"
)

func main() {
	//fmt.Printf("hello world")
	// Create the context

	config := z3.NewConfig()
	ctx := z3.NewContext(config)
	config.Close()
	defer ctx.Close()

	// Logic:
	//x<y
	//x>10 and x<100, y>10 and y<100 (to represent that both x and y are two digit numbers

	// Create the solver
	s := ctx.NewSolver()
	defer s.Close()

	// Vars
	x := ctx.Const(ctx.Symbol("x"), ctx.IntSort())
	y := ctx.Const(ctx.Symbol("y"), ctx.IntSort())
	hundred:= ctx.Int(1000,ctx.IntSort());
	ten:=ctx.Int(100,ctx.IntSort());

	//x<y
	formula:=x.Lt(y);
	s.Assert(formula);
	//x<100
	formula =x.Lt(hundred)
	s.Assert(formula);
	//y<100
	formula = y.Lt(hundred);
	s.Assert(formula);
	//x>10
	formula = x.Gt(ten);
	s.Assert(formula);
	//y>10
	formula = y.Gt(ten);
	s.Assert(formula);

	if v := s.Check(); v != z3.True {
		fmt.Println("Unsolveable")
		return
	}

	// Get the resulting model:
	m := s.Model()
	assignments := m.Assignments()
	m.Close()
	fmt.Printf("x = %s\n", assignments["x"])
	fmt.Printf("y = %s\n", assignments["y"])

	//Second solution
	//add another assertion saying that x!=what was assigned now, y!=what was assigned now and solve again.
	c := x.Eq(assignments["x"]).Not();
	s.Assert(c);
	c = y.Eq(assignments["y"]).Not();
	s.Assert(c);
	if v := s.Check(); v != z3.True {
		fmt.Println("Unsolveable")
		return
	}
	// Get the resulting model:
	s.Model()
	m = s.Model()
	assignments = m.Assignments()
	m.Close()
	fmt.Printf("x = %s\n", assignments["x"])
	fmt.Printf("y = %s\n", assignments["y"])

	//Third solution.
	//add another assertion saying that x!=what was assigned now, y!=what was assigned now and solve again.
	c = x.Eq(assignments["x"]).Not();
	s.Assert(c);
	c = y.Eq(assignments["y"]).Not();
	s.Assert(c);
	if v := s.Check(); v != z3.True {
		fmt.Println("Unsolveable")
		return
	}
	// Get the resulting model:
	s.Model()
	m = s.Model()
	assignments = m.Assignments()
	m.Close()
	fmt.Printf("x = %s\n", assignments["x"])
	fmt.Printf("y = %s\n", assignments["y"])
}

