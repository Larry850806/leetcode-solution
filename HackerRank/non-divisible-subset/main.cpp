#include <bits/stdc++.h>
#include <vector>

using namespace std;

vector<string> split_string(string);

void print(vector<int> &vec) {
  for (auto it = vec.begin(); it != vec.end(); it++) {
    cout << *it << " ";
  }
  cout << endl;
}

// Complete the nonDivisibleSubset function below.
int nonDivisibleSubset(int k, vector<int> S) {
  // If S = {1, 3, 5, 6} and k = 3
  // 1 % 3 = 1
  // 3 % 3 = 0
  // 5 % 3 = 2
  // 6 % 3 = 0
  // there is two 0, one 1 and one 2
  // => modByKAmount[0] = 2
  // => modByKAmount[1] = 1
  // => modByKAmount[2] = 1
  vector<int> modByKAmount(k);
  for (auto it = S.begin(); it != S.end(); it++) {
    int value = *it;
    modByKAmount[value % k]++;
  }

  print(modByKAmount);

  bool isKEven = k % 2 == 0;
  if (isKEven) {
    // if k is 6
    // add max(modByKAmount[1], modByKAmount[5])
    // add max(modByKAmount[2], modByKAmount[4])
    // if(modByKAmount[0] is bigger than 0), pick one
    // if(modByKAmount[3] is bigger than 0), pick one
    int result = 0;
    for (int i = 1; i < k / 2; i++) {
      int a = i, b = k - i;
      result += max(modByKAmount[a], modByKAmount[b]);
    }
    if (modByKAmount[0] > 0) {
      result++;
    }
    if (modByKAmount[k / 2] > 0) {
      result++;
    }
    return result;
  } else {
    // if k is 5
    // add max(modByKAmount[1], modByKAmount[4])
    // add max(modByKAmount[2], modByKAmount[3])
    // if(modByKAmount[0] is bigger than 0), pick one
    int result = 0;
    for (int i = 1; i <= k / 2; i++) {
      int a = i, b = k - i;
      result += max(modByKAmount[a], modByKAmount[b]);
    }
    if (modByKAmount[0] > 0) {
      result++;
    }
    return result;
  }
}

int main() {
  ofstream fout(getenv("OUTPUT_PATH"));

  string nk_temp;
  getline(cin, nk_temp);

  vector<string> nk = split_string(nk_temp);

  int n = stoi(nk[0]);

  int k = stoi(nk[1]);

  string S_temp_temp;
  getline(cin, S_temp_temp);

  vector<string> S_temp = split_string(S_temp_temp);

  vector<int> S(n);

  for (int i = 0; i < n; i++) {
    int S_item = stoi(S_temp[i]);

    S[i] = S_item;
  }

  int result = nonDivisibleSubset(k, S);

  fout << result << "\n";

  fout.close();

  return 0;
}

vector<string> split_string(string input_string) {
  string::iterator new_end =
      unique(input_string.begin(), input_string.end(),
             [](const char &x, const char &y) { return x == y and x == ' '; });

  input_string.erase(new_end, input_string.end());

  while (input_string[input_string.length() - 1] == ' ') {
    input_string.pop_back();
  }

  vector<string> splits;
  char delimiter = ' ';

  size_t i = 0;
  size_t pos = input_string.find(delimiter);

  while (pos != string::npos) {
    splits.push_back(input_string.substr(i, pos - i));

    i = pos + 1;
    pos = input_string.find(delimiter, i);
  }

  splits.push_back(
      input_string.substr(i, min(pos, input_string.length()) - i + 1));

  return splits;
}
