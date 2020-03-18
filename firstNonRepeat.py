#!/usr/bin/python3
import time
import random
from typing import List

"""
Given a string s consisting of lowercase English letters,
find and return the first instance of a non-repeating character
in it. If there is no such character return "_"
"""


def first_non_repeat_via_dict(s: str) -> str:
    # use set datastructure to use uniqueness property for checking
    # 2 loops across entire list O(2N)
    letter_occurrence_map = {}
    for char in s:
        if letter_occurrence_map.get(char) is None:
            letter_occurrence_map[char] = 1
        else:
            letter_occurrence_map[char] += 1
    for char in s:
        if letter_occurrence_map[char] == 1:
            return char
    return "_"


def first_non_repeat_via_str_methods(s: str) -> str:
    for char in s:
        if s.index(char) == s.rindex(char):
            return char
    return "_"


def main():
    s_length = 10
    s = "".join([random.choice("abcde") for i in range(s_length)])

    print(s)
    start = time.time()
    print(first_non_repeat_via_dict(s))
    print(
        f"Finding first non-repeating character with map took: {(time.time()-start)*1000} milliseconds"
    )

    start = time.time()
    print(first_non_repeat_via_str_methods(s))
    print(
        f"Finding first non-repeating character with r/index() took: {(time.time()-start)*1000} milliseconds"
    )


if __name__ == "__main__":
    main()
