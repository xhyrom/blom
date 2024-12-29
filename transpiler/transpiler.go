package transpiler

import "blom/ast"

type Transpiler interface {
	Transpile(program *ast.Program) (string, error)
	TranspileBlock(block ast.BlockStatement) string
	TranspileStatement(stmt ast.Statement) string
	TranspileFunctionDeclaration(declaration *ast.FunctionDeclaration) string
	TranspileBinaryExpression(expression *ast.BinaryExpression) string
	TranspileUnaryExpression(expression *ast.UnaryExpression) string
	TranspileDeclarationStatement(declaration *ast.VariableDeclarationStatement) string
	TranspileReturnStatement(statement *ast.ReturnStatement) string
	TranspileIfStatement(statement *ast.IfStatement) string
	TranspileFunctionCall(call *ast.FunctionCall) string
}
