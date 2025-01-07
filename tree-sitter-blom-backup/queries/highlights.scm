; General
((source_file) @source)

; Comments
((comment) @comment)

; Keywords
[
  "fun"
  "return"
  "if"
  "else"
  "for"
  "while"
] @keyword

; Function definitions
(native_function_definition
  name: (identifier) @name)
(regular_function_definition
  name: (identifier) @name)

(annotation) @attribute

; Statements
(expression_statement) @statement
(if_statement) @conditional

((variable_declaration
  (primitive_type_identifier) @type
  (identifier) @variable))
(for_statement) @loop
(while_statement) @loop
(block) @block

; Expressions
(binary_expression
  operator: _ @operator)
(unary_expression
  operator: _ @operator)
((function_call
  function: (function_name_identifier) @function.call
  parameters: (function_call_parameters) @parameter))

; Types
(primitive_type_identifier) @type

; Literals
(integer_literal) @number
(float_literal) @number
(string_literal) @string
(char_literal) @character
(boolean_literal) @boolean
