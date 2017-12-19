#include "sort.h"


void print(int a[], int n){
	if (n > 200){
		return;
	}
	int i=0;
	for(; i<n ; i++)
	{
		printf("%3d ", a[i]);
		if (((i+1) % 25) == 0){
			printf("\n");
		}
	}
	printf("\n");
}

void set_all(int a[], int n)
{
	srand(time(NULL));
	
	int i=0;
	for(; i <n ; i++)
	{
		a[i] = rand() % 100 + 1;
	}
}

int sorted(int a[], int n){
	int i=1;
	for(; i<n ; i++)
	{
		if (a[i-1] > a[i]){
			printf("not sorted!\n");
			return 0;
		}
	}
	printf("ok \n");
	return 1;
}

void max_heapify(int a[], int n, int i)
{
	int left = 2*i + 1;
	int right = left + 1;
	int largest = i;
	
	while(1){
	
		left = 2*i + 1;
		right = left + 1;
		if (left < n && a[left] > a[largest]){
			largest = left;
		}
		
		if (right < n && a[right] > a[largest]){
			largest = right;
		}
		
		if (largest != i){
			Swap(a, largest, i);
			i = largest;
			//max_heapify(a, n, largest);
		}
		else{
			break;
		}
	}
	
}

void build_heap(int a[], int n)
{
	int i=0;
	for( i = n/2; i > -1 ; --i){
		max_heapify(a, n, i);
	}		
}

void heapsort(int a[], int n)
{
	build_heap(a, n);
	
	int i = n-1;
	for(; i > 0; --i)
	{
		Swap(a, i, 0);
		max_heapify(a, --n, 0);
	}
}

int main(){
	int src[MAX_SIZE];
	set_all(src, MAX_SIZE);
	
	heapsort(src, MAX_SIZE);
	
	sorted(src, MAX_SIZE);
	print(src, MAX_SIZE);
	
	
	return 0;
}