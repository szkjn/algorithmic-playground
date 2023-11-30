/*
Create a program that constructs a linked list. The program should 
define a MyNode class representing list elements, and a Solution class to 
manage the list. Implement an insert method in Solution to add nodes to 
the list's end, and a display method to print the list. The user inputs 
the number of elements (T) and each element's value, which are added 
sequentially to the list.
*/
import * as readline from 'readline';

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

class MyNode {
    data: number;
    next: MyNode | null;

    constructor(data: number) {
        this.data = data;
        this.next = null;
    }
}

class Solution {
    display(head: MyNode | null): void {
        let current: MyNode | null = head;
        while (current !== null) {
            process.stdout.write(current.data + ' ');
            current = current.next;
        }
        console.log();
    }

    insert(head: MyNode | null, data: number): MyNode {
        const newNode = new MyNode(data);
        if (head === null) {
            head = newNode;
        } else {
            let current: MyNode = head;
            while (current.next !== null) {
                current = current.next;
            }
            current.next = newNode;
        }
        return head;
    }
}

const mylist = new Solution();
let head: MyNode | null = null;

rl.question("Enter number of elements: ", (T) => {
    const numberOfElements = parseInt(T);

    const readElement = (i: number) => {
        if (i < numberOfElements) {
            rl.question(`Enter data for element ${i + 1}: `, (data) => {
                head = mylist.insert(head, parseInt(data));
                readElement(i + 1);
            });
        } else {
            mylist.display(head);
            rl.close();
        }
    };

    readElement(0);
});
