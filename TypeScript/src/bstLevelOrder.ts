class MyNode {
    public left: MyNode | null;
    public right: MyNode | null;
    public data: number;

    constructor(data: number) {
        this.left = null;
        this.right = null;
        this.data = data;
    }
}

class Solution {
    insert(root: MyNode | null, data: number): MyNode {
        if (root === null) {
            return new MyNode(data);
        } else {
            if (data <= root.data) {
                root.left = this.insert(root.left, data);
            } else {
                root.right = this.insert(root.right, data);
            }
        }
        return root;
    }

    levelOrder(root: MyNode | null): void {
        const queue: MyNode[] = [];
        if (root !== null) {
            queue.push(root);
        }

        let output: string = "";
        while (queue.length > 0) {
            const node: MyNode = queue.shift()!;
            output += `${node.data} `;

            if (node.left !== null) {
                queue.push(node.left);
            }
            if (node.right !== null) {
                queue.push(node.right);
            }
        }
        console.log(output.trim());
    }
}

// Example usage with hardcoded inputs
let myTree: Solution = new Solution();
let root: MyNode | null = null;
let data: number[] = [3, 5, 4, 7, 2, 1]; // Example data

data.forEach(d => {
    root = myTree.insert(root, d);
});

myTree.levelOrder(root);
