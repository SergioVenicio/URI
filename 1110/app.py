def shuffle(n):
    if n <= 0:
        return ''

    discarteds = []
    cards = get_cards(n)

    while True:
        if len(cards) <= 1:
            break

        discarteds.append(cards.pop())
        top = cards.pop()

        cards = [top] + cards

    return (
        "Discarded cards: " + ', '.join(int_list_to_string(discarteds)) + "\n"
        "Remaining card: " + str(cards.pop())
    )


def get_cards(n):
    return [card + 1 for card in range(n)][::-1]


def int_list_to_string(int_list):
    return [str(x) for x in int_list]


if __name__ == '__main__':
    while True:
        n = int(input())
        if n == 0:
            break

        print(shuffle(n))
