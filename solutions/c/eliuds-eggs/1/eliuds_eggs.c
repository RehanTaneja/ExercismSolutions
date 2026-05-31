#include "eliuds_eggs.h"

int egg_count(int eggs){
    int cnt = 0;
    while (eggs){
        cnt += eggs & 1;
        eggs >>= 1;
    }
    return cnt;
}