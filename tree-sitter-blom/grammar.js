/**
 * @file A programming language built in Go, offering compilation with QBE, direct interpretation, and transpilation to multiple languages
 * @author Jozef Steinhübl <contact@xhyrom.dev>
 * @license EUPL-1.2
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "blom",

  extras: ($) => [/\s/, $.comment],

  conflicts: ($) => [[$.native_function_annotations, $.annotation]],

  rules: {
    /* General rules */
    source_file: ($) => repeat($._module_statement),

    _module_statement: ($) =>
      choice(/*$.import_statement,*/ $.function_definition),

    /* Comments */
    comment: ($) => token(seq("//", /.*/)),

    /* Function definition */
    function_definition: ($) =>
      choice($.native_function_definition, $.regular_function_definition),

    native_function_definition: ($) =>
      seq(
        "fun",
        optional(field("annotations", $.native_function_annotations)),
        field("name", $.identifier),
        field("arguments", $.function_arguments),
        "->",
        field("return_type", $._type),
        ";",
      ),

    regular_function_definition: ($) =>
      seq(
        "fun",
        optional(field("annotations", $.function_annotations)),
        field("name", $.identifier),
        field("arguments", $.function_arguments),
        optional(seq("->", field("return_type", $._type))),
        optional($.block),
      ),

    function_annotations: ($) => series_of(seq($.annotation), " "),

    native_function_annotations: ($) =>
      prec.right(seq("@native", optional($.function_annotations))),

    annotation: ($) =>
      field(
        "annotation",
        choice(seq("@", field("name", $.identifier)), "@native"),
      ),

    function_arguments: ($) =>
      seq(
        "(",
        optional(
          seq(
            series_of($.function_parameter, ","),
            optional(seq(",", $.variadic_parameter)),
          ),
        ),
        ")",
      ),

    function_parameter: ($) => seq($.identifier, ":", $._type),
    variadic_parameter: ($) => "...",

    /* Statements */
    _statement: ($) =>
      choice(
        $.expression_statement,
        $.if_statement,
        $.return_statement,
        $.variable_declaration,
        $.for_statement,
        $.while_statement,
        $.block,
      ),

    expression_statement: ($) => seq($._expression, ";"),

    if_statement: ($) =>
      seq(
        "if",
        field("condition", $._expression),
        field("then", $.block),
        optional(seq("else", field("else", $.block))),
      ),

    return_statement: ($) => seq("return", optional($._expression), ";"),

    variable_declaration: ($) =>
      seq(
        field("type", $._type),
        field("name", $.identifier),
        "=",
        field("value", $._expression),
        ";",
      ),

    for_statement: ($) =>
      choice(
        seq(
          "for",
          field("initializer", $.variable_declaration),
          field("condition", $._expression),
          ";",
          field("step", optional(seq($._expression, ";"))),
          field("body", $.block),
        ),
        seq(
          "for",
          field("condition", $._expression),
          ";",
          field("step", optional(seq($._expression, ";"))),
          field("body", $.block),
        ),
      ),

    while_statement: ($) =>
      seq("while", field("condition", $._expression), field("body", $.block)),

    block: ($) => seq("{", field("body", repeat($._statement)), "}"),

    _expression: ($) =>
      choice(
        $.assignment_expression,
        $.identifier,
        $.literal,
        $.function_call,
        $.binary_expression,
        $.unary_expression,
        $.parenthesized_expression,
      ),

    assignment_expression: ($) =>
      prec.right(
        seq(field("left", $.identifier), "=", field("right", $._expression)),
      ),

    binary_expression: ($) =>
      choice(
        binary_expression(10, $._expression, "*", $._expression),
        binary_expression(10, $._expression, "/", $._expression),
        binary_expression(10, $._expression, "%", $._expression),
        binary_expression(9, $._expression, "+", $._expression),
        binary_expression(9, $._expression, "-", $._expression),
        binary_expression(8, $._expression, "<<", $._expression),
        binary_expression(8, $._expression, ">>", $._expression),
        binary_expression(7, $._expression, "<", $._expression),
        binary_expression(7, $._expression, "<=", $._expression),
        binary_expression(7, $._expression, ">", $._expression),
        binary_expression(7, $._expression, ">=", $._expression),
        binary_expression(6, $._expression, "==", $._expression),
        binary_expression(5, $._expression, "&", $._expression),
        binary_expression(4, $._expression, "^", $._expression),
        binary_expression(3, $._expression, "|", $._expression),
        binary_expression(2, $._expression, "&&", $._expression),
        binary_expression(1, $._expression, "||", $._expression),
      ),

    unary_expression: ($) =>
      choice(
        unary_expression("+", $._expression),
        unary_expression("-", $._expression),
        unary_expression("~", $._expression),
      ),

    parenthesized_expression: ($) => seq("(", $._expression, ")"),

    function_call: ($) =>
      seq(
        field("function", $.function_name_identifier),
        field("parameters", $.function_call_parameters),
      ),

    function_call_parameters: ($) =>
      seq("(", optional(series_of($._expression, ",")), ")"),

    /* Types */
    _type: ($) => choice($.primitive_type_identifier),

    primitive_type_identifier: ($) =>
      choice(
        "i8",
        "u8",
        "i16",
        "u16",
        "i32",
        "u32",
        "i64",
        "u64",
        "f32",
        "f64",
        "bool",
        "string",
        "char",
      ),

    literal: ($) =>
      choice(
        $.integer_literal,
        $.float_literal,
        $.string_literal,
        $.char_literal,
        $.boolean_literal,
      ),

    integer_literal: ($) => token(/[0-9]+/),
    float_literal: ($) => token(/[0-9]+\.[0-9]+/),
    string_literal: ($) => token(/".*?"/),
    char_literal: ($) => token(/'.*?'/),
    boolean_literal: ($) => choice("true", "false"),

    identifier: ($) => /[a-zA-Z_]\w*/,
    function_name_identifier: ($) => choice(/[a-zA-Z_]\w*/, /@[a-zA-Z_]\w*/),
  },
});

function series_of(rule, separator) {
  return seq(rule, repeat(seq(separator, rule)), optional(separator));
}

function binary_expression(precedence, left, operator, right) {
  return prec.left(
    precedence,
    seq(
      field("left", left),
      field("operator", operator),
      field("right", right),
    ),
  );
}

function unary_expression(operator, operand) {
  return prec.left(seq(field("operator", operator), field("operand", operand)));
}
