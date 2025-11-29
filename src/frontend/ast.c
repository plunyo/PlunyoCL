#include "frontend/ast.h"

#include <stdlib.h>

ASTNode* CreateASTNode(ASTKind kind) {
    ASTNode* node = (ASTNode*)malloc(sizeof(ASTNode));
    node->kind = kind;

    switch (kind) {
        case AST_PROGRAM:
            node->program.statements = NULL;
            node->program.statementCount = 0;
            node->program.statementCapacity = 0;
            break;
        case AST_STATEMENT:
            node->statement.expr = NULL;
            break;
        case AST_NUMBER:
            node->number.value = 0.0;
            break;
        case AST_IDENTIFIER:
            node->identifier.name = NULL;
            break;
        case AST_BINARY:
            node->binary.op = '\0';
            node->binary.left = NULL;
            node->binary.right = NULL;
            break;
        case AST_UNARY:
            node->unary.op = '\0';
            node->unary.expr = NULL;
            break;
        case AST_DECLARATION:
            node->declaration.name = NULL;
            node->declaration.value = NULL;
            break;
        case AST_ASSIGN:
            node->assign.name = NULL;
            node->assign.value = NULL;
            break;
    }

    return node;
}

void DestroyAST(ASTNode* node) {
    if (!node) return;

    switch (node->kind) {
        case AST_PROGRAM:
            for (size_t i = 0; i < node->program.statementCount; i++) {
                DestroyAST(node->program.statements[i]);
            }
            free(node->program.statements);
            break;
        case AST_STATEMENT:
            DestroyAST(node->statement.expr);
            break;
        case AST_BINARY:
            DestroyAST(node->binary.left);
            DestroyAST(node->binary.right);
            break;
        case AST_UNARY:
            DestroyAST(node->unary.expr);
            break;
        case AST_DECLARATION:
            free(node->declaration.name);
            DestroyAST(node->declaration.value);
            break;
        case AST_ASSIGN:
            free(node->assign.name);
            DestroyAST(node->assign.value);
            break;
        case AST_IDENTIFIER:
            free(node->identifier.name);
            break;
        case AST_NUMBER:
            // what in the heck do i free
            break;
    }

    free(node);
}

void AddStatementToProgram(ASTNode* program, ASTNode* statement) {
    if (program->program.statementCount >= program->program.statementCapacity) {
        program->program.statementCapacity = program->program.statementCapacity == 0 ? 4 : program->program.statementCapacity * 2;
        program->program.statements = realloc(program->program.statements, program->program.statementCapacity * sizeof(ASTNode*));
    }
    program->program.statements[program->program.statementCount++] = statement;
}