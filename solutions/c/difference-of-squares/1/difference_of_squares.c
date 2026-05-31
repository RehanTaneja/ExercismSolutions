#include "difference_of_squares.h"

unsigned int square_of_sum(unsigned int n){
    unsigned int i = 1;
    unsigned int sum = 0;
    while(i<=n){
        sum+=i;
        i++;
    }
    return sum*sum;
}

unsigned int sum_of_squares(unsigned int n){
    unsigned int i = 1;
    unsigned int sum = 0;
    while (i<=n){
        sum+=(i*i);
        i++;
    }
    return sum;
}

unsigned int difference_of_squares(unsigned int n){
    unsigned int n1 = square_of_sum(n);
    unsigned int n2 = sum_of_squares(n);
    return n1-n2;
}