import bisect
import functools


class Solution:
    @functools.cache
    def maxInBefore(self, index: int) -> int:
        if index == 0:
            return 0
        return max(self.sorted_ends[index-1][2], self.maxInBefore(index-2) if index > 1 else 0)

    @functools.cache
    def maxInAfter(self, index: int) -> int:
        if index >= len(self.sorted_starts):
            return 0
        return max(self.sorted_starts[index][2], self.maxInAfter(index+1) if index < len(self.sorted_starts)-1 else 0)

    def maxTwoEvents(self, events: list[list[int]]) -> int:
        # 2054
        self.sorted_ends = []
        self.sorted_starts = []
        for i in range(len(events)):
            bisect.insort(self.sorted_ends, events[i], key=lambda e: e[1])
            bisect.insort(self.sorted_starts, events[i], key=lambda e: e[0])
        # print(f"{self.sorted_starts=}, {self.sorted_ends=}")
        max_value = 0
        for i in range(len(events)):
            # find events where end < start
            end_index = bisect.bisect_right(self.sorted_ends, events[i][0]-1, key=lambda e: e[1])
            # print(events[i], "start", events[i][0])
            # print(f"{end_index=} {self.maxInBefore(end_index)=}")
            max_1 = events[i][2] + self.maxInBefore(end_index)
            # find events where start > end
            start_index = bisect.bisect_left(self.sorted_starts, events[i][1]+1, key=lambda e: e[0])
            # print(events[i], "end", events[i][1])
            # print(f"{start_index=} {self.maxInAfter(start_index)=}")
            max_2 = events[i][2] + self.maxInAfter(start_index)
            # print(max_value, max_1, max_2)
            max_value = max(max_value, max_1, max_2)
        return max_value
