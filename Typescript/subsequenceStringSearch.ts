function hackerrankInStringTS(s: string): string {
    const hackerrank: string = "hackerrank";
    let idx: number = 0;

    for (let i: number = 0; i < s.length; i++) {
        if (s.charAt(i) === hackerrank.charAt(idx)) {
            idx++;
            if (idx === hackerrank.length) {
                return "YES";
            }
        }
    }

    return "NO";
}

console.log(hackerrankInStringTS("hereiamstackerrank"));
console.log(hackerrankInStringTS("hackerworld"));
console.log(hackerrankInStringTS("hereiamsthhaacckkekraraannkackerrank"));
