
#include "wiredtiger.h"
#include <stdlib.h>
#include <stdio.h>


#define error_check(call)                                              \
    do {                                                               \
        int __r;                                                       \
        if ((__r = (call)) != 0)                                        \
            printf("%d", __r);                                          \
    } while (0)

