# Advent of Code 2020

[![Tests](https://github.com/devries/advent_of_code_2020/actions/workflows/main.yml/badge.svg)](https://github.com/devries/advent_of_code_2020/actions/workflows/main.yml)

## Summary

Another year of [Advent of Code](https://adventofcode.com/) is over, and as usual
it provided a daily dose of fun in the countdown to Christmas. This year I needed
something like it more than most, as I think we all did. This was also the first
year that my daughter and I worked on problems together. She is just learning
how to write code in a high school class, so many of the concepts are very new
to her, but I enjoyed helping her and talking through the problems. Unfortunately
as her midterms started, she didn't have a lot of time to work on Advent of Code,
but we hopefully will be going back to look at some of the more interesting problems.

I really enjoyed the [Intcode problems from 2019](https://github.com/devries/advent_of_code_2019)
and missed having something we were building over time this year. I also thought
there were some extremely challenging problems last year, which at times made it
pretty difficult to keep up, but this year seemed to have fewer very difficut
problems. The final weekend of Advent of Code (days 19 and 20) were problems I
thought were the most difficult, but this year I was able to mainly work on the
problems before starting work each morning.

This year I tried to write more tests to get into a better habit of using tests,
and that mostly worked out, except for day 19.
Day 19 gave me the most trouble. I got the first part fairly quickly, but I had
a bug in my code which manifested itself in the second part. It took a long
time to debug the issue, and in that time I managed to introduce several more
bugs. I also had an error in my tests, which may have led me to keep trying to
debug the problem after it was working. I had to do a lot of cooking that night
and decided I would rewrite part 2 entirely after dinner. I then was able to
find my mistakes and finish the problem. In the end nearly 8 hours elapsed
between the time I finished part 1 and finished part 2.

I thought day 20 was a lot of fun, and I managed to get through this one with
few bugs, even though it took some time. I enjoyed considering all the orientations
as well as the rotations or flips required to go between them. I also enjoyed day
10 quite a bit. Juggling all those joltage adapters and figuring out the total
number of possible adapter combinations to charge your handheld.

I thought the hexagonal grid was pretty interesting, and I suspected most people
would use the same coordinate system I used (east and west are given by 2 moves
along the X-axis, whereas northeast, northwest, southeast, and southwest would be
given by 1 move in the X-axis and 1 move in the Y-axis), however I was surprised
that many people just tilted the Y-axis to align with either the northeast or
northwest direction, and treated east and west as one step in the X-axis,
northwest and southeast as 1 step in the Y direction, and northeast and southwest
as one step in each of the X and Y axes. 

Even though this is only my second year, I feel like this event really starts to
put me in the Christmas mood, and I look forward to next year.

## Index

- [Day 1: Report Repair](https://adventofcode.com/2020/day/1) - [part 1](day01_p1), [part 2](day01_p2)
- [Day 2: Password Philosophy](https://adventofcode.com/2020/day/2) - [part 1](day02_p1), [part 2](day02_p2) 
- [Day 3: Toboggan Trajectory](https://adventofcode.com/2020/day/3) - [part 1](day03_p1), [part 2](day03_p2)
- [Day 4: Passport Processing](https://adventofcode.com/2020/day/4) - [part 1](day04_p1), [part 2](day04_p2)
- [Day 5: Binary Boarding](https://adventofcode.com/2020/day/5) - [part 1](day05_p1), [part 2](day05_p2)
- [Day 6: Custom Customs](https://adventofcode.com/2020/day/6) - [part 1](day06_p1), [part 2](day06_p2), [part 2 refactor](day06_p2alt)
- [Day 7: Handy Haversacks](https://adventofcode.com/2020/day/7) - [part 1](day07_p1), [part 2](day07_p2), [part 1 refactor](day07_p1alt), [part 2 refactor](day07_p2alt)
- [Day 8: Handheld Halting](https://adventofcode.com/2020/day/8) - [part 1](day08_p1), [part 2](day08_p2)
- [Day 9: Encoding Error](https://adventofcode.com/2020/day/9) - [part 1](day09_p1), [part 2](day09_p2)
- [Day 10: Adapter Array](https://adventofcode.com/2020/day/10) - [part 1](day10_p1), [part 2](day10_p2)
- [Day 11: Seating System](https://adventofcode.com/2020/day/11) - [part 1](day11_p1), [part 2](day11_p2)
- [Day 12: Rain Risk](https://adventofcode.com/2020/day/12) - [part 1](day12_p1), [part 2](day12_p2)
- [Day 13: Shuttle Search](https://adventofcode.com/2020/day/13) - [part 1](day13_p1), [part 2](day13_p2)
- [Day 14: Docking Data](https://adventofcode.com/2020/day/14) - [part 1](day14_p1), [part 2](day14_p2)
- [Day 15: Rambunctious Recitation](https://adventofcode.com/2020/day/15) - [part 1](day15_p1), [part 2](day15_p2)
- [Day 16: Ticket Translation](https://adventofcode.com/2020/day/16) - [part 1](day16_p1), [part 2](day16_p2)
- [Day 17: Conway Cubes](https://adventofcode.com/2020/day/17) - [part 1](day17_p1), [part 2](day17_p2)
- [Day 18: Operation Order](https://adventofcode.com/2020/day/18) - [part 1](day18_p1), [part 2](day18_p2)
- [Day 19: Monster Messages](https://adventofcode.com/2020/day/19) - [part 1](day19_p1), [part 2](day19_p2)
- [Day 20: Jurassic Jigsaw](https://adventofcode.com/2020/day/20) - [part 1](day20_p1), [part 2](day20_p2)
- [Day 21: Allergen Assessment](https://adventofcode.com/2020/day/21) - [part 1](day21_p1), [part 2](day21_p2)
- [Day 22: Crab Combat](https://adventofcode.com/2020/day/22) - [part 1](day22_p1), [part 2](day22_p2)
- [Day 23: Crab Cups](https://adventofcode.com/2020/day/23) - [part 1](day23_p1), [part 2](day23_p2)
- [Day 24: Lobby Layout](https://adventofcode.com/2020/day/24) - [part 1](day24_p1), [part 2](day24_p2)
- [Day 25: Combo Breaker](https://adventofcode.com/2020/day/25) - [part 1](day25_p1)
