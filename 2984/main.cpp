# include <iostream>

using namespace std;

int main () {
    string user_text;
    cin >> user_text;

    int open  = 0;
    for (string::size_type i = 0; i < user_text.size(); i++) {
        if (user_text[i] == ')' && open > 0) {
            open -= 1;
        }
        else if (user_text[i] == '(') {
            open++;
        }
    }

    if (!(open)) {
        cout << "Partiu RU!" << endl;
    } else {
        cout << "Ainda temos " << open << " assunto(s) pendente(s)!" << endl;
    }
    return 0;
}
