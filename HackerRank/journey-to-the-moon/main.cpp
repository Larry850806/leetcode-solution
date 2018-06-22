#include <bits/stdc++.h>
#include <algorithm>
#include <climits>
#include <unordered_set>
#include <vector>

using namespace std;

vector<string> split_string(string);

template <class T>
void print(vector<T> &vec) {
  for (auto it = vec.begin(); it != vec.end(); it++) {
    cout << *it << " ";
  }
  cout << endl;
}

template <class T>
void print(unordered_set<T> &s) {
  for (auto it = s.begin(); it != s.end(); it++) {
    cout << *it << " ";
  }
  cout << endl;
}

long long int C(int n, int k) {
  long long int result = 1;
  for (int i = 0; i < k; i++) {
    result *= n - i;
  }
  for (int i = 1; i <= k; i++) {
    result /= i;
  }
  return result;
}

// The answer of Testcase 11 is 4999949998
// But 4999949998 is FUCKING bigger than INT_MAX(2147483647)
// so change "int journeyToMoon(...)" to "long long int journeyToMoon(...)"
// alse change "int result = journeyToMoon(...)" to long long in main function

// Complete the journeyToMoon function below.
long long int journeyToMoon(int n, vector<vector<int>> astronaut) {
  cout << INT_MAX << endl;
  // In the beginning, every astronaut from 0 to n-1 is a country
  // a country is a set that contains astronaut id
  // 0 -> {0}
  // 1 -> {2}
  // 2 -> {2}
  // ...
  vector<unordered_set<int> *> countries(n);
  for (int i = 0; i < n; i++) {
    countries[i] = new unordered_set<int>({i});
  }

  for (int i = 0; i < astronaut.size(); i++) {
    // if pair is [0, 1], merge 0 and 1 country
    // 0 -> {0, 1}
    // 1 -> {0, 1}
    // 2 -> {2}
    // ...
    auto pair = astronaut[i];
    auto countryA = countries[pair[0]];
    auto countryB = countries[pair[1]];
    if (countryA == countryB) {
      continue;
    }
    countryA->insert(countryB->begin(), countryB->end());
    for (auto it = countryB->begin(); it != countryB->end(); it++) {
      int mergedCountry = *it;
      countries[mergedCountry] = countryA;
    }
    delete countryB;
  }

  // after merged
  // 0 -> {0, 1, 4}
  // 1 -> {0, 1, 4}
  // 2 -> {2, 3}
  // 3 -> {2, 3}
  // 4 -> {0, 1, 4}
  // 5 -> {5}
  // remove duplicate sets
  // => [{0, 1, 4}, {2, 3}, {5}]
  unordered_set<unordered_set<int> *> uniq_countries_set(countries.begin(),
                                                         countries.end());
  vector<unordered_set<int> *> uniq_countries_vec(uniq_countries_set.begin(),
                                                  uniq_countries_set.end());

  // count amount
  // [{0, 1, 4}, {2, 3}, {5}] => [3, 2, 1]
  // sum = 6
  vector<int> amounts(uniq_countries_vec.size());
  int sum = 0;
  for (auto i = 0; i < uniq_countries_vec.size(); i++) {
    amounts[i] = uniq_countries_vec[i]->size();
    sum += amounts[i];
    delete uniq_countries_vec[i];
  }

  // [3, 2, 1] => C(3+2+1, 2) - C(3, 2) - C(2, 2)
  // = C(6, 2) - C(3, 2) - C(2, 2)
  // = 15 - 3 - 1
  // = 11
  long long int result = C(sum, 2);
  for (auto it = amounts.begin(); it != amounts.end(); it++) {
    int amount = *it;
    if (amount >= 2) {
      result -= C(amount, 2);
    }
  }

  return result;
}

int main() {
  ofstream fout(getenv("OUTPUT_PATH"));

  string np_temp;
  getline(cin, np_temp);

  vector<string> np = split_string(np_temp);

  int n = stoi(np[0]);

  int p = stoi(np[1]);

  vector<vector<int>> astronaut(p);
  for (int i = 0; i < p; i++) {
    astronaut[i].resize(2);

    for (int j = 0; j < 2; j++) {
      cin >> astronaut[i][j];
    }

    cin.ignore(numeric_limits<streamsize>::max(), '\n');
  }

  long long int result = journeyToMoon(n, astronaut);

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