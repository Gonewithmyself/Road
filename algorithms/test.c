#ifdef __cplusplus
extern "C"{
#endif // __cplusplus
#include "test.h"

student *init_student(int id)
{
    student *stu = (student*)malloc(sizeof(student));
    stu->id = id;
    dlq_node_init(&stu->node);
    return stu;
}

dlq_list* build_list(int a[], int n){
    dlq_list *list = dlq_list_init();
    student *stu;
    for(int i=0; i<n ; i++){
        stu = init_student(a[i]);
        dlq_list_append_front(list, &stu->node);
    }
    return list;
}

void print_list(dlq_list* list)
{
    if(list == NULL){
        return ;
    }
    student *stu;
    dlq_node *iter = dlq_list_head(list);
    for(; iter; iter= iter->next){
        stu = CONTAINER_OF(iter, student, node);
        printf("%d ", stu->id);
    }
    printf("\nlist len: %d\n", list->len);
}

dlq_node* find_node(dlq_list* list, int val)
{
    student *stu;
    dlq_node *iter = dlq_list_head(list);
    for(; iter; iter= iter->next){
        stu = CONTAINER_OF(iter, student, node);
        if(stu->id == val){
            return iter;
        }
    }
    return NULL;
}
void free_stu(dlq_node *node)
{
    if(node){
        free(CONTAINER_OF(node, student, node));
    }
}



int* get_array(int num)
{
    int *a = malloc(sizeof(int)*num);

    if(a)
    {
        srand(time(NULL));
        for(int i = 0; i<num; i++)
        {
            a[i] = rand() % num + 1;
        }
    }
    return a;
}

void print(int *a, int num)
{
    if(num > 1000)
        return;
    for(int i=0;i<num;i++)
    {
        printf("%-3d ", a[i]);
        if((i+1) % 20 == 0){
            //printf("\n");
        }
    }
    printf("\n");
}

void sorted(int *a, int num)
{
    for(int i=1;i<num;i++)
    {
        if(a[i-1]>a[i]){
            printf("Not sorted!\n");
            return;
        }
    }
    printf("Sorted.\n");
}

void check(int *a, int n)
{
    print(a, n);
    sorted(a, n);
}
#ifdef __cplusplus
}
#endif // __cplusplus
