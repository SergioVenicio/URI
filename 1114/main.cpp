#include <iostream>

using namespace std;


int main() {
    int password;
    while(true) {
        cin >> password;

        if(password == 2002) {
            cout << "Acesso Permitido" << endl;
            break;
        }

        cout << "Senha Invalida" << endl;
    }
    return 0;
}
