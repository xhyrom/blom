package transpiler

import "blom/ast"

type Transpiler interface {
	Transpile(program *ast.Program) (string, error)
	TranspileBlock(block ast.Block) string
	TranspileStatement(stmt ast.Statement) string
	TranspileFunctionDeclaration(declaration *ast.FunctionDeclaration) string
	TranspileBinaryExpression(expression *ast.BinaryExpression) string
	TranspileUnaryExpression(expression *ast.UnaryExpression) string
	TranspileDeclarationStatement(declaration *ast.VariableDeclaration) string
	TranspileReturnStatement(statement *ast.Return) string
	TranspileIfStatement(statement *ast.If) string
	TranspileFunctionCall(call *ast.FunctionCall) string
}
