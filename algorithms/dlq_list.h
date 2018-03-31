#ifndef DLQ_LIST_H_INCLUDED
#define DLQ_LIST_H_INCLUDED
#ifdef __cplusplus
extern "C"{
#endif // __cplusplus


#define ELEM_LEN(arr) sizeof((arr))/sizeof((arr)[0])

#define OFFSET_OF(type, mem) (long)(&(((type*)0)->mem))
#define CONTAINER_OF(ptr, type, mem) (type*)((char*)(ptr)-OFFSET_OF(type, mem))

#define list_for_each(list, iter) for(iter = (list)->head; iter; iter = iter->next)

typedef struct _dlq_node_t{
    struct _dlq_node_t *next;
    struct _dlq_node_t *prev;
}dlq_node;


typedef struct _dlq_list_t{
    dlq_node *head;
    dlq_node *tail;
    int len;
}dlq_list;

typedef void(*free_func_t)(dlq_node*);


dlq_list* dlq_list_init();
void dlq_list_clear(dlq_list* list, free_func_t free_func);
dlq_list* dlq_list_append(dlq_list* list, dlq_node* node);
dlq_list* dlq_list_append_front(dlq_list* list, dlq_node* node);
dlq_list* dlq_list_remove(dlq_list* list, dlq_node* node);

#define dlq_list_head(list) (list)->head
#define dlq_list_tail(list) (list)->tail

#define Malloc(ptr, size) do{   \
    (ptr) = malloc(size);       \
    if((ptr) == NULL){          \
        return NULL;            \
    }                           \
}while(0)

#define dlq_node_init(node) do{ \
    (node)->next = (node)->prev = NULL;\
}while(0)



#ifdef __cplusplus
}
#endif // __cplusplus
#endif // DLQ_LIST_H_INCLUDED
