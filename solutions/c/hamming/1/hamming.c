#include "hamming.h"

int compute(const char *lhs, const char *rhs){
    if (strlen(lhs)!=strlen(rhs)){
        return -1;
    }
    int cnt = 0;
    while (*lhs!='\0'){
        if (*lhs!=*rhs){
            cnt++;
        }
        lhs++;
        rhs++;
    }
    return cnt;
}
