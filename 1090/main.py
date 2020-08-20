import re


TRIANGULO_REX = re.compile(r'triangulo?s*')
CIRCULO_REX = re.compile(r'circulo?s*')
QUADRADO_REX = re.compile(r'quadrado?s*')


def value_mapper(card):
    if TRIANGULO_REX.search(card):
        return 0
    if CIRCULO_REX.search(card):
        return 1
    if QUADRADO_REX.search(card):
        return 2


def valid_counter(cards):
    cards_len = len(list(cards))

    if cards_len == 1 or cards_len == 3:
        return 1

    return 0


while True:
    entries = int(input())
    if entries == 0:
        break

    counter_cards = []
    cards = []
    for _ in range(entries):
        cards.append(value_mapper(input()))
        if len(cards) == 3:
            counter_cards.append(valid_counter(cards))
            cards = []
            continue

    print(sum(counter_cards))
