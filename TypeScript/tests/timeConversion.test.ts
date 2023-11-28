import { timeConversion } from '../src/timeConversion'

describe('timeConversion', () => {
  test('returns time in 24-hour format ', () => {
    expect(timeConversion("12:01:00PM")).toEqual("12:01:00");
    expect(timeConversion("12:01:00AM")).toEqual("00:01:00");
    expect(timeConversion("07:05:45PM")).toEqual("19:05:45");
  })
})