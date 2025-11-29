#pragma once

#include "frontend/ast.h"
#include "frontend/lexer.h"
#include "frontend/token.h"

typedef struct Parser {
    Lexer* lexer;
    Token* currentToken;
} Parser;

Parser* CreateParser(Lexer* lexer);
void DestroyParser(Parser* parser);

ASTNode* ParseStatement(Parser* parser);

Token* GetCurrentToken(Parser* parser);
Token* ExpectToken(Parser* parser, TokenType expectedType);
void AdvanceToken(Parser* parser);

ASTNode* GenerateAST(Parser* parser);
void DestroyAST(ASTNode* node);