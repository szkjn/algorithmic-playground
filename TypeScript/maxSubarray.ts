// Interview question of the week by Cassidy Williams (issue #281) :
// 
// Given an array of integers arr and an integer n, return a subarray of 
// arr of length n where the sum is the largest. Make sure you maintain 
// the order of the original array, and if n is greater than arr.length, 
// you can choose what you want to return.

const sum = (arr: number[]): number => {
    const sum: number = arr.reduce((a, b) => a + b, 0)
    return sum
}

const maxSubarray = (arr: number[], n: number): void => {
    let subarray: number[] = []
    for (let i=0; i < arr.length-n+1; i++) {
        const currentArr: number[] = arr.slice(i,i+n)    
        if (sum(currentArr) > sum(subarray)) {
            subarray = currentArr
        }
    }
    console.log(subarray)
}

maxSubarray([-4, 2, -5, 1, 2, 3, 6, -5, 1], 4)
maxSubarray([1, 2, 0, 5], 2)