package go_test_summary

import (
	"go/types"
	"go/token"
)

func Compute(expr string) (string, error) {
	tv, err := types.Eval(token.NewFileSet(), types.NewPackage("main", "main"), token.NoPos, expr)
	if err != nil {
		return "", err
	}
	return tv.Value.String(), nil
}
