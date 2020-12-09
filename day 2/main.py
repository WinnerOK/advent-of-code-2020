from collections import Counter

valid = 0


def describe_line(line: str) -> (int, int, str, str):
    rule, data = line.strip().split(":")
    rule_range, rule_symbol = rule.split(" ")
    rule_range_min, rule_range_max = map(int, rule_range.split("-"))
    return rule_range_min, rule_range_max, rule_symbol, data


with open("in.txt", "r") as f:
    for line in f:
        min_occurrences, max_occurrences, required_symbol, line_data = describe_line(line)
        cnt = Counter(line_data)
        if min_occurrences <= cnt[required_symbol] <= max_occurrences:
            valid += 1

print("1) Valid cnt:", valid)

valid = 0
with open("in.txt", "r") as f:
    for line in f:
        position_1, position_2, required_symbol, line_data = describe_line(line)
        # line is not stripped, so there is a space at index 0
        if (line_data[position_1] == required_symbol) ^ (line_data[position_2] == required_symbol):
            valid += 1

print("2) Valid cnt:", valid)
