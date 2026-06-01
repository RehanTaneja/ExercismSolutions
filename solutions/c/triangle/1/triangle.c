#include "triangle.h"


bool is_valid_triangle(const triangle_t sides){
    if (sides.a==0 || sides.b==0 || sides.c==0){
        return false;
    }
    bool s1 = (sides.a + sides.b)>=sides.c;
    bool s2 = (sides.a + sides.c)>=sides.b;
    bool s3 = (sides.b + sides.c)>=sides.a;
    return s1 && s2 && s3;
}

bool is_equilateral(const triangle_t sides){
    return is_valid_triangle(sides) && (sides.a==sides.b && sides.b==sides.c);
}

bool is_isosceles(const triangle_t sides){
    return is_valid_triangle(sides) && (sides.a==sides.b || sides.b==sides.c || sides.a==sides.c);
}

bool is_scalene(const triangle_t sides){
    return is_valid_triangle(sides) && !is_isosceles(sides);
}
