import sys

def insertionSort(arr): 
    for i in range(1, len(arr)): 
        key = arr[i] 
        j = i-1
        while j >=0 and key < arr[j] : 
                arr[j+1] = arr[j] 
                j -= 1
        arr[j+1] = key

arr = sys.argv[1:]
arr = list(map(int, arr))
insertionSort(arr)
print(' '.join(map(str, arr)))
