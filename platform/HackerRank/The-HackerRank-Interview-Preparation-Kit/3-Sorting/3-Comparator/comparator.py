#!/usr/bin/env python3

from functools import cmp_to_key
from dataclasses import dataclass


@dataclass
class Player:
    name: str
    score: int

    def __lt__(self, other):
        if self.score < other.score:
            return True
        return self.score == other.score and self.name > other.name

    @staticmethod
    def comparator(lhs, rhs):
        return (lhs < rhs) - (lhs > rhs)


def main():
    n = int(input())
    data = []
    for i in range(n):
        name, score = input().split()
        score = int(score)
        player = Player(name, score)
        data.append(player)

    data = sorted(data, key=cmp_to_key(Player.comparator))
    for i in data:
        print(i.name, i.score)


if __name__ == "__main__":
    main()
