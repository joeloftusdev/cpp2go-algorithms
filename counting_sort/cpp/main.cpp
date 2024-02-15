#include <iostream>
#include <vector>
#include <cstring>

// Function to perform count sort
void countsort(int arr[], int n, int k)
{
    // Create a vector to store the sorted array
    std::vector<int> output(n);

    // Create a vector to store frequency of each element, initialized with zeros
    std::vector<int> freq(k + 1, 0);

    // 1. Count the frequency of each element in the input array
    for (int i = 0; i < n; i++) {
        freq[arr[i]]++;
    }

    // 2. Calculate cumulative frequency
    int total = 0;
    for (int i = 0; i < k + 1; i++) {
        int oldCount = freq[i];
        freq[i] = total;
        total += oldCount;
    }

    // 3. Place elements in correct position in the output array
    for (int i = 0; i < n; i++) {
        output[freq[arr[i]]] = arr[i];
        freq[arr[i]]++;
    }

    // Copy the sorted elements back to the original array
    for (int i = 0; i < n; i++) {
        arr[i] = output[i];
    }
}

int main()
{
    // Input array
    int arr[] = { 4, 2, 10, 10, 1, 4, 2, 1, 10 };

    // Calculate the size of the array
    int n = sizeof(arr) / sizeof(arr[0]);

    // Range of array elements
    int k = 10;

    // Perform count sort
    countsort(arr, n, k);

    // Print the sorted array
    for (int i = 0; i < n; i++) {
        std::cout << arr[i] << " ";
    }
    std::cout << std::endl;

    return 0;
}
