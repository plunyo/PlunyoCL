#include "frontend/token.h"
#include <stdlib.h>

Token* CreateToken(TokenType type, const char* lexeme, size_t length) {
    Token* token = (Token*)malloc(sizeof(Token));

    token->type = type;
    token->length = length;
    token->next = NULL;

    token->lexeme = (char*)malloc(length + 1);
    for (size_t i = 0; i < length; i++) {
        token->lexeme[i] = lexeme[i];
    }
    token->lexeme[length] = '\0';
    

    return token;
}

void DestroyToken(Token* token) {
    free(token->lexeme);
    free(token);
}