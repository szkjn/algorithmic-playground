import { swapPairs } from './swapPairs'

describe('swapPairs', () => {
  test('swaps pairs in an array of integers', () => {
    expect(swapPairs([1, 2, 3, 4])).toEqual([2, 1, 4, 3]);
    expect(swapPairs([1, 2, 3])).toEqual([2, 1, 3]);
    expect(swapPairs([])).toEqual([]);
  })
})