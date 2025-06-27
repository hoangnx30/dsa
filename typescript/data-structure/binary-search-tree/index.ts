export class TreeNode<T = number> {
  data: T;
  left: TreeNode<T> | null;
  right: TreeNode<T> | null;

  constructor(data: T) {
    this.data = data;
    this.left = null;
    this.right = null;
  }
}

export class BinartSearchTree<T = number> {
  private root: TreeNode<T>;

  constructor(root: TreeNode<T>) {
    this.root = root;
  }

  insert(data: T) {
    this._insert(data, this.root);
    return this;
  }

  private _insert(data: T, node: TreeNode<T>) {
    const newNode = new TreeNode<T>(data);
    if (this.compare(data, node)) {
      if (node.right === null) {
        node.right = newNode;
      } else {
        this._insert(data, node.right);
      }
    } else {
      if (node.left === null) {
        node.left = newNode;
      } else {
        this._insert(data, node.left);
      }
    }
  }

  private compare(data: T, node: TreeNode<T>) {
    return data > node.data;
  }

  printInOrder() {
    this._printInOrder(this.root);
  }

  printPreOrder() {
    this._printPreOrder(this.root);
  }

  printPostOrder() {
    const arr = [];
    this._printPostOrder(this.root, arr);
    console.log(arr);
  }

  private _printInOrder(node: TreeNode<T> | null) {
    if (node !== null) {
      this._printInOrder(node.left);
      console.log(node.data);
      this._printInOrder(node.right);
    }
  }

  private _printPreOrder(node: TreeNode<T> | null) {
    if (node !== null) {
      console.log(node.data);
      this._printPreOrder(node.left);
      this._printPreOrder(node.right);
    }
  }

  private _printPostOrder(node: TreeNode<T> | null, arr: T[]) {
    if (node !== null) {
      this._printPostOrder(node.left, arr);
      this._printPostOrder(node.right, arr);
      arr.push(node.data);
    }
  }

  prettyPrint(node = this.root, prefix = "", isLeft = true) {
    if (node === null) {
      return;
    }
    if (node.right !== null) {
      this.prettyPrint(
        node.right,
        `${prefix}${isLeft ? "│   " : "    "}`,
        false
      );
    }
    console.log(`${prefix}${isLeft ? "└── " : "┌── "}${node.data}`);
    if (node.left !== null) {
      this.prettyPrint(node.left, `${prefix}${isLeft ? "    " : "│   "}`, true);
    }
  }

  bfs() {
    const queue: TreeNode<T>[] = [];
    queue.push(this.root);

    const result: T[] = [];

    while (queue.length > 0) {
      const node = queue.shift();

      result.push(node!.data);

      if (node?.left) {
        queue.push(node.left);
      }

      if (node?.right) {
        queue.push(node.right);
      }
    }

    console.log(result);
  }
}

const bst = new BinartSearchTree(new TreeNode(10));
bst.insert(6).insert(13).insert(5).insert(7);

// bst.printInOrder();
// bst.printPreOrder();
// bst.printPostOrder();

bst.prettyPrint();
bst.bfs();
