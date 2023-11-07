// Interview question of the week by Cassidy Williams (issue #260) :
//
// Given a list, swap every two adjacent nodes. Something to think 
// about: How would your code change if this were a linked list, 
// versus an array?

function swapPairs(arr) {
  
    let results = []
    
    while (arr.length > 0) {
        let firstEl = arr.shift()
        let secondEl = arr.shift()
        results.push(secondEl)
        results.push(firstEl)
    } 
    console.log(results)
}

swapPairs([1, 2, 3, 4])
swapPairs([])