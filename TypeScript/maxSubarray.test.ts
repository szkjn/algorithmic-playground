import { sum, maxSubarray } from './maxSubarray';

describe('sum', () => {
  test('calculates the sum of array elements', () => {
    expect(sum([1, 2, 3])).toBe(6);
    expect(sum([-1, -2, -3])).toBe(-6);
    expect(sum([])).toBe(0);
  });
});

describe('maxSubarray', () => {
  test('returns the subarray with the largest sum', () => {
    expect(maxSubarray([-4, 2, -5, 1, 2, 3, 6, -5, 1], 4)).toEqual([1, 2, 3, 6]);
    expect(maxSubarray([1, 2, 0, 5], 2)).toEqual([0, 5]);
  });

  test('handles empty array', () => {
    expect(maxSubarray([], 3)).toEqual([]);
  });
});
