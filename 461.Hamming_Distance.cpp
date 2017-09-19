// https://leetcode.com/problems/hamming-distance/description/
// Input: x = 1, y = 4

// Output: 2

// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)
//        ↑   ↑

// The above arrows point to positions where the corresponding bits are
// different.

// 5 -> 0101 -> 2
// 6 -> 0110 -> 2
// 7 -> 0111 -> 3

int countOne(int x) {
  int count = 0;
  while (x > 0) {
    if (x % 2 == 1) {
      count++;
    }
    x /= 2;
  }
  return count;
}

class Solution {
public:
  int hammingDistance(int x, int y) {
    int ans = countOne(x ^ y);
    return ans;
  }
};

// Other solution
// class Solution {
// public:
//     int hammingDistance(int x, int y) {

//         int num=x^y;
//         int count=0;
//         while(num){
//             count+=num&1;
//             num>>=1;
//         }
//         return count;
//     }
// };
