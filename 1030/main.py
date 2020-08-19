def main():
    test_cases = int(input())

    case = 0
    while case < test_cases:
        case += 1
        soldiers, jumps = map(int, input().split(' '))

        result = jump(soldiers, jumps) + 1
        print(f'Case {case}: {result}')


def jump(soldiers, jumps):
    result = 0
    for i in range(2, soldiers + 1):
        result = (result + jumps) % i

    return result


if __name__ == "__main__":
    main()
