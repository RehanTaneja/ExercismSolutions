#include "gigasecond.h"

void gigasecond(time_t input, char *output, size_t size) {
    time_t future = input + GS;
    struct tm *tm_struct = gmtime(&future);  // Use UTC

    if (tm_struct != NULL) {
        strftime(output, size, "%Y-%m-%d %H:%M:%S", tm_struct);
    }
}