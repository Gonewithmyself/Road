#ifndef TEST_H_INCLUDED
#define TEST_H_INCLUDED

#include <stdlib.h>
#include <stdio.h>
#include <time.h>
#include "dlq_list.h"

void test();
void free_stu(dlq_node *node);

typedef struct _student{
    int id;
    char name[10];
    dlq_node node;
}student;

typedef union item_{
    int id;
    char name;
}item;


#endif // TEST_H_INCLUDED
