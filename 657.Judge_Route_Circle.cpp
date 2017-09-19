// https://leetcode.com/problems/judge-route-circle/description/

class Solution {
 public:
  // O(nlogn)
  bool judgeCircle(string moves) {
    map<char, int> countMap;
    countMap['L'] = 0;
    countMap['R'] = 0;
    countMap['U'] = 0;
    countMap['D'] = 0;
    for (int i = 0; i < moves.size(); i++) {
      char direction = moves[i];
      countMap[direction]++;
    }
    return countMap['L'] == countMap['R'] && countMap['U'] == countMap['D'];
  }
};

// class Solution {
//  public:
//   bool judgeCircle(string moves) {
//     int up = 0;
//     int left = 0;
//     for (char c : moves) {
//       switch (c) {
//         case 'U':
//           ++up;
//           break;
//         case 'D':
//           --up;
//           break;
//         case 'L':
//           ++left;
//           break;
//         case 'R':
//           --left;
//           break;
//       }
//     }
//     if (!up && !left) {
//       return true;
//     }
//     return false;
//   }
// };