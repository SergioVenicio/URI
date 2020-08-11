def calc(dividend, divider):
    return dividend / divider


def main():
    calcs = int(input())
    while calcs:
        dividend, divider = map(int, input().split(' '))
        try:
            print(calc(dividend, divider))
        except ZeroDivisionError:
            print('divisao impossivel')

        calcs -= 1


if __name__ == '__main__':
    main()