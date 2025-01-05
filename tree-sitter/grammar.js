/**
 * @file A programming language built in Go, offering compilation with QBE, direct interpretation, and transpilation to multiple languages
 * @author Jozef Steinhübl <contact@xhyrom.dev>
 * @license EUPL-1.2
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "blom",

  rules: {
    source_file: ($) => repeat($._definition),

    _definition: ($) => choice($.function_definition),

    function_definition: ($) =>
      seq(
        "fun",
        $.identifier,
        $.parameter_list,
        optional(seq("->", $._type)),
        $.block,
      ),

    parameter_list: ($) => seq("(", optional(series_of($.parameter, ",")), ")"),

    parameter: ($) => seq($.identifier, ":", $._type),

    block: ($) => seq("{", repeat($._statement), "}"),

    _statement: ($) => choice($.expression_statement),

    expression_statement: ($) => seq($._expression, ";"),

    _expression: ($) => choice($.identifier, $.literal),

    _type: ($) =>
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
  },
});

function series_of(rule, separator) {
  return seq(rule, repeat(seq(separator, rule)), optional(separator));
}
