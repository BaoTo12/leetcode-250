def merge(arr: list[int], left: int, mid: int, right: int) -> None:
    n1: int = mid - left + 1
    n2: int = right - mid

    L: list[int] = [0] * n1
    R: list[int] = [0] * n2
    for i in range(n1):
        L[i] = arr[left + i]
    for i in range(n2):
        R[i] = arr[mid + 1 + i]

    i = 0
    j = 0
    k = left
    while i < n1 and j < n2:
        if L[i] <= R[j]:
            arr[k] = L[i]
            i += 1
        else:
            arr[k] = R[j]
            j += 1
        k += 1
    while i < n1:
        arr[k] = L[i]
        i += 1
        k += 1

    while j < n2:
        arr[k] = R[j]
        j += 1
        k += 1


def mergeSort(arr: list[int], left: int, right: int) -> None:
    if left < right:
        mid: int = (left + right) // 2
        mergeSort(arr, left, mid)
        mergeSort(arr, mid + 1, right)
        merge(arr, left, mid, right)


if __name__ == "__main__":
    arr: list[int] = [38, 27, 43, 10]
    mergeSort(arr, 0, len(arr) - 1)
    for i in arr:
        print(i, end=" ")
    print()
