import re


TRIANGULO_REX = re.compile(r'triangulo?s*')
CIRCULO_REX = re.compile(r'circulo?s*')
QUADRADO_REX = re.compile(r'quadrado?s*')
COUNTER = {
    1: 1,
    2: 0,
    3: 1,
}


def value_mapper(card):
    if TRIANGULO_REX.search(card):
        return 0
    if CIRCULO_REX.search(card):
        return 1
    if QUADRADO_REX.search(card):
        return 2


while True:
    entries = int(input())
    if entries == 0:
        break

    counter_cards = []
    cards = []
    result = 0
    for _ in range(entries):
        cards.append(value_mapper(input()))
        if len(cards) == 3:
            result += COUNTER[len(cards)]
            cards = []
            continue

    print(result)
