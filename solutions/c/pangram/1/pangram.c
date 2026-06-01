#include "pangram.h"

bool is_pangram(const char *sentence){
    if (sentence==NULL){
        return false;
    }
    int arr[26] = {0};
    for(size_t i=0;(unsigned char)sentence[i]!='\0';i++){
        if (isalpha((unsigned char)sentence[i])){
            char l = tolower((unsigned char)sentence[i]);
            arr[l-'a'] = 1;
        }
    }
    int j = 0;
    while (j<26){
        if (arr[j]<=0){
            return false;
        }
        j++;
    }
    return true;
}