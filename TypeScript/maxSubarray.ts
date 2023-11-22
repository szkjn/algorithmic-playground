// Interview question of the week by Cassidy Williams (issue #281) :
// 
// Given an array of integers arr and an integer n, return a subarray of 
// arr of length n where the sum is the largest. Make sure you maintain 
// the order of the original array, and if n is greater than arr.length, 
// you can choose what you want to return.

export const sum = (arr: number[]): number => {
    const sum: number = arr.reduce((a, b) => a + b, 0)
    return sum
}

export const maxSubarray = (arr: number[], n: number): number[] => {
    let subarray: number[] = []
    for (let i=0; i < arr.length-n+1; i++) {
        const currentArr: number[] = arr.slice(i,i+n)    
        if (sum(currentArr) > sum(subarray)) {
            subarray = currentArr
        }
    }
    return subarray
}