import itertools

with open('input.txt') as f:
    lines = f.readlines()

numbers = [int(v.strip()) for v in lines]
for a, b in itertools.combinations(numbers, 2):
    if a + b == 2020:
        print(a * b)
