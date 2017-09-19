// https://leetcode.com/problems/hamming-distance/description/

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
