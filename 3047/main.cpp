# include <iostream>

using namespace std;


int main() {
    int monica_age;
    int a, b, c;

    cin >> monica_age;
    cin >> a;
    cin >> b;
    c = monica_age - (a + b);

    if (a > b && a > c) {
        cout << a << endl;
    } else if (b > a && b > c) {
        cout << b << endl;
    } else {
        cout << c << endl;
    }
    return 0;
}
