#include "binary.h"

int convert(const char *input){
    int answer = 0;
    int len = strlen(input)-1;
    int i = 0;
    while (input[i]!='\0'){
        if (input[i]=='1'){
            answer += pow(2,len);
        }else if (input[i]!='0'){
            return INVALID;
        }
        i++;
        len--;
    }
    return answer;
}