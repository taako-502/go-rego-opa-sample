package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/v1/rego"
)

func main() {
	// Sample request input
	input := map[string]interface{}{
		"user": map[string]interface{}{
			"role":   "user",
			"groups": []string{"developers"},
			"age":    17,
		},
		"action": "read",
	}

	ctx := context.Background()

	// Load policy from policy directory and prepare evaluator
	r := rego.New(
		rego.Query("data.demo.allow"),
		rego.Load([]string{"policy"}, nil),
	)

	pq, err := r.PrepareForEval(ctx)
	if err != nil {
		panic(err)
	}

	rs, err := pq.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		panic(err)
	}

	// Interpret result
	allowed := false
	if len(rs) > 0 {
		if v, ok := rs[0].Expressions[0].Value.(bool); ok {
			allowed = v
		}
	}

	if allowed {
		fmt.Println("Allowed")
	} else {
		fmt.Println("Denied")
	}
}
