// Interview question of the week by Cassidy Williams (issue #260) :
//
// Given a list, swap every two adjacent nodes. Something to think 
// about: How would your code change if this were a linked list, 
// versus an array?

function swapPairs(arr: number[]): number[] {
    for (let i = 0; i < arr.length - 1; i += 2) {
        [arr[i], arr[i + 1]] = [arr[i + 1], arr[i]];
    }
    return arr;
}

console.log(swapPairs([1, 2, 3, 4])); // Outputs: [2, 1, 4, 3]
console.log(swapPairs([1, 2, 3])); // Outputs: [2, 1, 3]
console.log(swapPairs([])); // Outputs: []
