from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        result_list = ListNode(0)
        temp = result_list
        carry = 0
        while l1 is not None or l2 is not None:
            val1 = l1.val if l1 else 0
            val2 = l2.val if l2 else 0

            val = val1 + val2 + carry
            carry = val // 10

            temp.next = ListNode(val % 10)
            temp = temp.next

            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None

        if carry:
            temp.next = ListNode(carry)

        return result_list.next



