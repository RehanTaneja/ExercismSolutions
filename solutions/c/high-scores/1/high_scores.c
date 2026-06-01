#include "high_scores.h"

int32_t latest(const int32_t *scores, size_t scores_len){
    return scores[scores_len-1];
}

int32_t personal_best(const int32_t *scores, size_t scores_len){
    size_t i = 1;
    int32_t best = scores[0];
    while (i<scores_len){
        if (scores[i]>best){
            best = scores[i];
        }
        i++;
    }
    return best;
}

size_t personal_top_three(const int32_t *scores, size_t scores_len, int32_t *output){
    int32_t best = INT32_MIN;
    int32_t best2 = INT32_MIN;
    int32_t best3 = INT32_MIN;
    size_t i = 0;
    while (i<scores_len){
        if (scores[i]>best){
            best3 = best2;
            best2 = best;
            best = scores[i];
        }else if (scores[i]>best2){
            best3 = best2;
            best2 = scores[i];
        }else if (scores[i]>best3){
            best3 = scores[i];
        }
        i++;
    }
    size_t cnt = 0;
    if (scores_len>0){
        output[0] = best;
        cnt++;
    }
    if (scores_len>1){
        output[1] = best2;
        cnt++;
    }
    if (scores_len>2){
        output[2] = best3;
        cnt++;
    }
    return cnt;
}