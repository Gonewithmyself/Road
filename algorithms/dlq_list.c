#ifdef __cplusplus
extern "C"{
#endif // __cplusplus

#include "dlq_list.h"
#include <stdio.h>


dlq_list* dlq_list_init()
{
    dlq_list *list = malloc(sizeof(dlq_list));
    list->head=list->head=NULL;
    list->len=0;
    return list;
}

void dlq_list_clear(dlq_list* list, free_func_t free_func)
{
    if (list == NULL){
        return ;
    }

    dlq_node* node = list->head;
    while(node){
        dlq_node *cur = node;
        node = node->next;
        dlq_list_remove(list, cur);
        free_func(cur);
    }
    return;
}

dlq_list* dlq_list_append(dlq_list* list, dlq_node* node){
    if (list == NULL || node == NULL){
        return NULL;
    }
    if(list->len == 0){
        list->head = list->tail = node;
        node->next = node->prev = NULL;
    }else{
        dlq_node *last = list->tail;
        last->next = node;
        node->prev = last;
        list->tail = node;
        node->next = NULL;
    }
    list->len += 1;
    return list;
}

dlq_list* dlq_list_append_front(dlq_list* list, dlq_node* node){
    if (list == NULL || node == NULL){
        return NULL;
    }
    if(list->len == 0){
        list->head = list->tail = node;
        node->next = node->prev = NULL;
    }else{
        dlq_node *first = list->head;
        node->next = first;
        node->prev = NULL;
        list->head = node;
        first->prev = node;
    }
    list->len += 1;
    return list;
}
dlq_node* dlq_list_insert(dlq_list* list, dlq_node* pos, dlq_node *node)
{
    if (list == NULL || node == NULL || pos == NULL){
        return NULL;
    }

    dlq_node *next = pos->next;
    if(next){
        next->prev = node;
        pos->next = node;
        node->prev = pos;
        node->next = next;
    }else{
        dlq_list_append(list, node);
    }
    return node;
}
dlq_list* dlq_list_remove(dlq_list* list, dlq_node* node){
    if (list == NULL || list->len == 0 || node == NULL){
        return NULL;
    }

    dlq_node *prev = node->prev, *next = node ->next;
    if(prev){
        prev->next = next;
    }else{
        list ->head = next;
    }
    if(next){
        next -> prev = prev;
    }
    else{
        list ->tail = prev;
    }

    node->next = node->prev = NULL;
    list->len -= 1;
    return list;
}

#ifdef __cplusplus
}
#endif // __cplusplus
