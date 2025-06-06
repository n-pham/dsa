import collections
import functools
import heapq
import itertools
import math
import sortedcontainers


def romanToInt(s: str) -> int:
    # 13
    num = 0
    singleValue = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    doubleValue = {"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
    prev = ""
    for c in s + " ":
        if not prev:
            prev = c
            continue
        if prev + c in doubleValue:
            num += doubleValue[prev + c]
            prev = ""
        else:
            num += singleValue[prev]
            prev = c
    return num


def minTimeToReach(moveTime: list[list[int]]) -> int:
    # 3341
    n = len(moveTime)
    m = len(moveTime[0])
    INF = 1 << 30  # A large enough constant for our grid problems

    # Preallocate grid distances
    dist = [[INF] * m for _ in range(n)]
    dist[0][0] = 0

    pq = [(0, 0, 0)]
    # Directions: up, down, left, right
    DIRS = [(-1, 0), (1, 0), (0, -1), (0, 1)]

    while pq:
        d, i, j = heapq.heappop(pq)
        if i == n - 1 and j == m - 1:
            return d
        if d > dist[i][j]:
            continue
        for dx, dy in DIRS:
            x = i + dx
            y = j + dy
            if 0 <= x < n and 0 <= y < m:
                nxtTime = dist[i][j]
                cellTime = moveTime[x][y]
                # Avoid calling max() for small ints by explicit comparison, a bit faster
                t = cellTime + 1 if cellTime > nxtTime else nxtTime + 1
                if t < dist[x][y]:
                    dist[x][y] = t
                    heapq.heappush(pq, (t, x, y))


def minTimeToReach3342(moveTime: list[list[int]]) -> int:
    # 3342
    n, m = len(moveTime), len(moveTime[0])
    dist = [[math.inf] * m for _ in range(n)]
    dist[0][0] = 0

    # Priority queue for Dijkstra's algorithm starting at top-left corner
    pq = [(0, 0, 0)]

    dirs = (-1, 0, 1, 0, -1)

    while True:
        d, i, j = heapq.heappop(pq)  # Pop the smallest distance node

        # If we've reached the bottom-right corner, return the total distance
        if i == n - 1 and j == m - 1:
            return d

        # Skip if we've found a shorter path to (i, j) already
        if d > dist[i][j]:
            continue

        # Explore all 4 possible directions
        for a, b in itertools.pairwise(dirs):
            x, y = i + a, j + b
            # Check if the new position is within bounds
            if 0 <= x < n and 0 <= y < m:
                # New time is the greater of either the given cell's moving time
                # or the currently calculated time plus a constant based on parity of coordinates
                t = max(moveTime[x][y], dist[i][j]) + (i + j) % 2 + 1
                # If found a shorter path to (x, y), update distance and push to queue
                if dist[x][y] > t:
                    dist[x][y] = t
                    heapq.heappush(pq, (t, x, y))


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def mergeTwoLists(list1: ListNode, list2: ListNode) -> ListNode:
    # 21
    head = ls = ListNode()
    ls = head
    while list1 and list2:
        if list1.val < list2.val:
            ls.next = list1
            list1 = list1.next
        else:
            ls.next = list2
            list2 = list2.next
        ls = ls.next
    ls.next = list1 or list2
    return head.next


def triangleType(nums: list[int]) -> str:
    # 3024
    return (
        "none"
        if nums[0] + nums[1] <= nums[2]
        or nums[0] + nums[2] <= nums[1]
        or nums[2] + nums[1] <= nums[0]
        else {1: "equilateral", 2: "isosceles", 3: "scalene"}[len(set(nums))]
    )


def maxRemoval(nums: list[int], queries: list[list[int]]) -> int:
    # 3362
    q = collections.deque(sorted(queries))
    available = sortedcontainers.SortedList()  # available `r`s
    running = sortedcontainers.SortedList()  # running `r`s

    for i, num in enumerate(nums):
        while q and q[0][0] <= i:
            available.add(q.popleft()[1])
        while running and running[0] < i:
            running.pop(0)
        while num > len(running):
            if not available or available[-1] < i:
                return -1
            running.add(available.pop())

    return len(available)


assert triangleType([3, 3, 3]) == "equilateral"
assert triangleType([3, 4, 5]) == "scalene"
assert triangleType([3, 3, 5]) == "isosceles"
assert triangleType([3, 4, 9]) == "none"

# assert minTimeToReach([[0, 4], [4, 4]]) == 6

# assert romanToInt("LVIII") == 58
# assert romanToInt("MCMXCIV") == 1994
