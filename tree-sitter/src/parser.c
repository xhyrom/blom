#include "tree_sitter/parser.h"

#if defined(__GNUC__) || defined(__clang__)
#pragma GCC diagnostic ignored "-Wmissing-field-initializers"
#endif

#define LANGUAGE_VERSION 14
#define STATE_COUNT 36
#define LARGE_STATE_COUNT 2
#define SYMBOL_COUNT 45
#define ALIAS_COUNT 0
#define TOKEN_COUNT 30
#define EXTERNAL_TOKEN_COUNT 0
#define FIELD_COUNT 0
#define MAX_ALIAS_SEQUENCE_LENGTH 6
#define PRODUCTION_ID_COUNT 1

enum ts_symbol_identifiers {
  anon_sym_fun = 1,
  anon_sym_DASH_GT = 2,
  anon_sym_LPAREN = 3,
  anon_sym_COMMA = 4,
  anon_sym_RPAREN = 5,
  anon_sym_COLON = 6,
  anon_sym_LBRACE = 7,
  anon_sym_RBRACE = 8,
  anon_sym_SEMI = 9,
  anon_sym_i8 = 10,
  anon_sym_u8 = 11,
  anon_sym_i16 = 12,
  anon_sym_u16 = 13,
  anon_sym_i32 = 14,
  anon_sym_u32 = 15,
  anon_sym_i64 = 16,
  anon_sym_u64 = 17,
  anon_sym_f32 = 18,
  anon_sym_f64 = 19,
  anon_sym_bool = 20,
  anon_sym_string = 21,
  anon_sym_char = 22,
  sym_integer_literal = 23,
  sym_float_literal = 24,
  sym_string_literal = 25,
  sym_char_literal = 26,
  anon_sym_true = 27,
  anon_sym_false = 28,
  sym_identifier = 29,
  sym_source_file = 30,
  sym__definition = 31,
  sym_function_definition = 32,
  sym_parameter_list = 33,
  sym_parameter = 34,
  sym_block = 35,
  sym__statement = 36,
  sym_expression_statement = 37,
  sym__expression = 38,
  sym__type = 39,
  sym_literal = 40,
  sym_boolean_literal = 41,
  aux_sym_source_file_repeat1 = 42,
  aux_sym_parameter_list_repeat1 = 43,
  aux_sym_block_repeat1 = 44,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [anon_sym_fun] = "fun",
  [anon_sym_DASH_GT] = "->",
  [anon_sym_LPAREN] = "(",
  [anon_sym_COMMA] = ",",
  [anon_sym_RPAREN] = ")",
  [anon_sym_COLON] = ":",
  [anon_sym_LBRACE] = "{",
  [anon_sym_RBRACE] = "}",
  [anon_sym_SEMI] = ";",
  [anon_sym_i8] = "i8",
  [anon_sym_u8] = "u8",
  [anon_sym_i16] = "i16",
  [anon_sym_u16] = "u16",
  [anon_sym_i32] = "i32",
  [anon_sym_u32] = "u32",
  [anon_sym_i64] = "i64",
  [anon_sym_u64] = "u64",
  [anon_sym_f32] = "f32",
  [anon_sym_f64] = "f64",
  [anon_sym_bool] = "bool",
  [anon_sym_string] = "string",
  [anon_sym_char] = "char",
  [sym_integer_literal] = "integer_literal",
  [sym_float_literal] = "float_literal",
  [sym_string_literal] = "string_literal",
  [sym_char_literal] = "char_literal",
  [anon_sym_true] = "true",
  [anon_sym_false] = "false",
  [sym_identifier] = "identifier",
  [sym_source_file] = "source_file",
  [sym__definition] = "_definition",
  [sym_function_definition] = "function_definition",
  [sym_parameter_list] = "parameter_list",
  [sym_parameter] = "parameter",
  [sym_block] = "block",
  [sym__statement] = "_statement",
  [sym_expression_statement] = "expression_statement",
  [sym__expression] = "_expression",
  [sym__type] = "_type",
  [sym_literal] = "literal",
  [sym_boolean_literal] = "boolean_literal",
  [aux_sym_source_file_repeat1] = "source_file_repeat1",
  [aux_sym_parameter_list_repeat1] = "parameter_list_repeat1",
  [aux_sym_block_repeat1] = "block_repeat1",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [anon_sym_fun] = anon_sym_fun,
  [anon_sym_DASH_GT] = anon_sym_DASH_GT,
  [anon_sym_LPAREN] = anon_sym_LPAREN,
  [anon_sym_COMMA] = anon_sym_COMMA,
  [anon_sym_RPAREN] = anon_sym_RPAREN,
  [anon_sym_COLON] = anon_sym_COLON,
  [anon_sym_LBRACE] = anon_sym_LBRACE,
  [anon_sym_RBRACE] = anon_sym_RBRACE,
  [anon_sym_SEMI] = anon_sym_SEMI,
  [anon_sym_i8] = anon_sym_i8,
  [anon_sym_u8] = anon_sym_u8,
  [anon_sym_i16] = anon_sym_i16,
  [anon_sym_u16] = anon_sym_u16,
  [anon_sym_i32] = anon_sym_i32,
  [anon_sym_u32] = anon_sym_u32,
  [anon_sym_i64] = anon_sym_i64,
  [anon_sym_u64] = anon_sym_u64,
  [anon_sym_f32] = anon_sym_f32,
  [anon_sym_f64] = anon_sym_f64,
  [anon_sym_bool] = anon_sym_bool,
  [anon_sym_string] = anon_sym_string,
  [anon_sym_char] = anon_sym_char,
  [sym_integer_literal] = sym_integer_literal,
  [sym_float_literal] = sym_float_literal,
  [sym_string_literal] = sym_string_literal,
  [sym_char_literal] = sym_char_literal,
  [anon_sym_true] = anon_sym_true,
  [anon_sym_false] = anon_sym_false,
  [sym_identifier] = sym_identifier,
  [sym_source_file] = sym_source_file,
  [sym__definition] = sym__definition,
  [sym_function_definition] = sym_function_definition,
  [sym_parameter_list] = sym_parameter_list,
  [sym_parameter] = sym_parameter,
  [sym_block] = sym_block,
  [sym__statement] = sym__statement,
  [sym_expression_statement] = sym_expression_statement,
  [sym__expression] = sym__expression,
  [sym__type] = sym__type,
  [sym_literal] = sym_literal,
  [sym_boolean_literal] = sym_boolean_literal,
  [aux_sym_source_file_repeat1] = aux_sym_source_file_repeat1,
  [aux_sym_parameter_list_repeat1] = aux_sym_parameter_list_repeat1,
  [aux_sym_block_repeat1] = aux_sym_block_repeat1,
};

static const TSSymbolMetadata ts_symbol_metadata[] = {
  [ts_builtin_sym_end] = {
    .visible = false,
    .named = true,
  },
  [anon_sym_fun] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_DASH_GT] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COMMA] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RPAREN] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_COLON] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_LBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_RBRACE] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_SEMI] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_i8] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_u8] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_i16] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_u16] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_i32] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_u32] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_i64] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_u64] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_f32] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_f64] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_bool] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_string] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_char] = {
    .visible = true,
    .named = false,
  },
  [sym_integer_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_float_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_string_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_char_literal] = {
    .visible = true,
    .named = true,
  },
  [anon_sym_true] = {
    .visible = true,
    .named = false,
  },
  [anon_sym_false] = {
    .visible = true,
    .named = false,
  },
  [sym_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym_source_file] = {
    .visible = true,
    .named = true,
  },
  [sym__definition] = {
    .visible = false,
    .named = true,
  },
  [sym_function_definition] = {
    .visible = true,
    .named = true,
  },
  [sym_parameter_list] = {
    .visible = true,
    .named = true,
  },
  [sym_parameter] = {
    .visible = true,
    .named = true,
  },
  [sym_block] = {
    .visible = true,
    .named = true,
  },
  [sym__statement] = {
    .visible = false,
    .named = true,
  },
  [sym_expression_statement] = {
    .visible = true,
    .named = true,
  },
  [sym__expression] = {
    .visible = false,
    .named = true,
  },
  [sym__type] = {
    .visible = false,
    .named = true,
  },
  [sym_literal] = {
    .visible = true,
    .named = true,
  },
  [sym_boolean_literal] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_source_file_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_parameter_list_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_block_repeat1] = {
    .visible = false,
    .named = false,
  },
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
};

static const uint16_t ts_non_terminal_alias_map[] = {
  0,
};

static const TSStateId ts_primary_state_ids[STATE_COUNT] = {
  [0] = 0,
  [1] = 1,
  [2] = 2,
  [3] = 3,
  [4] = 4,
  [5] = 5,
  [6] = 6,
  [7] = 7,
  [8] = 8,
  [9] = 9,
  [10] = 10,
  [11] = 11,
  [12] = 12,
  [13] = 13,
  [14] = 14,
  [15] = 15,
  [16] = 16,
  [17] = 17,
  [18] = 18,
  [19] = 19,
  [20] = 20,
  [21] = 21,
  [22] = 22,
  [23] = 23,
  [24] = 24,
  [25] = 25,
  [26] = 26,
  [27] = 27,
  [28] = 28,
  [29] = 29,
  [30] = 30,
  [31] = 31,
  [32] = 32,
  [33] = 33,
  [34] = 34,
  [35] = 35,
};

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(36);
      ADVANCE_MAP(
        '"', 2,
        '\'', 3,
        '(', 39,
        ')', 41,
        ',', 40,
        '-', 16,
        ':', 42,
        ';', 45,
        'b', 27,
        'c', 21,
        'f', 10,
        'i', 5,
        's', 33,
        't', 29,
        'u', 6,
        '{', 43,
        '}', 44,
      );
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(0);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(59);
      END_STATE();
    case 1:
      if (lookahead == '"') ADVANCE(2);
      if (lookahead == '\'') ADVANCE(3);
      if (lookahead == 'f') ADVANCE(67);
      if (lookahead == 't') ADVANCE(71);
      if (lookahead == '}') ADVANCE(44);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(1);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(59);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 2:
      if (lookahead == '"') ADVANCE(61);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(2);
      END_STATE();
    case 3:
      if (lookahead == '\'') ADVANCE(62);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(3);
      END_STATE();
    case 4:
      if (lookahead == ')') ADVANCE(41);
      if (('\t' <= lookahead && lookahead <= '\r') ||
          lookahead == ' ') SKIP(4);
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 5:
      if (lookahead == '1') ADVANCE(14);
      if (lookahead == '3') ADVANCE(8);
      if (lookahead == '6') ADVANCE(12);
      if (lookahead == '8') ADVANCE(46);
      END_STATE();
    case 6:
      if (lookahead == '1') ADVANCE(15);
      if (lookahead == '3') ADVANCE(9);
      if (lookahead == '6') ADVANCE(13);
      if (lookahead == '8') ADVANCE(47);
      END_STATE();
    case 7:
      if (lookahead == '2') ADVANCE(54);
      END_STATE();
    case 8:
      if (lookahead == '2') ADVANCE(50);
      END_STATE();
    case 9:
      if (lookahead == '2') ADVANCE(51);
      END_STATE();
    case 10:
      if (lookahead == '3') ADVANCE(7);
      if (lookahead == '6') ADVANCE(11);
      if (lookahead == 'a') ADVANCE(23);
      if (lookahead == 'u') ADVANCE(25);
      END_STATE();
    case 11:
      if (lookahead == '4') ADVANCE(55);
      END_STATE();
    case 12:
      if (lookahead == '4') ADVANCE(52);
      END_STATE();
    case 13:
      if (lookahead == '4') ADVANCE(53);
      END_STATE();
    case 14:
      if (lookahead == '6') ADVANCE(48);
      END_STATE();
    case 15:
      if (lookahead == '6') ADVANCE(49);
      END_STATE();
    case 16:
      if (lookahead == '>') ADVANCE(38);
      END_STATE();
    case 17:
      if (lookahead == 'a') ADVANCE(31);
      END_STATE();
    case 18:
      if (lookahead == 'e') ADVANCE(63);
      END_STATE();
    case 19:
      if (lookahead == 'e') ADVANCE(65);
      END_STATE();
    case 20:
      if (lookahead == 'g') ADVANCE(57);
      END_STATE();
    case 21:
      if (lookahead == 'h') ADVANCE(17);
      END_STATE();
    case 22:
      if (lookahead == 'i') ADVANCE(26);
      END_STATE();
    case 23:
      if (lookahead == 'l') ADVANCE(32);
      END_STATE();
    case 24:
      if (lookahead == 'l') ADVANCE(56);
      END_STATE();
    case 25:
      if (lookahead == 'n') ADVANCE(37);
      END_STATE();
    case 26:
      if (lookahead == 'n') ADVANCE(20);
      END_STATE();
    case 27:
      if (lookahead == 'o') ADVANCE(28);
      END_STATE();
    case 28:
      if (lookahead == 'o') ADVANCE(24);
      END_STATE();
    case 29:
      if (lookahead == 'r') ADVANCE(34);
      END_STATE();
    case 30:
      if (lookahead == 'r') ADVANCE(22);
      END_STATE();
    case 31:
      if (lookahead == 'r') ADVANCE(58);
      END_STATE();
    case 32:
      if (lookahead == 's') ADVANCE(19);
      END_STATE();
    case 33:
      if (lookahead == 't') ADVANCE(30);
      END_STATE();
    case 34:
      if (lookahead == 'u') ADVANCE(18);
      END_STATE();
    case 35:
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(60);
      END_STATE();
    case 36:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 37:
      ACCEPT_TOKEN(anon_sym_fun);
      END_STATE();
    case 38:
      ACCEPT_TOKEN(anon_sym_DASH_GT);
      END_STATE();
    case 39:
      ACCEPT_TOKEN(anon_sym_LPAREN);
      END_STATE();
    case 40:
      ACCEPT_TOKEN(anon_sym_COMMA);
      END_STATE();
    case 41:
      ACCEPT_TOKEN(anon_sym_RPAREN);
      END_STATE();
    case 42:
      ACCEPT_TOKEN(anon_sym_COLON);
      END_STATE();
    case 43:
      ACCEPT_TOKEN(anon_sym_LBRACE);
      END_STATE();
    case 44:
      ACCEPT_TOKEN(anon_sym_RBRACE);
      END_STATE();
    case 45:
      ACCEPT_TOKEN(anon_sym_SEMI);
      END_STATE();
    case 46:
      ACCEPT_TOKEN(anon_sym_i8);
      END_STATE();
    case 47:
      ACCEPT_TOKEN(anon_sym_u8);
      END_STATE();
    case 48:
      ACCEPT_TOKEN(anon_sym_i16);
      END_STATE();
    case 49:
      ACCEPT_TOKEN(anon_sym_u16);
      END_STATE();
    case 50:
      ACCEPT_TOKEN(anon_sym_i32);
      END_STATE();
    case 51:
      ACCEPT_TOKEN(anon_sym_u32);
      END_STATE();
    case 52:
      ACCEPT_TOKEN(anon_sym_i64);
      END_STATE();
    case 53:
      ACCEPT_TOKEN(anon_sym_u64);
      END_STATE();
    case 54:
      ACCEPT_TOKEN(anon_sym_f32);
      END_STATE();
    case 55:
      ACCEPT_TOKEN(anon_sym_f64);
      END_STATE();
    case 56:
      ACCEPT_TOKEN(anon_sym_bool);
      END_STATE();
    case 57:
      ACCEPT_TOKEN(anon_sym_string);
      END_STATE();
    case 58:
      ACCEPT_TOKEN(anon_sym_char);
      END_STATE();
    case 59:
      ACCEPT_TOKEN(sym_integer_literal);
      if (lookahead == '.') ADVANCE(35);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(59);
      END_STATE();
    case 60:
      ACCEPT_TOKEN(sym_float_literal);
      if (('0' <= lookahead && lookahead <= '9')) ADVANCE(60);
      END_STATE();
    case 61:
      ACCEPT_TOKEN(sym_string_literal);
      if (lookahead == '"') ADVANCE(61);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(2);
      END_STATE();
    case 62:
      ACCEPT_TOKEN(sym_char_literal);
      if (lookahead == '\'') ADVANCE(62);
      if (lookahead != 0 &&
          lookahead != '\n') ADVANCE(3);
      END_STATE();
    case 63:
      ACCEPT_TOKEN(anon_sym_true);
      END_STATE();
    case 64:
      ACCEPT_TOKEN(anon_sym_true);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 65:
      ACCEPT_TOKEN(anon_sym_false);
      END_STATE();
    case 66:
      ACCEPT_TOKEN(anon_sym_false);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 67:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'a') ADVANCE(70);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('b' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 68:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'e') ADVANCE(64);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 69:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'e') ADVANCE(66);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 70:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'l') ADVANCE(72);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 71:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'r') ADVANCE(73);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 72:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 's') ADVANCE(69);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 73:
      ACCEPT_TOKEN(sym_identifier);
      if (lookahead == 'u') ADVANCE(68);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    case 74:
      ACCEPT_TOKEN(sym_identifier);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(74);
      END_STATE();
    default:
      return false;
  }
}

static const TSLexMode ts_lex_modes[STATE_COUNT] = {
  [0] = {.lex_state = 0},
  [1] = {.lex_state = 0},
  [2] = {.lex_state = 0},
  [3] = {.lex_state = 0},
  [4] = {.lex_state = 1},
  [5] = {.lex_state = 1},
  [6] = {.lex_state = 1},
  [7] = {.lex_state = 1},
  [8] = {.lex_state = 0},
  [9] = {.lex_state = 0},
  [10] = {.lex_state = 4},
  [11] = {.lex_state = 0},
  [12] = {.lex_state = 0},
  [13] = {.lex_state = 4},
  [14] = {.lex_state = 0},
  [15] = {.lex_state = 4},
  [16] = {.lex_state = 0},
  [17] = {.lex_state = 0},
  [18] = {.lex_state = 0},
  [19] = {.lex_state = 0},
  [20] = {.lex_state = 0},
  [21] = {.lex_state = 0},
  [22] = {.lex_state = 0},
  [23] = {.lex_state = 0},
  [24] = {.lex_state = 4},
  [25] = {.lex_state = 0},
  [26] = {.lex_state = 0},
  [27] = {.lex_state = 0},
  [28] = {.lex_state = 0},
  [29] = {.lex_state = 0},
  [30] = {.lex_state = 4},
  [31] = {.lex_state = 0},
  [32] = {.lex_state = 0},
  [33] = {.lex_state = 0},
  [34] = {.lex_state = 0},
  [35] = {.lex_state = 0},
};

static const uint16_t ts_parse_table[LARGE_STATE_COUNT][SYMBOL_COUNT] = {
  [0] = {
    [ts_builtin_sym_end] = ACTIONS(1),
    [anon_sym_fun] = ACTIONS(1),
    [anon_sym_DASH_GT] = ACTIONS(1),
    [anon_sym_LPAREN] = ACTIONS(1),
    [anon_sym_COMMA] = ACTIONS(1),
    [anon_sym_RPAREN] = ACTIONS(1),
    [anon_sym_COLON] = ACTIONS(1),
    [anon_sym_LBRACE] = ACTIONS(1),
    [anon_sym_RBRACE] = ACTIONS(1),
    [anon_sym_SEMI] = ACTIONS(1),
    [anon_sym_i8] = ACTIONS(1),
    [anon_sym_u8] = ACTIONS(1),
    [anon_sym_i16] = ACTIONS(1),
    [anon_sym_u16] = ACTIONS(1),
    [anon_sym_i32] = ACTIONS(1),
    [anon_sym_u32] = ACTIONS(1),
    [anon_sym_i64] = ACTIONS(1),
    [anon_sym_u64] = ACTIONS(1),
    [anon_sym_f32] = ACTIONS(1),
    [anon_sym_f64] = ACTIONS(1),
    [anon_sym_bool] = ACTIONS(1),
    [anon_sym_string] = ACTIONS(1),
    [anon_sym_char] = ACTIONS(1),
    [sym_integer_literal] = ACTIONS(1),
    [sym_float_literal] = ACTIONS(1),
    [sym_string_literal] = ACTIONS(1),
    [sym_char_literal] = ACTIONS(1),
    [anon_sym_true] = ACTIONS(1),
    [anon_sym_false] = ACTIONS(1),
  },
  [1] = {
    [sym_source_file] = STATE(31),
    [sym__definition] = STATE(9),
    [sym_function_definition] = STATE(9),
    [aux_sym_source_file_repeat1] = STATE(9),
    [ts_builtin_sym_end] = ACTIONS(3),
    [anon_sym_fun] = ACTIONS(5),
  },
};

static const uint16_t ts_small_parse_table[] = {
  [0] = 2,
    STATE(22), 1,
      sym__type,
    ACTIONS(7), 13,
      anon_sym_i8,
      anon_sym_u8,
      anon_sym_i16,
      anon_sym_u16,
      anon_sym_i32,
      anon_sym_u32,
      anon_sym_i64,
      anon_sym_u64,
      anon_sym_f32,
      anon_sym_f64,
      anon_sym_bool,
      anon_sym_string,
      anon_sym_char,
  [19] = 2,
    STATE(18), 1,
      sym__type,
    ACTIONS(9), 13,
      anon_sym_i8,
      anon_sym_u8,
      anon_sym_i16,
      anon_sym_u16,
      anon_sym_i32,
      anon_sym_u32,
      anon_sym_i64,
      anon_sym_u64,
      anon_sym_f32,
      anon_sym_f64,
      anon_sym_bool,
      anon_sym_string,
      anon_sym_char,
  [38] = 8,
    ACTIONS(11), 1,
      anon_sym_RBRACE,
    ACTIONS(13), 1,
      sym_integer_literal,
    ACTIONS(19), 1,
      sym_identifier,
    STATE(32), 1,
      sym_boolean_literal,
    ACTIONS(17), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(35), 2,
      sym__expression,
      sym_literal,
    ACTIONS(15), 3,
      sym_float_literal,
      sym_string_literal,
      sym_char_literal,
    STATE(5), 3,
      sym__statement,
      sym_expression_statement,
      aux_sym_block_repeat1,
  [69] = 8,
    ACTIONS(13), 1,
      sym_integer_literal,
    ACTIONS(19), 1,
      sym_identifier,
    ACTIONS(21), 1,
      anon_sym_RBRACE,
    STATE(32), 1,
      sym_boolean_literal,
    ACTIONS(17), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(35), 2,
      sym__expression,
      sym_literal,
    ACTIONS(15), 3,
      sym_float_literal,
      sym_string_literal,
      sym_char_literal,
    STATE(6), 3,
      sym__statement,
      sym_expression_statement,
      aux_sym_block_repeat1,
  [100] = 8,
    ACTIONS(23), 1,
      anon_sym_RBRACE,
    ACTIONS(25), 1,
      sym_integer_literal,
    ACTIONS(34), 1,
      sym_identifier,
    STATE(32), 1,
      sym_boolean_literal,
    ACTIONS(31), 2,
      anon_sym_true,
      anon_sym_false,
    STATE(35), 2,
      sym__expression,
      sym_literal,
    ACTIONS(28), 3,
      sym_float_literal,
      sym_string_literal,
      sym_char_literal,
    STATE(6), 3,
      sym__statement,
      sym_expression_statement,
      aux_sym_block_repeat1,
  [131] = 2,
    ACTIONS(37), 4,
      anon_sym_RBRACE,
      sym_float_literal,
      sym_string_literal,
      sym_char_literal,
    ACTIONS(39), 4,
      sym_integer_literal,
      anon_sym_true,
      anon_sym_false,
      sym_identifier,
  [144] = 3,
    ACTIONS(41), 1,
      ts_builtin_sym_end,
    ACTIONS(43), 1,
      anon_sym_fun,
    STATE(8), 3,
      sym__definition,
      sym_function_definition,
      aux_sym_source_file_repeat1,
  [156] = 3,
    ACTIONS(5), 1,
      anon_sym_fun,
    ACTIONS(46), 1,
      ts_builtin_sym_end,
    STATE(8), 3,
      sym__definition,
      sym_function_definition,
      aux_sym_source_file_repeat1,
  [168] = 3,
    ACTIONS(48), 1,
      anon_sym_RPAREN,
    ACTIONS(50), 1,
      sym_identifier,
    STATE(23), 1,
      sym_parameter,
  [178] = 3,
    ACTIONS(52), 1,
      anon_sym_COMMA,
    ACTIONS(54), 1,
      anon_sym_RPAREN,
    STATE(12), 1,
      aux_sym_parameter_list_repeat1,
  [188] = 3,
    ACTIONS(56), 1,
      anon_sym_COMMA,
    ACTIONS(58), 1,
      anon_sym_RPAREN,
    STATE(16), 1,
      aux_sym_parameter_list_repeat1,
  [198] = 3,
    ACTIONS(50), 1,
      sym_identifier,
    ACTIONS(58), 1,
      anon_sym_RPAREN,
    STATE(23), 1,
      sym_parameter,
  [208] = 3,
    ACTIONS(60), 1,
      anon_sym_DASH_GT,
    ACTIONS(62), 1,
      anon_sym_LBRACE,
    STATE(17), 1,
      sym_block,
  [218] = 3,
    ACTIONS(50), 1,
      sym_identifier,
    ACTIONS(64), 1,
      anon_sym_RPAREN,
    STATE(11), 1,
      sym_parameter,
  [228] = 3,
    ACTIONS(66), 1,
      anon_sym_COMMA,
    ACTIONS(69), 1,
      anon_sym_RPAREN,
    STATE(16), 1,
      aux_sym_parameter_list_repeat1,
  [238] = 1,
    ACTIONS(71), 2,
      ts_builtin_sym_end,
      anon_sym_fun,
  [243] = 2,
    ACTIONS(62), 1,
      anon_sym_LBRACE,
    STATE(26), 1,
      sym_block,
  [250] = 2,
    ACTIONS(73), 1,
      anon_sym_LPAREN,
    STATE(14), 1,
      sym_parameter_list,
  [257] = 1,
    ACTIONS(75), 2,
      anon_sym_DASH_GT,
      anon_sym_LBRACE,
  [262] = 1,
    ACTIONS(77), 2,
      ts_builtin_sym_end,
      anon_sym_fun,
  [267] = 1,
    ACTIONS(79), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [272] = 1,
    ACTIONS(69), 2,
      anon_sym_COMMA,
      anon_sym_RPAREN,
  [277] = 2,
    ACTIONS(50), 1,
      sym_identifier,
    STATE(23), 1,
      sym_parameter,
  [284] = 1,
    ACTIONS(81), 2,
      anon_sym_DASH_GT,
      anon_sym_LBRACE,
  [289] = 1,
    ACTIONS(83), 2,
      ts_builtin_sym_end,
      anon_sym_fun,
  [294] = 1,
    ACTIONS(85), 2,
      anon_sym_DASH_GT,
      anon_sym_LBRACE,
  [299] = 1,
    ACTIONS(87), 2,
      ts_builtin_sym_end,
      anon_sym_fun,
  [304] = 1,
    ACTIONS(89), 2,
      anon_sym_DASH_GT,
      anon_sym_LBRACE,
  [309] = 1,
    ACTIONS(91), 1,
      sym_identifier,
  [313] = 1,
    ACTIONS(93), 1,
      ts_builtin_sym_end,
  [317] = 1,
    ACTIONS(95), 1,
      anon_sym_SEMI,
  [321] = 1,
    ACTIONS(97), 1,
      anon_sym_SEMI,
  [325] = 1,
    ACTIONS(99), 1,
      anon_sym_COLON,
  [329] = 1,
    ACTIONS(101), 1,
      anon_sym_SEMI,
};

static const uint32_t ts_small_parse_table_map[] = {
  [SMALL_STATE(2)] = 0,
  [SMALL_STATE(3)] = 19,
  [SMALL_STATE(4)] = 38,
  [SMALL_STATE(5)] = 69,
  [SMALL_STATE(6)] = 100,
  [SMALL_STATE(7)] = 131,
  [SMALL_STATE(8)] = 144,
  [SMALL_STATE(9)] = 156,
  [SMALL_STATE(10)] = 168,
  [SMALL_STATE(11)] = 178,
  [SMALL_STATE(12)] = 188,
  [SMALL_STATE(13)] = 198,
  [SMALL_STATE(14)] = 208,
  [SMALL_STATE(15)] = 218,
  [SMALL_STATE(16)] = 228,
  [SMALL_STATE(17)] = 238,
  [SMALL_STATE(18)] = 243,
  [SMALL_STATE(19)] = 250,
  [SMALL_STATE(20)] = 257,
  [SMALL_STATE(21)] = 262,
  [SMALL_STATE(22)] = 267,
  [SMALL_STATE(23)] = 272,
  [SMALL_STATE(24)] = 277,
  [SMALL_STATE(25)] = 284,
  [SMALL_STATE(26)] = 289,
  [SMALL_STATE(27)] = 294,
  [SMALL_STATE(28)] = 299,
  [SMALL_STATE(29)] = 304,
  [SMALL_STATE(30)] = 309,
  [SMALL_STATE(31)] = 313,
  [SMALL_STATE(32)] = 317,
  [SMALL_STATE(33)] = 321,
  [SMALL_STATE(34)] = 325,
  [SMALL_STATE(35)] = 329,
};

static const TSParseActionEntry ts_parse_actions[] = {
  [0] = {.entry = {.count = 0, .reusable = false}},
  [1] = {.entry = {.count = 1, .reusable = false}}, RECOVER(),
  [3] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_source_file, 0, 0, 0),
  [5] = {.entry = {.count = 1, .reusable = true}}, SHIFT(30),
  [7] = {.entry = {.count = 1, .reusable = true}}, SHIFT(22),
  [9] = {.entry = {.count = 1, .reusable = true}}, SHIFT(18),
  [11] = {.entry = {.count = 1, .reusable = true}}, SHIFT(21),
  [13] = {.entry = {.count = 1, .reusable = false}}, SHIFT(32),
  [15] = {.entry = {.count = 1, .reusable = true}}, SHIFT(32),
  [17] = {.entry = {.count = 1, .reusable = false}}, SHIFT(33),
  [19] = {.entry = {.count = 1, .reusable = false}}, SHIFT(35),
  [21] = {.entry = {.count = 1, .reusable = true}}, SHIFT(28),
  [23] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_block_repeat1, 2, 0, 0),
  [25] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_block_repeat1, 2, 0, 0), SHIFT_REPEAT(32),
  [28] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_block_repeat1, 2, 0, 0), SHIFT_REPEAT(32),
  [31] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_block_repeat1, 2, 0, 0), SHIFT_REPEAT(33),
  [34] = {.entry = {.count = 2, .reusable = false}}, REDUCE(aux_sym_block_repeat1, 2, 0, 0), SHIFT_REPEAT(35),
  [37] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_expression_statement, 2, 0, 0),
  [39] = {.entry = {.count = 1, .reusable = false}}, REDUCE(sym_expression_statement, 2, 0, 0),
  [41] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_source_file_repeat1, 2, 0, 0),
  [43] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_source_file_repeat1, 2, 0, 0), SHIFT_REPEAT(30),
  [46] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_source_file, 1, 0, 0),
  [48] = {.entry = {.count = 1, .reusable = true}}, SHIFT(29),
  [50] = {.entry = {.count = 1, .reusable = true}}, SHIFT(34),
  [52] = {.entry = {.count = 1, .reusable = true}}, SHIFT(13),
  [54] = {.entry = {.count = 1, .reusable = true}}, SHIFT(20),
  [56] = {.entry = {.count = 1, .reusable = true}}, SHIFT(10),
  [58] = {.entry = {.count = 1, .reusable = true}}, SHIFT(25),
  [60] = {.entry = {.count = 1, .reusable = true}}, SHIFT(3),
  [62] = {.entry = {.count = 1, .reusable = true}}, SHIFT(4),
  [64] = {.entry = {.count = 1, .reusable = true}}, SHIFT(27),
  [66] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_parameter_list_repeat1, 2, 0, 0), SHIFT_REPEAT(24),
  [69] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_parameter_list_repeat1, 2, 0, 0),
  [71] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_function_definition, 4, 0, 0),
  [73] = {.entry = {.count = 1, .reusable = true}}, SHIFT(15),
  [75] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_parameter_list, 3, 0, 0),
  [77] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_block, 2, 0, 0),
  [79] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_parameter, 3, 0, 0),
  [81] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_parameter_list, 4, 0, 0),
  [83] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_function_definition, 6, 0, 0),
  [85] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_parameter_list, 2, 0, 0),
  [87] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_block, 3, 0, 0),
  [89] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_parameter_list, 5, 0, 0),
  [91] = {.entry = {.count = 1, .reusable = true}}, SHIFT(19),
  [93] = {.entry = {.count = 1, .reusable = true}},  ACCEPT_INPUT(),
  [95] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_literal, 1, 0, 0),
  [97] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_boolean_literal, 1, 0, 0),
  [99] = {.entry = {.count = 1, .reusable = true}}, SHIFT(2),
  [101] = {.entry = {.count = 1, .reusable = true}}, SHIFT(7),
};

#ifdef __cplusplus
extern "C" {
#endif
#ifdef TREE_SITTER_HIDE_SYMBOLS
#define TS_PUBLIC
#elif defined(_WIN32)
#define TS_PUBLIC __declspec(dllexport)
#else
#define TS_PUBLIC __attribute__((visibility("default")))
#endif

TS_PUBLIC const TSLanguage *tree_sitter_blom(void) {
  static const TSLanguage language = {
    .version = LANGUAGE_VERSION,
    .symbol_count = SYMBOL_COUNT,
    .alias_count = ALIAS_COUNT,
    .token_count = TOKEN_COUNT,
    .external_token_count = EXTERNAL_TOKEN_COUNT,
    .state_count = STATE_COUNT,
    .large_state_count = LARGE_STATE_COUNT,
    .production_id_count = PRODUCTION_ID_COUNT,
    .field_count = FIELD_COUNT,
    .max_alias_sequence_length = MAX_ALIAS_SEQUENCE_LENGTH,
    .parse_table = &ts_parse_table[0][0],
    .small_parse_table = ts_small_parse_table,
    .small_parse_table_map = ts_small_parse_table_map,
    .parse_actions = ts_parse_actions,
    .symbol_names = ts_symbol_names,
    .symbol_metadata = ts_symbol_metadata,
    .public_symbol_map = ts_symbol_map,
    .alias_map = ts_non_terminal_alias_map,
    .alias_sequences = &ts_alias_sequences[0][0],
    .lex_modes = ts_lex_modes,
    .lex_fn = ts_lex,
    .primary_state_ids = ts_primary_state_ids,
  };
  return &language;
}
#ifdef __cplusplus
}
#endif
