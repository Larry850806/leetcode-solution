#include <bits/stdc++.h>
#include <cstdio>
#include <string>

using namespace std;

const string numToWord[] = {
    "zero",        "one",        "two",          "three",        "four",
    "five",        "six",        "seven",        "eight",        "nine",
    "ten",         "eleven",     "twelve",       "thirteen",     "fourteen",
    "fifteen",     "sixteen",    "seventeen",    "eighteen",     "nineteen",
    "twenty",      "twenty one", "twenty two",   "twenty three", "twenty four",
    "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine",
    "thirty"};

// Complete the timeInWords function below.
string timeInWords(int h, int m) {
  char str[50];
  cout << h << " " << m << endl;

  if (m == 0) {
    sprintf(str, "%s o' clock", numToWord[h].c_str());
  } else if (m >= 1 && m <= 30) {
    if (m == 1) {
      sprintf(str, "one minute past %s", numToWord[h].c_str());
    } else if (m == 15) {
      sprintf(str, "quarter past %s", numToWord[h].c_str());
    } else if (m == 30) {
      sprintf(str, "half past %s", numToWord[h].c_str());
    } else {
      sprintf(str, "%s minutes past %s", numToWord[m].c_str(),
              numToWord[h].c_str());
    }
  } else if (m == 45) {
    sprintf(str, "quarter to %s", numToWord[h + 1].c_str());
  } else if (m == 59) {
    sprintf(str, "one minute to %s", numToWord[h + 1].c_str());
  } else {
    sprintf(str, "%s minutes to %s", numToWord[60 - m].c_str(),
            numToWord[h + 1].c_str());
  }
  return string(str);
}

int main() {
  ofstream fout(getenv("OUTPUT_PATH"));

  int h;
  cin >> h;
  cin.ignore(numeric_limits<streamsize>::max(), '\n');

  int m;
  cin >> m;
  cin.ignore(numeric_limits<streamsize>::max(), '\n');

  string result = timeInWords(h, m);

  fout << result << "\n";

  fout.close();

  return 0;
}
