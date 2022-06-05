#include "mathlib.h"

uint32_t add(uint32_t v1,uint32_t v2) {
  return v1+v2;
}

int main() {
    printf("C is working \n");

    int answer = add(1,2);
    
    printf("add function working: Value of 1+2 = %d \n",answer);

    return 0;
}