module github.com/xhyrom/blom

go 1.23.6

replace blom/analyzer => ./toolchain/analyzer

replace blom/ast => ./toolchain/ast

replace blom/compiler => ./toolchain/compiler

replace blom/debug => ./toolchain/debug

replace blom/interpreter => ./toolchain/interpreter

replace blom/lexer => ./toolchain/lexer

replace blom/parser => ./toolchain/parser

replace blom/scope => ./toolchain/scope

replace blom/tokens => ./toolchain/tokens

replace blom/qbe => ./qbe
