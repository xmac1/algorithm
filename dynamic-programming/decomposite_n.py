from typing import Dict, Tuple

def decomposite_n(n: int, start: int, m: Dict[Tuple[int, int], int]):
    """
    top-down method with memorization for dynamic programming
    """
    total = 0
    if (n, start) in m:
        return m[(n, start)]
    if n == 1 and n > start:
        m[(1, 1)] = 1
        return 1
    for i in range(start, int(n/2)+1):
        if i == 0:
            total += 1
            continue
        if n-i < i:
            break
        total += 1
        total += decomposite_n(n-i, i, m)
    m[(n, start)] = total
    return total

def bottom_up_decomposite_n(n: int):
    """
    bottom-up method for dynamic programming
    """
    m = dict()
    for i in range(n+1):
        m[(0, i)] = 0
        m[(i,i)] = 0
    for i in range(1, int(n/2)+1):
        for j in range(i, n - i):
            if j > i:
                




                

if __name__ == "__main__":
    bottom_up_decomposite_n(5)