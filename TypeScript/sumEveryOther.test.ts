import { sumEveryOtherTS } from './sumEveryOther'

describe('sumEveryOtherTS', () => {
  test('calculates the sum of every other digit', () => {
    expect(sumEveryOtherTS(548915381)).toEqual(26);
    expect(sumEveryOtherTS(10)).toEqual(0);
    expect(sumEveryOtherTS(1010.11)).toEqual(1);
  })
})