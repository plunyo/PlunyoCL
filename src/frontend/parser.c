#include "frontend/parser.h"
#include "frontend/ast.h"
#include <stdlib.h>
#include <stdio.h>

ASTNode* GenerateAST(Parser* parser) {
    ASTNode* program = CreateASTNode(AST_PROGRAM);

    while (GetCurrentToken(parser)->type != TOKEN_EOF) {
        AddStatementToProgram(program, ParseStatement(parser));
    }

    return program;
}

Token* GetCurrentToken(Parser* parser) {
    return parser->currentToken;
}

void AdvanceToken(Parser* parser) {
    if (parser->currentToken->next) {
        parser->currentToken = parser->currentToken->next;
    }
}

Token* ExpectToken(Parser* parser, TokenType expectedType) {
    Token* current = GetCurrentToken(parser);

    if (current->type != expectedType) {
        fprintf(stderr, "Unexpected token: expected %d but got %d\n", expectedType, current->type);
        exit(EXIT_FAILURE);
    }

    AdvanceToken(parser);
    return current;
}

Parser* CreateParser(Lexer* lexer) {
    Parser* parser = (Parser*)malloc(sizeof(Parser));
    parser->lexer = lexer;
    parser->currentToken = lexer->tokens;
    return parser;
}

void DestroyParser(Parser* parser) {
    free(parser);
}