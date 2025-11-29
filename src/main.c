#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "frontend/lexer.h"

void readFile(const char* filename, char** buffer, size_t* length) {
    FILE* file = fopen(filename, "r");

    if (!file) {
        *buffer = NULL;
        *length = 0;
        return;
    }

    fseek(file, 0, SEEK_END);
    *length = ftell(file);

    fseek(file, 0, SEEK_SET);
    *buffer = (char*)malloc(*length + 1);

    fread(*buffer, 1, *length, file);
    (*buffer)[*length] = '\0';
    
    fclose(file);
}

void repl(Lexer* lexer) {
    char line[1024];

    while (1) {
        printf("pcl > ");
        if (!fgets(line, sizeof(line), stdin) || strcmp(line, "exit\n") == 0) {
            break;
        }

        ParseTokens(lexer, line, strlen(line));
        PrintTokens(lexer);
    }
} 

int main(int argc, char* argv[]) {
    // if file as arg
    if (argc > 2) {
        printf("usage: %s <source_file>\n", argv[0]);
        return 1;
    }

    Lexer* lexer = CreateLexer();

    if (argc == 2) {
        const char* sourceFile = argv[1];
        char* source;
        size_t length;
        
        readFile(sourceFile, &source, &length );

        if (!source) {
            printf("Could not read file: %s\n", sourceFile);
            return 1;
        }

        ParseTokens(lexer, source, length);
        PrintTokens(lexer);
    } else {
        repl(lexer);
    }

    DestroyLexer(lexer);

    return 0;
}
