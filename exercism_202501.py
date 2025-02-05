class EmptyListException(Exception):
    pass


class Node:
    def __init__(self, value, next_node=None):
        self._value = value
        self.next_node = next_node

    def value(self):
        return self._value

    def next(self):
        return self.next_node

class LinkedList:
    def __init__(self, values=None):
        self.head_node = None
        if values:
            self.head_node = Node(values[0])
            for value in values[1:]:
                node = Node(value)
                node.next_node = self.head_node
                self.head_node = node

    def __iter__(self):
        return self.__next__()

    def __next__(self):
        current_node = self.head_node
        while current_node is not None:
            yield current_node.value()
            current_node = current_node.next()


    def __len__(self):
        length = 0
        current = self.head_node
        while current:
            length += 1
            current = current.next_node
        return length

    def head(self):
        if not self.head_node:
            raise EmptyListException("The list is empty.")
        return self.head_node

    def push(self, value):
        new_node = Node(value)
        new_node.next_node = self.head_node
        self.head_node = new_node

    def pop(self):
        if not self.head_node:
            raise EmptyListException("The list is empty.")
        value = self.head_node.value()
        self.head_node = self.head_node.next_node
        return value

    def reversed(self):
        reversed_list = LinkedList()
        current = self.head_node
        while current:
            reversed_list.push(current.value())
            current = current.next_node
        return reversed_list

def rectangles(strings):
    """
     0123456
    0   +--+
    1  ++  |
    2+-++--+
    3|    |
    4+--+--+
    (0,3),(0,6),(1,2),(1,3),(2,0),(2,2),(2,3),(2,6),(4,0),(4,3),(4,6)
    test_rectangles_must_have_four_sides
    """
    pass
