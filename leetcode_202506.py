import collections
import string


def maxCandies(
    status: list[int],
    candies: list[int],
    keys: list[list[int]],
    containedBoxes: list[list[int]],
    initialBoxes: list[int],
) -> int:
    # 1298
    # Track total candies collected
    total_candies = 0

    # Track boxes we have but can't open yet
    found_boxes = set()
    # Track keys we have
    found_keys = set()
    # Queue for BFS
    queue = collections.deque()

    # Add initial boxes to queue if open, otherwise to found_boxes
    for box in initialBoxes:
        if status[box] == 1:
            queue.append(box)
        else:
            found_boxes.add(box)

    # BFS to process boxes
    while queue:
        box = queue.popleft()

        # Collect candies from current box
        total_candies += candies[box]

        # Add keys found in current box
        for key in keys[box]:
            found_keys.add(key)
            # If we have the box for this key but couldn't open it before
            if key in found_boxes:
                queue.append(key)
                found_boxes.remove(key)

        # Process boxes found inside current box
        for new_box in containedBoxes[box]:
            if status[new_box] == 1 or new_box in found_keys:
                # Can open this box immediately
                queue.append(new_box)
            else:
                # Store for later when we get the key
                found_boxes.add(new_box)

    return total_candies


def clearStars(self, s: str) -> str:
    # Dictionary to keep track of the indices of each alphabet character.
    char_indices = collections.defaultdict(list)
    n = len(s)  # Length of the input string.
    remove = [False] * n  # Boolean array to mark characters for removal.
    for i, char in enumerate(s):
        if char == "*":
            # Mark '*' for removal.
            remove[i] = True
            # Attempt to find the most recent alphabet character to remove.
            for alphabet in string.ascii_lowercase:
                if char_indices[alphabet]:
                    # Mark the most recent occurrence for removal and break loop.
                    remove[char_indices[alphabet].pop()] = True
                    break
        else:
            # Record the index of the current alphabet character.
            char_indices[char].append(i)
    # Return the filtered string after removing marked characters.
    return "".join(char for i, char in enumerate(s) if not remove[i])


def maxDifference(s: str, k: int) -> float:
    def get_status(a, b):
        return ((a & 1) << 1) | (b & 1)

    n = len(s)
    ans = float("-inf")

    for ca in range(5):
        for cb in range(5):
            if ca == cb:
                continue
            best = [float("inf")] * 4
            cnt_a = cnt_b = prev_a = prev_b = 0
            left = -1
            for right in range(n):
                cnt_a += s[right] == str(ca)
                cnt_b += s[right] == str(cb)
                while right - left >= k and cnt_b - prev_b >= 2:
                    status = get_status(prev_a, prev_b)
                    best[status] = min(best[status], prev_a - prev_b)
                    left += 1
                    prev_a += s[left] == str(ca)
                    prev_b += s[left] == str(cb)
                status = get_status(cnt_a, cnt_b)
                if best[status ^ 2] < float("inf"):
                    ans = max(ans, cnt_a - cnt_b - best[status ^ 2])

    return ans

def partitionArray(nums: list[int], k: int) -> int:
    # 2294
    nums.sort()
    ans = 1
    mn = nums[0]
    for i in range(1, len(nums)):
        if nums[i] - mn > k:
            ans += 1
            mn = nums[i]
    return ans

def maxDistance(self, s: str, k: int) -> int:
    # 3443
    return max(self._flip(s, k, 'NE'), self._flip(s, k, 'NW'),
               self._flip(s, k, 'SE'), self._flip(s, k, 'SW'))

def _flip(self, s: str, k: int, direction: str) -> int:
    res = 0
    pos = 0
    opposite = 0

    for c in s:
        if c in direction:
            pos += 1
        else:
            pos -= 1
            opposite += 1
        res = max(res, pos + 2 * min(k, opposite))

    return res

def kthSmallestProduct(nums1: list[int], nums2: list[int], k: int) -> int:
    # 2040
    A1 = [-num for num in nums1 if num < 0][::-1]
    A2 = [num for num in nums1 if num >= 0]
    B1 = [-num for num in nums2 if num < 0][::-1]
    B2 = [num for num in nums2 if num >= 0]

    negCount = len(A1) * len(B2) + len(A2) * len(B1)

    if k > negCount:  # Find the (k - negCount)-th positive.
        k -= negCount
        sign = 1
    else:
        k = negCount - k + 1  # Find the (negCount - k + 1)-th abs(negative).
        sign = -1
        B1, B2 = B2, B1

    def numProductNoGreaterThan(A: list[int], B: list[int], m: int) -> int:
        ans = 0
        j = len(B) - 1
        for i in range(len(A)):
            # For each A[i], find the first index j s.t. A[i] * B[j] <= m
            # So numProductNoGreaterThan m for this row will be j + 1
            while j >= 0 and A[i] * B[j] > m:
                j -= 1
            ans += j + 1
        return ans

    l = 0
    r = 10**10

    while l < r:
        m = (l + r) // 2
        if (numProductNoGreaterThan(A1, B1, m) +
                numProductNoGreaterThan(A2, B2, m) >= k):
            r = m
        else:
            l = m + 1

    return sign * l

def longestSubsequenceRepeatedK(s: str, k: int) -> str:
    # 2014
    def is_subsequence(subseq: str, s: str, k: int) -> bool:
        i = 0
        for c in s:
            if c == subseq[i]:
                i += 1
                if i == len(subseq):
                    k -= 1
                    if k == 0:
                        return True
                    i = 0
        return False

    count = [0] * 26
    possible_chars = []
    bfs_queue = collections.deque([""])

    for c in s:
        count[ord(c) - ord('a')] += 1

    for c in range(26):
        if count[c] >= k:
            possible_chars.append(chr(ord('a') + c))

    ans = ""
    while bfs_queue:
        curr_subseq = bfs_queue.popleft()
        if len(curr_subseq) * k > len(s):
            return ans
        for c in possible_chars:
            new_subseq = curr_subseq + c
            if is_subsequence(new_subseq, s, k):
                bfs_queue.append(new_subseq)
                ans = new_subseq

    return ans

assert kthSmallestProduct([2,5], [3,4], 2) == 8
# assert maxDifference("1122211", 3) == 1
# assert maxDifference("12233", 4) == -1
# assert maxDifference("000112233444", 5) == 2

# assert (
#     maxCandies(
#         [1, 0, 1, 0], [7, 5, 4, 100], [[], [], [1], []], [[1, 2], [3], [], []], [0]
#     )
#     == 16
# )
