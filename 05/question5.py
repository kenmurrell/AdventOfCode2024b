from typing import Dict
from typing import List
from typing import Tuple

def getMiddle(l: List[int]) -> int:
    return l[int(len(l)/2)]

def checkOrder(rules: Dict[int, List[int]], print: List[int]) -> bool:
    for i, c in enumerate(print):
        if c not in rules:
            continue
        afterNumbers = rules[c]
        if any(n in afterNumbers for n in print[:i]):
             return False
    return True

def load(filename: str) -> Tuple[Dict[int, List[int]], List[int]]:
    rules = dict()
    prints = list()
    isRules = True
    with open(filename) as f:
        for line in f:
            if len(line.strip("\n")) == 0:
                isRules = False
                continue
            if isRules:
                first = int(line.split("|")[0])
                second = int(line.split("|")[1])
                if first in rules:
                        rules[first].append(second)
                else:
                        rules[first] = [second]
            else:
                prints.append([int(c) for c in line.split(",")])
    return rules, prints

def Part1(filename: str):
    rules, prints = load(filename)
    sum = 0
    for print in prints:
        if checkOrder(rules, print):
             sum += getMiddle(print)
    return sum

def slowSort(rules: Dict[int, List[int]], print: List[int]) -> List[int]:
    for i, c in enumerate(print):
        if c not in rules:
            continue
        afterNumbers = rules[c]
        if any(n in afterNumbers for n in print[:i]):
            print[i], print[i-1] = print[i-1], print[i]
            return slowSort(rules, print)
    return print

def Part2(filename: str):
    rules, prints = load(filename)
    sum = 0
    for print in prints:
        if not checkOrder(rules, print):
            sortedprint = slowSort(rules, print)
            sum += getMiddle(sortedprint)
    return sum

test1 = Part1("test.txt")
print("Test1: {} (expected 143)".format(test1))
ans1 = Part1("data.txt")
print("Answer1: {} (expected 6260)".format(ans1))

test2 = Part2("test.txt")
print("Test2: {} (expected 123)".format(test2))
ans2 = Part2("data.txt")
print("Answer2: {} (expected )".format(ans2))
