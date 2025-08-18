from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        current = head
        previous = None

        while current:
            next = current.next

            current.next = previous
            previous = current
            current = next

        return previous

    def printList(self, head):

        current = head
        while current:
            print(current.val)
            current = current.next


if __name__ == '__main__':
    head = ListNode(1)
    head.next = ListNode(2)
    head.next.next = ListNode(3)

    solution = Solution()
    solution.printList(solution.reverseList(head))
