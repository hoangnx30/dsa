import { LinkedListNode } from "./linked-list-node";

export class SingleLinkedList<T> {
  private head: LinkedListNode<T> | null = null;
  private tail: LinkedListNode<T> | null = null;
  constructor() {}

  prepend(value: T) {
    const newNode = new LinkedListNode(value);

    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;

      return this;
    }

    newNode.next = this.head;
    this.head = newNode;

    return this;
  }

  append(value: T) {
    const newNode = new LinkedListNode(value);

    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;
      return this;
    }

    this.tail!.next = newNode;
    this.tail = newNode;

    return this;
  }

  getHeadValue(): T | null {
    return this.head?.value ?? null;
  }

  deleteHead() {
    if (!this.head) return this;

    if (this.head.next) {
      this.head = this.head.next;
    } else {
      this.head = null;
      this.tail = null;
    }

    return this;
  }

  deleteTail() {
    if (!this.tail) return this;

    if (!this.head?.next) {
      this.head = null;
      this.tail = null;

      return this;
    }

    let currNode = this.head;

    while (currNode) {
      if (currNode?.next && !currNode?.next?.next) {
        this.tail = currNode;
        currNode.next = null;
      } else {
        currNode = currNode.next!;
      }
    }

    return this;
  }

  toArray() {
    let currNode = this.head;
    const result: T[] = [];
    while (currNode) {
      result.push(currNode.value);

      currNode = currNode.next;
    }

    return result;
  }
}
