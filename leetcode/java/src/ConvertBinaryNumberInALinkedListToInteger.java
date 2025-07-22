

public class ConvertBinaryNumberInALinkedListToInteger {
    static class ListNode {
        int val;
        ListNode next;

        ListNode() {
        }

        ListNode(int val) {
            this.val = val;
        }

        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }
    }

    static int getDecimalValue(ListNode head) {
        int result = head.val;
        while (head.next != null) {
            result = (result << 1) | head.next.val;
            head = head.next;
        }
        return result;
    }

    static int getDecimalValue1(ListNode head) {
        StringBuilder decimalNumber = new StringBuilder();
        while (head != null) {
            decimalNumber.append(head.val);
            head = head.next;
        }
        int result = 0;
        System.out.println(decimalNumber);
        for (int i = 0; i < decimalNumber.length(); i++) {

            char num = decimalNumber.charAt(decimalNumber.length() - i - 1);
            if (Character.getNumericValue(num) == 1) {
                result += Math.pow(2, i);
            }
        }

        return result;
    }

    public static void main(String[] args) {
        ListNode node1 = new ListNode(1);
        ListNode node2 = new ListNode(0);
        ListNode node3 = new ListNode(1);

        node1.next = node2;
        node2.next = node3;

        int result = getDecimalValue1(node1);

        System.out.printf("result: %d", result);

    }
}
