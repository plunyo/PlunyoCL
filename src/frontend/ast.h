#pragma once

#include <stddef.h>

typedef enum ASTKind {
    AST_PROGRAM,
    AST_STATEMENT,
    AST_NUMBER,
    AST_IDENTIFIER,
    AST_BINARY,
    AST_UNARY,
    AST_DECLARATION,
    AST_ASSIGN,
} ASTKind;

typedef struct ASTNode ASTNode;

typedef struct ASTNode {
    ASTKind kind;

    union {
        struct {
            ASTNode** statements;
            size_t statementCount;
            size_t statementCapacity;
        } program;

        struct {
            ASTNode* expr;
        } statement;

        struct {
            double value;
        } number;

        struct {
            char* name;
        } identifier;

        struct {
            char op;   // '+', '-', '*', '/'
            ASTNode* left;
            ASTNode* right;
        } binary;

        struct {
            char op;   // '-', '!'
            ASTNode* expr;
        } unary;

        struct {
            char* name;
            ASTNode* value;
        } declaration;

        struct {
            char* name;
            ASTNode* value;
        } assign;
    };
} ASTNode;

ASTNode* CreateASTNode(ASTKind kind);
void DestroyAST(ASTNode* node);

void AddStatementToProgram(ASTNode* program, ASTNode* statement);