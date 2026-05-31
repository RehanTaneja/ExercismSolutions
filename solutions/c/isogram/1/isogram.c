#include "isogram.h"

bool is_isogram(const char phrase[]){
    if (phrase == NULL) return false;
    int a[26] = {0};
    for(int i = 0;phrase[i]!='\0';i++){
        if (isalpha((unsigned char)phrase[i])){
            char lower = tolower((unsigned char)phrase[i]);
            a[lower-'a']++;
            if (a[lower-'a']>1){
                return false;
            }
        }
    }
    return true;
}