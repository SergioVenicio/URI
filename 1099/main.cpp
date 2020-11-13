#include <iostream>

using namespace std;

int get_ods_sums(int min_value, int max_value) {
    int sum = 0;
    for (int aux = min_value + 1; aux < max_value; aux++) {
        if (aux % 2 != 0) {
            sum += aux;
        }
    }
    return sum;
}

int main() {
    int N;
    cin >> N;

    int X;
    int Y;
    int result = 0;
    for (int i = 0; i < N; i++) {
        cin >> X;
        cin >> Y;

        if (X <= Y) {
            result = get_ods_sums(X, Y);
        } else {
            result = get_ods_sums(Y, X);
        }
        cout << result << endl;
    }
    return 0;
}
