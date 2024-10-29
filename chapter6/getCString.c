#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* getCString(char* str){
    // malloc(): 입력 문자열과 동일한 크기의 메모리를 동적으로 할당
    // +1: null 문자 포함
    char* result = (char*)malloc(strlen(str) + 1);
    if (result != NULL) {
        strcpy(result, str); // 할당 성공 후 문자열 복사
    }
    return result; 
}