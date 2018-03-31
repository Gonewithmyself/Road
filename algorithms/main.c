#ifdef __cplusplus
extern "C"{
#endif // __cplusplus
#include "test.h"

int main()
{
    test();
    return 0;
}

void test_item()
{
    item it;

    it.id = 0x12345678;

    printf("%x, %x\n", it.id, it);

    it.name = 'a';
    printf("%d, %x\n", it.name, it);
}

void test()
{
    int *a, n = 2000000;

    a = get_array(n);
    check(a, n);
    quick_sort(a, n);
    //heap_sort(a, n);
    check(a, n);
}

void test_hash()
{
    int seed = hash_table_seed(1313131);
    printf("%d\n", seed);
}

void test_dlq_list()
{
    int a[] = {1,2,3,4,5, 6};
    dlq_list *list = build_list(a, ELEM_LEN(a));
    print_list(list);

    dlq_node *node = find_node(list, 5);
    dlq_list_remove(list, node);
    print_list(list);

    student *stu = init_student(99);
    node = find_node(list, 3);
    dlq_list_insert(list, list->tail, &stu->node);
    print_list(list);

    dlq_list_clear(list, free_stu);
    print_list(list);

}
#ifdef __cplusplus
extern "C"{
#endif // __cplusplus
