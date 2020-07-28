def createNumbers(nMin, nMax):
    if nMin >= nMax:
        return [nMax]

    return [nMin] + createNumbers(nMin+1, nMax)


def parseInput(userString):
    numbers = [n for n in map(int, userString.split(' '))]

    if any(map(lambda x: x <= 0, numbers)):
        return None

    return numbers

def toString(numbers):
    return ' '.join(map(str, numbers))

def userInput():
    numbers = parseInput(input())

    if numbers is None:
        return

    elements = createNumbers(*sorted(numbers))
    print("{} Sum={}".format(
        toString(elements),
        sum(elements)
    ))
    userInput()


userInput()