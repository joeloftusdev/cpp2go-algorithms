#include <iostream>
#include <cstdlib>
#include <ctime>

#define N 15

// Merge two sorted subarrays `arr[low … mid]` and `arr[mid+1 … high]`
void Merge(int arr[], int aux[], int low, int mid, int high)
{
    int k = low, i = low, j = mid + 1;

    while (i <= mid && j <= high)
    {
        if (arr[i] <= arr[j]) {
            aux[k++] = arr[i++];
        }
        else {
            aux[k++] = arr[j++];
        }
    }

    while (i <= mid) {
        aux[k++] = arr[i++];
    }

    for (int i = low; i <= high; i++) {
        arr[i] = aux[i];
    }
}

void mergesort(int arr[], int aux[], int low, int high)
{
    if (high <= low) {
        return;
    }

    int mid = low + ((high - low) >> 1);

    mergesort(arr, aux, low, mid);
    mergesort(arr, aux, mid + 1, high);

    Merge(arr, aux, low, mid, high);
}

int isSorted(int arr[])
{
    for (int i = 1; i < N; i++)
    {
        if (arr[i - 1] > arr[i])
        {
            std::cout << "MergeSort Fails!!";
            return 0;
        }
    }

    return 1;
}

int main()
{
    int arr[N], aux[N];
    srand(time(NULL));

    for (int i = 0; i < N; i++) {
        aux[i] = arr[i] = (rand() % 100) - 50;
    }

    mergesort(arr, aux, 0, N - 1);

    if (isSorted(arr))
    {
        for (int i = 0; i < N; i++) {
            std::cout << arr[i] << " ";
        }
    }

    return 0;
}
