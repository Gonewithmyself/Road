#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define MAX_SIZE 100

#define Swap(src, a, b) do{ \
		int temp = src[a];  \
		src[a] = src[b];    \
		src[b] = temp;  \
}while (0)
	
void print(int a[], int n);
int sorted(int a[], int n);

void set_all(int a[], int n);
