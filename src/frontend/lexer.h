#pragma once

#include <stddef.h>
#include "frontend/token.h"

typedef struct Lexer {
    Token* tokens;
    size_t tokenCount;
} Lexer;

Lexer* CreateLexer();
void AddToken(Lexer* lexer, Token* newToken);
void ParseTokens(Lexer* lexer, const char* source, size_t length);
void PrintTokens(Lexer* lexer); 
void DestroyLexer(Lexer* lexer);
