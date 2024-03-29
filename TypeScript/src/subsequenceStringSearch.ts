/*
We say that a string contains the word hackerrank if a subsequence
of its characters spell the word hackerrank. Remember that a subsequence
maintains the order of characters selected from a sequence.

For each query, print YES on a new line if the string contains hackerrank,
otherwise, print NO.

Function Description
- Complete the hackerrankInString function in the editor below.

Parameter:
- string s: a string

Returns:
- string: YES or NO

Sample Input:

hereiamstackerrank 	// should return YES
hackerworld 		// should return NO
hhaacckkekraraannk 	// should return YES
*/

export function subsequenceStringSearch(s: string): string {
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