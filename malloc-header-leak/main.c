#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main() {
    malloc(72);
    long* ptr = (long*)malloc(sizeof(long)*9);
    *(ptr -1) = 32;
    free(ptr);
    sleep(2);
    for (int i = 0 ; i < 12 ; i++) {
        printf("%ld\n", *(ptr-i+8));
    }
    return 0;
}
