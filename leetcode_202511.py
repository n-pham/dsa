import bisect
from collections import defaultdict


def findXSum(nums: list[int], k: int, x: int) -> list[int]:
    # 3318 3321
    ans, top_set, other_set = [], [], []
    window_sum_container = {"value": 0}
    count = defaultdict(int)

    def add(num: int) -> None:
        if count[num] == 0:
            return
        priority_tuple = (count[num], num)
        # If top set is empty or new element has higher priority than minimum in top set
        if top_set and priority_tuple > top_set[0]:
            window_sum_container["value"] += count[num] * num
            bisect.insort(top_set, priority_tuple)
        else:
            bisect.insort(other_set, priority_tuple)

    def remove(num: int) -> None:
        if count[num] == 0:
            return
        priority_tuple = (count[num], num)

        i = bisect.bisect_left(top_set, priority_tuple)
        if i < len(top_set) and top_set[i] == priority_tuple:
            window_sum_container["value"] -= count[num] * num
            top_set.pop(i)
        else:
            i = bisect.bisect_left(other_set, priority_tuple)
            if i < len(other_set) and other_set[i] == priority_tuple:
                other_set.pop(i)

    for i, num in enumerate(nums):
        remove(num)
        count[num] += 1
        add(num)
        window_start = i - k + 1
        if window_start < 0:
            continue
        while other_set and len(top_set) < x:
            element_to_promote = other_set.pop()
            bisect.insort(top_set, element_to_promote)
            window_sum_container["value"] += (
                element_to_promote[0] * element_to_promote[1]
            )
        while len(top_set) > x:
            element_to_demote = top_set.pop(0)
            window_sum_container["value"] -= element_to_demote[0] * element_to_demote[1]
            bisect.insort(other_set, element_to_demote)
        ans.append(window_sum_container["value"])
        # Remove the leftmost element from window for next iteration
        leftmost_element = nums[window_start]
        remove(leftmost_element)
        count[leftmost_element] -= 1
        if count[leftmost_element] > 0:
            add(leftmost_element)
    return ans


def test_findXSum():
    assert findXSum([3, 8, 7, 8, 7, 5], 2, 2) == [11, 15, 15, 15, 12]
    assert findXSum([1, 1, 2, 2, 3, 4, 2, 3], 6, 2) == [6, 10, 12]


if __name__ == "__main__":
    test_findXSum()
