import { subsequenceStringSearch } from '../src/subsequenceStringSearch'

describe('subsequenceStringSearch', () => {
    test('returns "YES" if subsequence is contained in "hackerrank", else "NO"', () => {
      expect(subsequenceStringSearch("hereiamstackerrank")).toEqual("YES");
      expect(subsequenceStringSearch("hackerworld")).toEqual("NO");
      expect(subsequenceStringSearch("hhaacckkekraraannk")).toEqual("YES");
    })
  })