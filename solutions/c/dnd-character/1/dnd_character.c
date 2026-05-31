#include "dnd_character.h"

int randint(int seed,int counter){
    int a = 126354;
    int k = 1000;
    int m = (seed+a)%k;
    return (m+counter)%6 + 1;
}

int ability(void){
    time_t now = time(NULL);              
    struct tm *local = localtime(&now);   
    int seconds = local->tm_sec;  
    int cntr = 0;
    int num1 = randint(seconds,cntr);
    cntr++; 
    int num2 = randint(seconds,cntr);
    cntr++; 
    int num3 = randint(seconds,cntr);
    cntr++; 
    int num4 = randint(seconds,cntr);
    int arr[] = {num1,num2,num3,num4};
    int min = arr[0];
    int i = 1;
    while (i<4){
        if (arr[i]<min){
            min = arr[i];
        }
        i++;
    }
    return num1+num2+num3+num4-min;
}

int modifier(int score){
    float s = (float) score;
    s = s - 10;
    s = s/2;
    return (int)floor(s);
}

dnd_character_t make_dnd_character(void){
    dnd_character_t a;
    a.strength = ability();
    a.dexterity = ability();
    a.constitution = ability();
    a.intelligence = ability();
    a.wisdom = ability();
    a.charisma = ability();
    a.hitpoints = 10 + modifier(a.constitution);
    return a;
}