#pragma once

#include <stddef.h>

typedef enum TokenType {
    TOKEN_IDENTIFIER, // 0
    TOKEN_NUMBER,     // 1
    TOKEN_STRING,     // 2
    TOKEN_PLUS,       // 3
    TOKEN_MINUS,      // 4
    TOKEN_ASTERISK,   // 5
    TOKEN_WHITESPACE, // 6
    TOKEN_SLASH,      // 7
    TOKEN_LPAREN,     // 8
    TOKEN_RPAREN,     // 9
    TOKEN_SEMICOLON,  // 10
    TOKEN_EQUALS,     // 11
    TOKEN_EOF,        // 12
} TokenType;

typedef struct Token {
    TokenType type;
    char* lexeme;
    size_t length;

    struct Token* next;
    struct Token* previous;
} Token;

Token* CreateToken(TokenType type, const char* lexeme, size_t length);
void DestroyToken(Token* token);