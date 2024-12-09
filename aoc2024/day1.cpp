
#include <bits/stdc++.h>

using namespace std;

#define pii pair<int,int>


void pt1() {
    priority_queue<int, vector<int>, greater<int>> lq;
    priority_queue<int, vector<int>, greater<int>> rq;

    int l, r;
    int i = 0;
    while(cin >> l >> r) {
        lq.push(l);
        rq.push(r);
    }

    int total = 0;
    while (!lq.empty()) {
        auto l = lq.top();
        auto r = rq.top();

        lq.pop();
        rq.pop();

        total += abs(l - r);
    }

    cout << total << endl;
}

void part2() {
    map<int, int> freq;
    vector<int> lhs;

    int l, r;
    while(cin >> l >> r) {
        lhs.push_back(l);
        freq[r]++;
    }

    int total = 0;
    for (auto l : lhs) {
        total += l * freq[l];
    }

    cout << total << endl;
}

int main() {
    part2();
}