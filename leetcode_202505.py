import functools
import heapq
import itertools
import math


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
                # Calculate the time to reach this new position
                t = max(moveTime[x][y], dist[i][j]) + 1
                # If found a shorter path to (x, y), update distance and push to queue
                if dist[x][y] > t:
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


# assert minTimeToReach([[0, 4], [4, 4]]) == 6

# assert romanToInt("LVIII") == 58
# assert romanToInt("MCMXCIV") == 1994
