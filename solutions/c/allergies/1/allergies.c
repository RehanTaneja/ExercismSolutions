#include "allergies.h"

bool is_allergic_to(allergen_t allergen, int num){
    num &= 255;
    switch (allergen){
        case ALLERGEN_EGGS:
            return num & 1;
        case ALLERGEN_PEANUTS:
            return num & 2;
        case ALLERGEN_SHELLFISH:
            return num & 4;
        case ALLERGEN_STRAWBERRIES:
            return num & 8;
        case ALLERGEN_TOMATOES:
            return num & 16;
        case ALLERGEN_CHOCOLATE:
            return num & 32;
        case ALLERGEN_POLLEN:
            return num & 64;
        case ALLERGEN_CATS:
            return num  & 128;
        default:
            return false;
    }
}

allergen_list_t get_allergens(int num){
    int cnt = 0;
    allergen_list_t a = {0};
    num &= 255;
    for (int i = 0;i<8;i++){
        if (num & (1<<i)){
            a.allergens[i]=true;
            cnt++;
        }
    }
    a.count = cnt;
    return a;
}