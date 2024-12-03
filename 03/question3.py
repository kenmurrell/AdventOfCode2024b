from enum import Enum

class Mode(Enum):
    Idle = 1
    Parse = 2
    FirstNumber = 3
    SecondNumber = 4

class Window:
    def __init__(self):
        self.win = ""
    def add(self, i: str):
        if len(self.win) == 7:
            self.win = self.win[1:] + i
        else:
            self.win += i
    def isMult(self):
        return self.__check__(self.win, "mul(")
    def isOn(self):
        return self.__check__(self.win, "do()")
    def isOff(self):
        return self.__check__(self.win, "don't()")
    @staticmethod
    def __check__(s1: str, s2: str) -> bool:
        return s1[-len(s2):] == s2

def Part1(filename: str) -> int:
    data = ""
    with open(filename) as f:
        for line in f:
            data += line.strip()
    window, sum = Window(), 0
    m, firstNumber, secondNumber = Mode.Parse, "", ""

    for c in data:
        window.add(c)
        if m == Mode.Parse and window.isMult():
            m = Mode.FirstNumber
        elif m == Mode.FirstNumber:
            if c.isdigit():
                firstNumber += c
            elif c == ",":
                m = Mode.SecondNumber
            else:
                m, firstNumber, secondNumber = Mode.Parse, "", ""
        elif m == Mode.SecondNumber:
            if c.isdigit():
                secondNumber += c
            elif c == ")":
                sum += int(firstNumber) * int(secondNumber)
                m, firstNumber, secondNumber = Mode.Parse, "", ""
            else:
                m, firstNumber, secondNumber = Mode.Parse, "", ""
    return sum

def Part2(filename: str) -> int:
    data = ""
    with open(filename) as f:
        for line in f:
            data += line.strip()
    window, sum = Window(), 0
    m, firstNumber, secondNumber = Mode.Parse, "", ""

    for c in data:
        window.add(c)
        if m == Mode.Idle and window.isOn():
            m, firstNumber, secondNumber = Mode.Parse, "", ""
        elif m == Mode.Parse:
            if window.isMult():
                m = Mode.FirstNumber
            elif window.isOff():
                m = Mode.Idle
        elif m == Mode.FirstNumber:
            if c.isdigit():
                firstNumber += c
            elif c == ",":
                m = Mode.SecondNumber
            else:
                m, firstNumber, secondNumber = Mode.Parse, "", ""
        elif m == Mode.SecondNumber:
            if c.isdigit():
                secondNumber += c
            elif c == ")":
                sum += int(firstNumber) * int(secondNumber)
                m, firstNumber, secondNumber = Mode.Parse, "", ""
            else:
                m, firstNumber, secondNumber = Mode.Parse, "", ""
    return sum

test1 = Part1("test1.txt")
print("Test1: {} (expected 161)".format(test1))
ans1 = Part1("data.txt")
print("Answer1: {} (expected 184576302)".format(ans1))

test2 = Part2("test2.txt")
print("Test2: {} (expected 48)".format(test2))
ans2 = Part2("data.txt")
print("Answer2: {} (expected 118173507)".format(ans2))