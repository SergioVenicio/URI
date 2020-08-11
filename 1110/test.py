import pytest

from app import get_cards, shuffle


def test_get_card():
    n = 10
    cards = get_cards(n)

    assert cards == [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]


def test_one_card_suffle():
    n = 1
    assert shuffle(n) == (
        'Discarded cards: \n'
        'Remaining card: 1'
    )


def test_seven_cards_suffle():
    n = 7
    assert shuffle(n) == (
        'Discarded cards: 1, 3, 5, 7, 4, 2\n'
        'Remaining card: 6'
    )


def test_udebug_case():
    for i in range(50):
        n = i + 1
        expeted = shuffle(n)
        assert 'Discarded' in expeted and 'Remaining' in expeted


def test_zero_cards_suffle():
    n = 0
    assert shuffle(n) == ''
