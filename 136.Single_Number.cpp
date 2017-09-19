// https://leetcode.com/problems/single-number/description/

class Solution {
 public:
  // 0 ^ x = x
  // x ^ y = y ^ x
  // x ^ x = 0
  // 1 ^ 2 ^ 1 = 1 ^ 1 ^ 2 = 0 ^ 2 = 2
  // 2 is single number
  int singleNumber(vector<int>& nums) {
    int result = 0;
    for (int i = 0; i < nums.size(); i++) {
      result ^= nums[i];
    }
    return result;
  }
};

// class Solution {
//  public:
//   int singleNumber(vector<int>& nums) {
//     int result = 0;
//     for (int num : nums) {
//       result ^= num;
//     }
//     return result;
//   }
// };