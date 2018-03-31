#ifdef __cplusplus
extern "C"{
#endif // __cplusplus

#include "test.h"

void _swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

void insert_sort(int *a, int n)
{
    for(int i=1; i<n ; i++)
    {
        int temp = a[i], j=i;
        for(;j>0;j--)
        {
            if(a[j-1]>temp)
            {
                a[j] = a[j-1];
            }
            else{
                break;
            }
        }
        a[j] = temp;
    }
}

void max_heapfiy(int *a, int n , int i)
{
    int left = 2*i ;
    int right = left + 1;
    int largest = i;

    if (left < n && a[left] > a[largest])
    {
        largest  = left;
    }
    if (right < n && a[right] > a[largest])
    {
        largest  = right;
    }

    if( largest != i)
    {
        int temp = a[i];
        a[i] = a[largest];
        a[largest] = temp;
        max_heapfiy(a, n, largest);
    }
}

void build_heap(int *a, int n)
{
    for(int i=n/2; i>-1 ; i--)
    {
        max_heapfiy(a, n, i);
    }
}

void heap_sort(int *a, int n)
{
    build_heap(a, n);
    for(int i = 1; i < n; i++ )
    {
        int temp = a[0];
        a[0] = a[n-i];
        a[n-i] = temp;
        max_heapfiy(a, n-i, 0);
    }
}

void merge(int *a, int p, int q, int r)
{
    int m = q - p + 1 , n = r-q;
    int *L =malloc(sizeof(int)*m), *R = malloc(sizeof(int)*n);


    for(int i=0; i<m;i++)
    {
        L[i] = a[p+i];
    }
    for(int j=0; j<n;j++)
    {
        R[j] = a[q+j+1];
    }/*
    printf("A:");
    //print(a, 20);
    printf("m= %d, n=%d, p=%d q=%d r=%d\n", m, n, p, q, r);
    printf("L: ");
    //print(L, m);
    printf("R: ");
    print(R, n);
    //printf("L=%d, R=%d\n", L[0], R[0]);*/

    int i = 0, j = 0, k = p;

    for(;i<m && j<n;k++)
    {
        if(L[i]<=R[j])
        {
            //printf("i = %d, %d, %d, %d\n", i, j, L[i], k);
            a[k] = L[i];
            i++;
        }
        else
        {
            //printf("j = %d, %d, %d, %d\n", i, j, R[j], k);
            a[k] = R[j];
            j++;
        }
    }

    while(i<m)
    {
        a[k++] = L[i++];
    }
    while(j<n)
    {
        a[k++] = R[j++];
    }

}

void merge_sort(int *a, int p, int r)
{
    if(p < r)
    {
        int q = (p+r)/2;
        merge_sort(a, p, q);
        merge_sort(a, q+1, r);
        merge(a, p, q, r);
    }
}

void mergesort(int *a, int n)
{
    merge_sort(a, 0, n-1);
}

int _qsort(int *a, int i, int j)
{
    //if(i<j)
    {
        int k = a[i], s=i;
        j++;
        do
        {
            while(a[++i]<k);
            while(a[--j]>k);
            if(i<j)
            {
                int temp = a[i];
                a[i] = a[j];
                a[j] = temp;
            }
        }while(i<j);
        _swap(&a[s], &a[j]);
        return j;
        //_qsort(a, s, j-1);
        //_qsort(a, j+1, e);
    }
}

int _qsort1(int *a, int p, int r)
{
    int x = a[r], i = p-1, j = p;

    for(;j < r; j++)
    {
        if(a[j]<=x)
        {
            i++;
            _swap(&a[i], &a[j]);
        }
    }
    _swap(&a[i+1], &a[r]);
    return i+1;
}

void _quick_sort(int *a, int p, int r)
{
    while(p < r)
    {
        int q = _qsort1(a, p, r);
        _quick_sort(a, p, q-1);
        p = q+1;
    }
}

void quick_sort(int *a, int n)
{
    _quick_sort(a, 0, n-1);
}

#ifdef __cplusplus
extern "C"{
#endif // __cplusplus
