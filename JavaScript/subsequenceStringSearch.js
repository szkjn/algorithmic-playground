function hackerrankInString(s) {
    let hackerrank = "hackerrank";
    let idx = 0;

    for (let i = 0; i < s.length; i++) {
        if (s.charAt(i) === hackerrank.charAt(idx)) {
            idx++;
            if (idx === hackerrank.length) {
                return "YES";
            }
        }
    }

    return "NO";
}

console.log(hackerrankInString("hereiamstackerrank"));
console.log(hackerrankInString("hackerworld"));
console.log(hackerrankInString("hereiamsthhaacckkekraraannkackerrank"));