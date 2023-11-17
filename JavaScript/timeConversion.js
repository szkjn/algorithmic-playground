/*
Given a time in 12-hour AM/PM format, convert it to 24-hour format.

- 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

Example:
- s = 12:01:00PM -> returns 12:01:00
- s = 12:01:00AM -> returns 00:01:00

Function parameter:
- string s: time in 12 hour format

Returns:
- string: time in 24 hour format

Input Format:
- A single string of time in 12-hour format (i.e.: hh:mm:ssAM or hh:mm:ssPM).

Constraints:
- All input times are valid
*/

const readline = require('readline');

function timeConversion(s) {
    let res = s.substring(0, 8);
    const hh = s.substring(0, 2);
    const meridiem = s.substring(8).trim();

    let hhInt = parseInt(hh, 10);
    if (isNaN(hhInt)) {
        throw new Error("Invalid hour format");
    }

    if (meridiem === "PM" && hhInt < 12) {
        res = (hhInt + 12).toString() + res.substring(2);
    } else if (meridiem === "AM" && hhInt === 12) {
        res = "00" + res.substring(2);
    }
    return res;
}

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question('Enter time in 12-hours format hh:mm:ss<AM/PM>:\n', (input) => {
    try {
        const result = timeConversion(input);
        console.log(result);
    } catch (error) {
        console.error('Error:', error.message);
    }
    rl.close();
});
