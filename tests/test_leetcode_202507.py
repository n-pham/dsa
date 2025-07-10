
import unittest
from leetcode_202507 import maxFreeTimeII

class TestMaxFreeTimeII(unittest.TestCase):
    def test_simple_case(self):
        eventTime = 100
        startTime = [10, 60]
        endTime = [20, 70]
        # Gaps are [10, 40, 30]. Initial max free time is 40.
        # Durations are [10, 10].
        # Move meeting 1 ([10, 20]):
        #   Remove it -> new gap is 10 + 10 + 40 = 60. Other gap is 30.
        #   Place it in the 30 gap -> max free time is 60.
        # Move meeting 2 ([60, 70]):
        #   Remove it -> new gap is 40 + 10 + 30 = 80. Other gap is 10.
        #   Place it in the 10 gap -> max free time is 80.
        self.assertEqual(maxFreeTimeII(eventTime, startTime, endTime), 80)

    def test_no_meetings(self):
        self.assertEqual(maxFreeTimeII(100, [], []), 100)

    def test_one_meeting(self):
        self.assertEqual(maxFreeTimeII(100, [10], [20]), 90)

if __name__ == '__main__':
    unittest.main()
