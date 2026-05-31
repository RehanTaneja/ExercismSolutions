#include "armstrong_numbers.h"

bool is_armstrong_number(int candidate){
    if (candidate<0){
        return false;
    }else if (candidate==0){
        return true;
    }
    int num = candidate;
    int numDigits = 0;
    while (num!=0){
        num = num / 10;
        numDigits++;
    }
    int num2 = candidate;
    int s = 0;
    while (candidate!=0){
        int d = candidate%10;
        s += pow(d,numDigits);
        candidate = candidate / 10;
    }
    return (num2==s);
}