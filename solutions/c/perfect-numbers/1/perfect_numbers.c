#include "perfect_numbers.h"

kind classify_number(int num){
    if (num<=0){
        return -1;
    }
    int s = 0;
    int i = 1;
    while (i<=num/2){
        if (num%i==0){
            s += i;
        }
        i++;
    }
    if (s==num){
        return PERFECT_NUMBER;
    }else if (s>num){
        return ABUNDANT_NUMBER;
    }
    return DEFICIENT_NUMBER;
}