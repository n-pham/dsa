import bisect  # noqa: F401
import functools  # noqa: F401


def addSpaces2109(s: str, spaces: list[int]) -> str:
    positions = [
        {"end": space_position, "start": spaces[index - 1] if index else 0}
        for index, space_position in enumerate(spaces)
    ] + [{"end": len(s), "start": spaces[-1]}]
    return " ".join([s[position["start"] : position["end"]] for position in positions])


# print(addSpaces2109("EnjoyYourCoffee", [5, 9]))


def minEnd3133(n: int, x: int) -> int:
    """
    | n1 |   n2  |
    +----+-------+
    ...00|10̲1110̲1| OK <-- x
    ...00|1011110|
    ...00|10̲1111̲1| OK
    ...00|11̲1110̲1| OK
    ...00|11̲1111̲1| OK
    ...01|0000000|
    ...01|0000001|
    ..............
    ...01|1011100|
    ...01|10̲1110̲1| OK
    ...01|1011110|
    ...01|10̲1111̲1| OK
    ...01|11̲1110̲1| OK
    ...01|11̲1111̲1| OK
    """
    zero_count = bin(x).count("0") - 1
    if zero_count == 0:
        return ((n - 1) << x.bit_length()) | x
    n1 = (n - 1) // (2**zero_count)
    n2 = (n - 1) % (2**zero_count)
    list_2 = list(bin(x)[2:])
    n2_reversed_number = reversed(bin(n2)[2:])
    for i in range(len(list_2) - 1, 0, -1):
        try:
            if list_2[i] == "0":
                list_2[i] = next(n2_reversed_number)
        except StopIteration:
            break
    # return int(f"{bin(n1)[2:]}{''.join(list_2)}",2)
    return f"{bin(n1)[2:]}{''.join(list_2)}"


# print(minEnd3133(6715154, 7193485))
# print(minEnd3133(4, 1))
# print(minEnd3133(3, 4))
# print(minEnd3133(2, 2))


def largestCombination2275(candidates: list[int]) -> int:
    """
      10000
      10001
    1000111
     111110
       1100
      11000
       1110
    -------
    1144432
    """
    largest = 0
    for i in range(max(candidates).bit_length()):
        largest = max(largest, sum([1 for number in candidates if number & (1 << i)]))
        # largest = max(largest, len(list(filter(lambda number: number & (1 << i), candidates))))
        # largest = max(largest, functools.reduce(lambda cnt, number: cnt +1 if number & (1 << i) else cnt, candidates, 0))
    return largest


# print(largestCombination2275([16, 17, 71, 62, 12, 24, 14]))
# print(largestCombination2275([8, 8]))


def canChange2337_1(start: str, target: str) -> bool:
    i, j = 0, 0
    while i < len(start) and j < len(start):
        while i < len(start) and start[i] == "_":
            i += 1
        while j < len(target) and target[j] == "_":
            j += 1
        if i == len(start) and j == len(target):
            return True
        if i == len(start) or j == len(target) or start[i] != target[j]:
            return False
        i += 1
        j += 1
    return True

def canChange2337(start: str, target: str) -> bool:
    pass

# assert canChange2337("_L__R__R_", "L______RR") is True
# assert canChange2337("_R", "_R") is False

def maxTwoEvents_greedy_fail(events: list[list[int]]) -> int:
    # sorted_events = sorted(events, key=lambda e: -e[2])
    sorted_events = []
    count_by_value = {}
    for event in events:
        bisect.insort(sorted_events, event, key=lambda e: e[2])
        count_by_value[event[2]] = count_by_value.get(event[2], 0) + 1
    print(sorted_events, count_by_value)
    for i in range(len(sorted_events)-1, -1, -1):
        previous_single_value = sorted_events[i+1][2] if i < len(sorted_events)-1 else 0
        for j in range(i-1, -1, -1):
            print(previous_single_value, sorted_events[i], sorted_events[j])
            if previous_single_value >= sorted_events[i][2] + sorted_events[j][2]:
                return previous_single_value
            if sorted_events[i][0] > sorted_events[j][1] or sorted_events[i][1] < sorted_events[j][0]:
                return sorted_events[i][2] + sorted_events[j][2]
    return sorted_events[len(sorted_events)-1][2]

def maxTwoEvents_time(events: list[list[int]]) -> int:
    print(events)
    sorted_ends = []
    sorted_end_values = []
    sorted_starts = []
    for i in range(len(events)):
        bisect.insort(sorted_ends, events[i], key=lambda e: e[1])
        bisect.insort(sorted_starts, events[i], key=lambda e: e[0])
    print(sorted_starts, sorted_ends)
    max_value = 0
    for i in range(len(events)):
        # find events where end < start
        end_index = bisect.bisect_right(sorted_ends, events[i][0]-1, key=lambda e: e[1])
        before_events = sorted_ends[:end_index]
        # print("start", events[i][0], before_events)
        max_1 = events[i][2] + (max(event[2] for event in before_events) if before_events else 0)
        # find events where start > end
        start_index = bisect.bisect_left(sorted_starts, events[i][1]+1, key=lambda e: e[0])
        after_events = sorted_starts[start_index:]
        # print("end", events[i][1], after_events)
        max_2 = events[i][2] + (max(event[2] for event in after_events) if after_events else 0)
        # print(max_value, max_1, max_2)
        max_value = max(max_value, max_1, max_2)
    return max_value

@functools.cache
def cached_max(events: tuple) -> int:
    return max([event[2] for event in events])

def maxTwoEvents_memory(events: list[list[int]]) -> int:
    print(events)
    sorted_ends = []
    sorted_end_values = []
    sorted_starts = []
    for i in range(len(events)):
        bisect.insort(sorted_ends, tuple(events[i]), key=lambda e: e[1])
        bisect.insort(sorted_starts, tuple(events[i]), key=lambda e: e[0])
    sorted_starts = tuple(sorted_starts)
    sorted_ends = tuple(sorted_ends)
    print(f"{sorted_starts=}, {sorted_ends=}")
    max_value = 0
    for i in range(len(events)):
        # find events where end < start
        end_index = bisect.bisect_right(sorted_ends, events[i][0]-1, key=lambda e: e[1])
        before_events = sorted_ends[:end_index]
        # print("start", events[i][0], before_events)
        # max_1 = events[i][2] + (max(event[2] for event in before_events) if before_events else 0)
        max_1 = events[i][2] + (cached_max(before_events) if before_events else 0)
        # find events where start > end
        start_index = bisect.bisect_left(sorted_starts, events[i][1]+1, key=lambda e: e[0])
        after_events = sorted_starts[start_index:]
        # print("end", events[i][1], after_events)
        max_2 = events[i][2] + (cached_max(after_events) if after_events else 0)
        # print(max_value, max_1, max_2)
        max_value = max(max_value, max_1, max_2)
    return max_value

class MaxTwoEvents:
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
        # print(events)
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

def findScore2593(nums: list[int]) -> int:
    nums = sorted([(v,i) for i,v in enumerate(nums)])
    print(nums)
    total = 0
    marks = set()
    for n in nums:
        if n[1] in marks:
            continue
        total += n[0]
        marks.add(n[1]-1)
        marks.add(n[1]+1)
    return total

assert findScore2593([2,2,1,3,1,5,2]) == 6
# test5 = MaxTwoEvents()
# assert test5.maxTwoEvents([[74152,91909,110],[63319,68196,6629],[17475,97226,9029],[24193,97786,1334],[60034,97031,3102],[26142,68823,8761],[78252,80542,6180],[71848,90915,2439],[20674,78352,7658],[63076,95726,3141],[62982,96724,7889],[98082,99947,43],[31163,47662,9460],[70569,86690,1035],[43569,89641,7558],[65944,93785,4173],[50013,50083,4450],[62002,75752,7006],[97425,98445,285],[72146,87680,1385],[52925,89578,5929],[10445,68400,873],[35307,63039,9828],[86514,92613,5192],[7454,56328,5970],[25498,65159,8321],[78801,93100,1571],[14317,55241,8917],[88634,94849,8492],[5177,24045,8173],[14310,32864,3246],[78570,87242,2104],[49745,81203,7630],[62436,89965,4737],[61550,64865,7674],[57890,94634,5209],[79252,81484,7675],[39874,99033,1204],[68191,90317,3413],[95493,98535,5204],[52920,87027,9386],[8707,13035,875],[28750,83816,2008],[27172,43423,6305],[70951,94311,1062],[2013,82730,1816],[38132,50489,4106],[80314,86023,9044],[41330,73377,8801],[90085,92868,2565],[57099,74996,5146],[12525,37798,8817],[78952,99468,3657],[5550,8366,2739],[51403,86622,4999],[35468,57148,6629],[6244,40399,9101],[92871,99835,730],[72881,85666,3158],[21596,80389,360],[9228,90520,7013],[1840,59008,4642],[31319,60953,690],[43946,93749,2225],[94446,97474,6225],[35670,84186,2150],[62019,62507,3064],[42328,49012,2083],[64411,99385,661],[66672,87953,5536],[14956,23059,1665],[91518,93667,1166],[51681,59765,9468],[12151,75034,1312],[51104,87056,2493],[73525,75050,2687],[96494,98592,2302],[84977,94379,9381],[64710,79934,4163],[49724,74504,293],[59186,90677,4585],[32139,67408,424],[34920,76783,9111],[72420,86433,2480],[51804,67291,1863],[42434,58222,1555],[14793,16890,599],[71284,81336,3381],[8005,89372,9982],[36626,58499,4899],[85225,93811,3572],[41891,99423,3675],[33912,40452,4727],[87461,97836,7903],[79543,96407,7745],[40311,70140,7404],[82944,89459,3638],[12683,99729,4973],[13726,80014,1498],[53377,56550,2382],[35455,36773,8566],[30471,57452,5390],[20125,31151,4485],[87771,90371,294],[49323,62221,7892],[72015,93870,8304],[76621,92690,1184],[1727,9743,1600],[84312,90410,4283],[43880,56359,291],[28570,61471,2253],[52423,81021,4023],[53556,62537,8187],[17397,99559,3184],[36089,94699,4784],[85328,97303,120],[45511,61427,2797],[4316,4558,2043],[85206,94328,252],[5080,6369,3490],[36273,88727,2651],[93585,97123,7835],[15137,30578,1362],[3166,83735,436],[54282,65088,9542],[35392,73512,1437],[82835,83468,4241],[83029,98402,8102],[22870,85083,7091],[14460,78771,3743],[28203,55328,4743],[46847,86754,5738],[4251,63054,2723],[33448,56737,1553],[73232,91529,1569],[32522,78589,155],[56304,74441,2632],[42353,69439,8123],[62210,95420,7901],[47877,96636,9516],[62817,96445,3903],[57007,68223,812],[94839,94887,6343],[67305,89684,6902],[90505,97351,9496],[64736,83562,5332],[93246,94217,798],[92158,96770,9485],[7175,28464,4072],[10089,79270,2632],[9607,90527,63],[10261,57879,4340],[95869,98664,1535],[95697,97847,3659],[65145,92460,4018],[30869,48893,2673],[24819,28014,9213],[68291,69882,3567],[36927,58583,3172],[2305,58949,1719],[4061,90615,716],[3126,12120,3156],[429,16955,4606],[14403,29189,218],[92883,97975,1509],[25046,64308,3915],[11778,16532,6168],[81903,94464,4455],[72619,78857,6866],[90479,99458,8462],[87802,89586,5848],[9731,68836,5417],[67589,92849,3994],[45779,46913,2116],[81219,82119,1079],[25930,71362,2751],[71046,87067,3168],[60983,92601,335],[62728,78893,4198],[27366,78463,3254],[45421,94174,2528],[77203,90969,4500],[36622,90036,6004],[71493,94374,3345],[21815,91262,8476],[39023,90389,868],[41239,49833,7617],[8022,17893,1103],[74013,99363,7928],[83984,92681,9754],[21697,81079,119],[90957,94475,75],[38187,77628,8773],[74394,94803,3995],[19345,79434,138],[92301,97442,3792],[92000,96049,2245],[17750,33970,7608],[83758,88095,7386],[44228,96587,9295],[43629,51313,7436],[66914,71930,5304],[21909,43615,1703],[47464,53157,9259],[73829,80931,5698],[75880,79295,2305],[11327,87239,5792],[62066,67496,6198],[49724,67230,8135],[14817,26585,2075],[58510,87112,1815],[95168,98476,1444],[99563,99688,5863],[89899,96336,7904],[15483,25656,2712],[19724,77837,4153],[41474,90911,8671]]) == 19582
# test,test2,test3,test4 = MaxTwoEvents(),MaxTwoEvents(),MaxTwoEvents(),MaxTwoEvents()
# assert test.maxTwoEvents([[1,3,2],[4,5,2],[2,4,3]]) == 4
# assert test2.maxTwoEvents([[72, 80, 70], [35, 90, 47]]) == 70
# assert test3.maxTwoEvents([[1,3,2],[4,5,2],[1,5,5]]) == 5
# assert test4.maxTwoEvents([[1,5,3],[1,5,1],[6,6,5]]) == 8
