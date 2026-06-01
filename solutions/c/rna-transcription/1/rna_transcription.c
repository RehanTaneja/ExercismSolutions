#include "rna_transcription.h"

char *to_rna(const char *dna){
    char *answer = malloc(strlen(dna)+1);
    int i = 0;
    while (*dna!='\0'){
        switch (*dna){
            case 'G':
                answer[i] = 'C';
                break;
            case 'C':
                answer[i] = 'G';
                break;
            case 'T':
                answer[i] = 'A';
                break;
            case 'A':
                answer[i] = 'U';
                break;
        }
        i++;
        dna++;
    }
    answer[i] = '\0';
    return answer;
}