import java.util.Stack;

class Solution1 {

    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode() {
        }

        TreeNode(int val) {
            this.val = val;
        }

        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    public String simplifyPath(String path) {
        String[] components = path.split("/");
        Stack<String> stack = new Stack<>();
        for (String comp : components) {
            if (comp.isEmpty() || comp.equals(".")) {
                continue;
            }

            if (comp.equals("..")) {
                if (!stack.isEmpty()) {
                    stack.pop();
                }
            } else {
                stack.push(comp);
            }
        }

        StringBuilder sb = new StringBuilder();

        while (!stack.isEmpty()) {
            sb.insert(0, "/" + stack.pop());
        }
        return sb.isEmpty() ? "/" : sb.toString();
    }
}

public class SimplifyPath {
    public static void main(String[] args) {
        Solution solution = new Solution();

        String path = "/home/";
        String result = solution.simplifyPath(path);

        System.out.printf("Input: %s, Value: %s", path, result);
    }
}
