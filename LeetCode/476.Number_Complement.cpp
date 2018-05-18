vector<long> getPower2() {
  vector<long> power2;
  power2.push_back(1);
  for (int i = 0; i < 31; i++) {
    power2.push_back(power2[i] * 2);
  }
  // power2 = [1, 2, 4, 8, 16, ... , 2^31]
  return power2;
}

class Solution {
 public:
  // find number in power2 bigger than 5 -> 8
  // (8 -1) - 5 = 2
  int findComplement(int num) {
    vector<long> power2 = getPower2();
    int biggerValue = *upper_bound(power2.begin(), power2.end(), num);
    return biggerValue - 1 - num;
  }
};

// class Solution {
//  public:
//   int findComplement(int num) {
//     long loop = num;
//     for (long i = 1; i <= loop; i <<= 1) {
//       num ^= i;
//     }
//     return num;
//   }
// };