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


assert (
    maxCandies(
        [1, 0, 1, 0], [7, 5, 4, 100], [[], [], [1], []], [[1, 2], [3], [], []], [0]
    )
    == 16
)
