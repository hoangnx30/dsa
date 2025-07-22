import java.util.ArrayList;
import java.util.List;

public class BinaryTreeInOrderTraversal {
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

    class Solution {
        public List<Integer> inorderTraversal(TreeNode root) {
            List<Integer> list = new ArrayList<>();
            this.inorderTraversal(root, list);
            return list;
        }

        private List<Integer> inorderTraversal(TreeNode root, List<Integer> result) {
            if (root != null) {
                result = inorderTraversal(root.left, result);
                result.add(root.val);
                result = inorderTraversal(root.right, result);

                return result;
            }

            return result;
        }
    }
}
