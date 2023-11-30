'use strict';

import * as readline from 'readline';

export function dayOfProgrammer(year: number): string {
    if (year === 1918) {
        // Special case for the year 1918
        return "26.09.1918";
    }

    let isLeap = false;
    if (year < 1918) {
        // Julian calendar
        isLeap = year % 4 === 0;
    } else {
        // Gregorian calendar
        isLeap = (year % 400 === 0) || (year % 4 === 0 && year % 100 !== 0);
    }

    if (isLeap) {
        return `12.09.${year}`;
    } else {
        return `13.09.${year}`;
    }
}

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question('Enter a year: ', (input) => {
    const year = parseInt(input.trim(), 10);
    const result = dayOfProgrammer(year);
    console.log(result);
    rl.close();
});
