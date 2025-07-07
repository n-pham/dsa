import collections
import heapq
import sys

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
