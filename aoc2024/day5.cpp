
#include <bits/stdc++.h>

using namespace std;

void pt1() {
    string tmp;

    map<int, set<int>> map;
    while (getline(cin, tmp)) {
        if (tmp == "") {
            break;
        }

        auto pos = tmp.find("|");
        auto l = stoi(tmp.substr(0, pos));
        auto r = stoi(tmp.substr(pos + 1));
        map[r].insert(l);
    }

    vector<vector<int>> vvi;
    while(getline(cin, tmp)) {
        std::vector<int> vi;
        std::stringstream ss (tmp);
        std::string item;

        while (getline(ss, item, ',')) {
            vi.push_back(stoi(item));
        }

        vvi.push_back(vi);
    }

    int sum = 0;
    for (auto ups : vvi) {
        for (auto up : ups) { cout << up << " "; }
        cout << endl;

        for (auto i = 1; i < ups.size(); ++i) {
            for (auto j = 0; j < i; ++j) {
                if (map[ups[j]].count(ups[i])) {
                    goto end;
                }
            }
        }
        sum += ups[ups.size() / 2];
        end:;
    }

    cout << endl <<  sum << endl;
}

void pt2() {
    string tmp;

    map<int, set<int>> map;
    while (getline(cin, tmp)) {
        if (tmp == "") {
            break;
        }

        auto pos = tmp.find("|");
        auto l = stoi(tmp.substr(0, pos));
        auto r = stoi(tmp.substr(pos + 1));

    }

    vector<vector<int>> vvi;
    while(getline(cin, tmp)) {
        std::vector<int> vi;
        std::stringstream ss (tmp);
        std::string item;

        while (getline(ss, item, ',')) {
            vi.push_back(stoi(item));
        }

        vvi.push_back(vi);
    }
}


int main() {
    pt2();
}