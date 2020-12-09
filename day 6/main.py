from functools import reduce
from typing import List

plane_answers = []
with open("in.txt", "r") as f:
    group_answers = []
    for line in f:
        line_stripped = line.strip()
        if len(line_stripped) > 0:
            group_answers.append(line_stripped)
        else:
            plane_answers.append(group_answers)
            group_answers = []
    plane_answers.append(group_answers)


# Part 1
def count_distinct_questions(answers: List[str]) -> int:
    questions = set([
        q
        for answer in answers
        for q in answer
    ])
    return len(questions)


print(
    "Part 1:",
    reduce(
        lambda x, y: x + y,
        map(
            count_distinct_questions,
            plane_answers
        )
    )
)


# Part 2
def count_all_yes_questions(answers: List[str]) -> int:
    questions = [
        set(answer)
        for answer in answers
    ]
    questions_all_yes = reduce(lambda x, y: x.intersection(y), questions)
    return len(questions_all_yes)


print(
    "Part 2:",
    reduce(
        lambda x, y: x + y,
        map(
            count_all_yes_questions,
            plane_answers
        )
    )
)
