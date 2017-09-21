// https://leetcode.com/problems/array-partition-i/description/

class Solution {
 public:
  // [1, 4, 2, 3]
  // => [1, 2, 3, 4]
  // => [1, 3]
  // => 4
  int arrayPairSum(vector<int>& nums) {
    vector<int> sortedNums(nums);
    sort(sortedNums.begin(), sortedNums.end());

    int maxSum = 0;
    for (int i = 0; i < sortedNums.size(); i++) {
      if (i % 2 == 0) {
        maxSum += sortedNums[i];
      }
    }

    return maxSum;
  }
};