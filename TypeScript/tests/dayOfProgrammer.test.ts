import { readFileSync } from 'fs';
import { join } from 'path';
import { dayOfProgrammer } from '../src/dayOfProgrammer';


interface TestCase {
  year: number;
  expected: string;
}


const testDataPath = join(__dirname, '../../test_data/day_of_programmer.json');
const testCases: TestCase[] = JSON.parse(readFileSync(testDataPath, 'utf-8'));


describe('dayOfProgrammer', () => {
  testCases.forEach(testCase => {
    test(`find day of programmer for year ${testCase.year}`, () => {
      expect(dayOfProgrammer(testCase.year)).toEqual(testCase.expected);
    });
  });
});