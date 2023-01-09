// Interview question of the week by Cassidy Williams (issue #282) :
// 
// Given a number, sum every second digit in that number.
// Example:
// > sumEveryOther(548915381)
// > 26 // 4+9+5+8
// > sumEveryOther(10)
// > 0 
// > sumEveryOther(1010.11)
// > 1 // 0+0+1

function sumEveryOther(n) {
  n = String(n)
  res = 0
  for (let i = 1; i < n.length; i += 2) { res += Number(n[i]) }
  console.log(res)
}

sumEveryOther(548915381)
sumEveryOther(10)
sumEveryOther(1010.11)
