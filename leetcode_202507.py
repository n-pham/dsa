import collections
import heapq
import sys
import bisect
import functools

def possibleStringCount(word: str, k: int) -> int:
    # 3333
    mod = 10**9 + 7
    cnt = []
    total = 1
    i = 0
    while i < len(word):
        j = i
        while i < len(word) and word[i] == word[j]:
            i += 1
        if i > j + 1:
            cnt.append(i - j - 1)
            total = total * (i - j) % mod
        k -= 1
    if k <= 0:
        return total

    dp = [0] * k
    dp[0] = 1
    for c in cnt:
        for i in range(1, k):
            dp[i] = (dp[i] + dp[i - 1]) % mod
        for i in range(k - 1, c, -1):
            dp[i] = (dp[i] - dp[i - c - 1]) % mod
    for i in range(1, k):
        dp[i] = (dp[i] + dp[i - 1]) % mod
    return (total - dp[k - 1]) % mod

def maxEvents(events: list[list[int]]) -> int:
    # 1353
    # Create a default dictionary to hold events keyed by start date
    event_dict = collections.defaultdict(list)
    
    # Initialize variables to track the earliest and latest event dates
    earliest_start, latest_end = sys.maxsize, 0
    
    # Populate event_dict with events and update earliest_start and latest_end
    for start, end in events:
        event_dict[start].append(end)
        earliest_start = min(earliest_start, start)
        latest_end = max(latest_end, end)
    
    # Initialize an empty min-heap to store active events' end dates
    min_heap = []
    
    # Counter for the maximum number of events one can attend
    max_events_attended = 0
    
    # Iterate over each day within the range of event dates
    for day in range(earliest_start, latest_end + 1):
        # Remove events that have already ended
        while min_heap and min_heap[0] < day:
            heapq.heappop(min_heap)
        
        # Push all end dates of events starting today onto the heap
        for end in event_dict[day]:
            heapq.heappush(min_heap, end)
        
        # If there are any events available to attend today, attend one and increment count
        if min_heap:
            max_events_attended += 1
            heapq.heappop(min_heap)  # Remove the event that was attended
    
    # Return the total number of events attended
    return max_events_attended

def maxValue(self, events: list[list[int]], k: int) -> int:
    # 1751
    n = len(events)
    events.sort()

    starts = [e[0] for e in events]

    @functools.lru_cache(None)
    def dp(i, k_left):
        if k_left == 0 or i >= n:
            return 0

        # Option 1: Skip event i.
        # This is the value if we don't attend event i.
        res = dp(i + 1, k_left)

        # Option 2: Attend event i.
        # Find the index of the first event that starts after event i ends.
        end_day = events[i][1]
        next_i = bisect.bisect_right(starts, end_day)
        
        # The value is the current event's value + max value from non-overlapping
        # future events, with one less event allowed to be attended.
        res = max(res, events[i][2] + dp(next_i, k_left - 1))
        
        return res

    return dp(0, k)

def maxFreeTime(eventTime: int, k: int, startTime: list[int], endTime: list[int]) -> int:
    gaps = ([startTime[0]] +
            [startTime[i] - endTime[i - 1] for i in range(1, len(startTime))] +
            [eventTime - endTime[-1]])
    windowSum = sum(gaps[:k + 1])
    ans = windowSum

    for i in range(k + 1, len(gaps)):
        windowSum += gaps[i] - gaps[i - k - 1]
        ans = max(ans, windowSum)

    return ans

