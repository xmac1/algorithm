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