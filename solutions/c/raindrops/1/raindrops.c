#include "raindrops.h"

void convert(char result[], int drops){
    if (drops%3==0){
        sprintf(result+strlen(result),"Pling");
    }
    if (drops%5==0){
        sprintf(result+strlen(result),"Plang");
    }
    if (drops%7==0){
        sprintf(result+strlen(result),"Plong");
    }
    if (drops%3!=0 && drops%5!=0 && drops%7!=0){
        sprintf(result,"%d",drops);
    }
}