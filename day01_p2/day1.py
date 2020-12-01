import itertools

with open('input.txt') as f:
    lines = f.readlines()

numbers = [int(v.strip()) for v in lines]
for a, b, c in itertools.combinations(numbers, 3):
    if a + b + c == 2020:
        print(a * b * c)
