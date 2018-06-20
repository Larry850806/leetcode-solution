#include <bits/stdc++.h>
#include <algorithm>
#include <climits>
#include <iterator>
#include <map>

using namespace std;

vector<string> split_string(string);

// O(n)
void rm_dup(vector<int> &vec) {
  auto it = unique(vec.begin(), vec.end());
  vec.resize(distance(vec.begin(), it));
}

void print(vector<int> &vec) {
  for (auto it = vec.begin(); it != vec.end(); it++) {
    cout << *it << " ";
  }
  cout << endl;
}

// Complete the climbingLeaderboard function below.
vector<int> climbingLeaderboard(vector<int> scores, vector<int> alice) {
  rm_dup(scores);
  reverse(scores.begin(), scores.end());
  scores.insert(scores.begin(), INT_MIN);
  scores.push_back(INT_MAX);

  print(scores);
  print(alice);

  // scores = [INT_MIN 10 20 40 50 100 INT_MAX]
  // alice  = [5 5 25 50 120]

  vector<int> ranks;
  auto scores_it = scores.begin();
  auto alice_it = alice.begin();

  while (alice_it != alice.end()) {
    int score = *scores_it;
    int alice_score = *alice_it;

    cout << score << " " << alice_score << endl;

    if (alice_score == score) {
      int rank = distance(scores_it, scores.end()) - 1;
      ranks.push_back(rank);
      alice_it++;
      continue;
    }

    if (alice_score > score) {
      scores_it++;
      continue;
    }

    if (alice_score < score) {
      int rank = distance(scores_it, scores.end());
      ranks.push_back(rank);
      alice_it++;
      continue;
    }
  }

  return ranks;
}

int main() {
  ofstream fout(getenv("OUTPUT_PATH"));

  int scores_count;
  cin >> scores_count;
  cin.ignore(numeric_limits<streamsize>::max(), '\n');

  string scores_temp_temp;
  getline(cin, scores_temp_temp);

  vector<string> scores_temp = split_string(scores_temp_temp);

  vector<int> scores(scores_count);

  for (int i = 0; i < scores_count; i++) {
    int scores_item = stoi(scores_temp[i]);

    scores[i] = scores_item;
  }

  int alice_count;
  cin >> alice_count;
  cin.ignore(numeric_limits<streamsize>::max(), '\n');

  string alice_temp_temp;
  getline(cin, alice_temp_temp);

  vector<string> alice_temp = split_string(alice_temp_temp);

  vector<int> alice(alice_count);

  for (int i = 0; i < alice_count; i++) {
    int alice_item = stoi(alice_temp[i]);

    alice[i] = alice_item;
  }

  vector<int> result = climbingLeaderboard(scores, alice);

  for (int i = 0; i < result.size(); i++) {
    fout << result[i];

    if (i != result.size() - 1) {
      fout << "\n";
    }
  }

  fout << "\n";

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
