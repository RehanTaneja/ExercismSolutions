#include "luhn.h"

bool luhn(const char *num){
    int digits = 0;
    int s = 0;
    bool dd = false;
    int i = strlen(num)-1;
    while(i>=0){
        if (isspace(num[i])){
            i--;
            continue;
        }else if (!isdigit(num[i])){
            return false;
        }
        int d = num[i] - '0';
        if (dd){
            d *= 2;
            if (d>9) d-=9;
        }
        dd = !dd;
        s += d;
        i--;
        digits++;
    }
    return digits > 1 && s%10==0;
}