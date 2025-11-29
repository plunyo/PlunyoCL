#include "frontend/lexer.h"
#include "frontend/token.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

Lexer* CreateLexer() {
    Lexer* lexer = (Lexer*)malloc(sizeof(Lexer));

    lexer->tokens = NULL;
    lexer->tokenCount = 0;

    return lexer;
}

void AddToken(Lexer* lexer, Token* newToken) {
    if (!lexer->tokens) {
        lexer->tokens = newToken;
    } else {
        Token* current = lexer->tokens;

        while (current->next) {
            current = current->next;
        }

        current->next = newToken;
        newToken->previous = current;
    }

    lexer->tokenCount++;
}

void ParseTokens(Lexer* lexer, const char* source, size_t length) {
    for (size_t i = 0; i < length; i++) {
        char currentChar = source[i];

        switch (currentChar) {
            case '+':
                AddToken(lexer, CreateToken(TOKEN_PLUS, "+", 1));
                break;
            case '-':
                AddToken(lexer, CreateToken(TOKEN_MINUS, "-", 1));
                break;
            case '*':
                AddToken(lexer, CreateToken(TOKEN_ASTERISK, "*", 1));
                break;
            case '/':
                AddToken(lexer, CreateToken(TOKEN_SLASH, "/", 1));
                break;
            case '(':
                AddToken(lexer, CreateToken(TOKEN_LPAREN, "(", 1));
                break;
            case ')':
                AddToken(lexer, CreateToken(TOKEN_RPAREN, ")", 1));
                break;
            case ';':
                AddToken(lexer, CreateToken(TOKEN_SEMICOLON, ";", 1));
                break;
            case '=':
                AddToken(lexer, CreateToken(TOKEN_EQUALS, "=", 1));
                break;
            case ' ':
            case '\t':
            case '\n': {
                size_t start = i;

                while (i + 1 < length &&
                    (source[i + 1] == ' ' ||
                    source[i + 1] == '\t' ||
                    source[i + 1] == '\n')) {
                    i++;
                }

                size_t ws_length = (i - start) + 1;

                char* ws = malloc(ws_length + 1);
                memcpy(ws, &source[start], ws_length);
                ws[ws_length] = '\0';

                AddToken(lexer, CreateToken(TOKEN_WHITESPACE, ws, ws_length));
                break;
            }
            default:
                if ((currentChar >= '0' && currentChar <= '9')) {
                    size_t start = i;

                    while (i < length && (source[i] >= '0' && source[i] <= '9')) {
                        i++;
                    }

                    AddToken(lexer, CreateToken(TOKEN_NUMBER, &source[start], i - start));
                    i--;
                } else if ((currentChar >= 'a' && currentChar <= 'z') || (currentChar >= 'A' && currentChar <= 'Z') || currentChar == '_') {
                    size_t start = i;

                    while (i < length && ((source[i] >= 'a' && source[i] <= 'z') || (source[i] >= 'A' && source[i] <= 'Z') || (source[i] >= '0' && source[i] <= '9') || source[i] == '_')) {
                        i++;
                    }

                    AddToken(lexer, CreateToken(TOKEN_IDENTIFIER, &source[start], i - start));
                    i--;
                }

                fprintf(stderr, "Unknown Token ( Type: %d | Lexeme: %s )", currentChar, &source[i]);
                break;
        }
    }

    AddToken(lexer, CreateToken(TOKEN_EOF, NULL, 0));
}

void PrintTokens(Lexer* lexer) {
    Token* current = lexer->tokens;

    while (current) {
        printf("(Type: %d | Lexeme: '%s' | Length: %zu)\n", current->type, current->lexeme, current->length);
        current = current->next;
    }
}

void DestroyLexer(Lexer* lexer) {
    Token* current = lexer->tokens;

    while (current) {
        Token* next = current->next;

        free(current->lexeme);
        free(current);

        current = next;
    }

    free(lexer);
}