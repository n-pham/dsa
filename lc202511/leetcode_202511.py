import bisect
from collections import defaultdict
import orjson


def find_json_depth(obj, current_depth=1) -> int:
    if isinstance(obj, dict):
        if not obj:
            return current_depth
        return max(find_json_depth(v, current_depth + 1) for v in obj.values())
    elif isinstance(obj, list):
        if not obj:
            return current_depth
        return max(find_json_depth(item, current_depth + 1) for item in obj)
    else:
        return current_depth


def test_find_json_depth():
    obj = orjson.loads('{"name": "Mixed nested","data": {"level1": {"level2": [{"level3": {"level4": {"level5": "deep"}}}]}}}')
    result = find_json_depth(obj)
    assert result == 8

def findXSum(nums: list[int], k: int, x: int) -> list[int]:
    # 3318 3321
    ans, top_set, other_set = [], [], []
    window_sum_container = {"value": 0}
    count = defaultdict(int)

    def update(num: int, count_delta: int):
        old_count = count[num]

        # Part 1: Remove old entry if it exists
        if old_count > 0:
            old_priority_tuple = (old_count, num)
            i = bisect.bisect_left(top_set, old_priority_tuple)
            if i < len(top_set) and top_set[i] == old_priority_tuple:
                window_sum_container["value"] -= old_count * num
                top_set.pop(i)
            else:
                i = bisect.bisect_left(other_set, old_priority_tuple)
                if i < len(other_set) and other_set[i] == old_priority_tuple:
                    other_set.pop(i)

        # Part 2: Update count
        new_count = old_count + count_delta
        count[num] = new_count

        # Part 3: Add new entry if count > 0
        if new_count > 0:
            new_priority_tuple = (new_count, num)
            if top_set and new_priority_tuple > top_set[0]:
                window_sum_container["value"] += new_count * num
                bisect.insort(top_set, new_priority_tuple)
            else:
                bisect.insort(other_set, new_priority_tuple)

    for i, num in enumerate(nums):
        update(num, 1)
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
        update(leftmost_element, -1)
    return ans


def test_findXSum():
    assert findXSum([3, 8, 7, 8, 7, 5], 2, 2) == [11, 15, 15, 15, 12]
    assert findXSum([1, 1, 2, 2, 3, 4, 2, 3], 6, 2) == [6, 10, 12]


if __name__ == "__main__":
    # test_findXSum()
    test_find_json_depth()
