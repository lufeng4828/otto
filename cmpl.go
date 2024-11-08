package otto

import (
	"github.com/lufeng4828/otto/ast"
	"github.com/lufeng4828/otto/file"
)

type compiler struct {
	file    *file.File
	program *ast.Program
}
