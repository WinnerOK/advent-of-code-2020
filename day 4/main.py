import re

REQUIRED_FIELDS = {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}


def transform_passport(s: str) -> dict:
    return dict([
        data.split(":")
        for data in s.split(" ")
    ])


def validate_passport(passport_data: dict) -> bool:
    def validate_int(number: str, minimum: int, maximum: int) -> bool:
        try:
            n = int(number)
        except ValueError:
            return False
        else:
            return minimum <= n <= maximum

    def validate_byr(birth_year: str) -> bool:
        return validate_int(birth_year, 1920, 2002)

    def validate_iyr(issue_year: str) -> bool:
        return validate_int(issue_year, 2010, 2020)

    def validate_eyr(expiration_year: str) -> bool:
        return validate_int(expiration_year, 2020, 2030)

    def validate_hgt(height: str) -> bool:
        match = re.fullmatch(r'(\d+)(in|cm)', height)
        if not match:
            return False
        else:
            numerical, unit = match.groups()
            if unit == "cm":
                return validate_int(numerical, 150, 193)
            elif unit == "in":
                return validate_int(numerical, 59, 76)
            else:
                raise ValueError("Unknown unit passed regex")

    def validate_hcl(color: str) -> bool:
        return bool(re.fullmatch(r'#(\d|[a-f]){6}', color))

    def validate_ecl(color: str) -> bool:
        return color in ("amb", "blu", "brn", "gry", "grn", "hzl", "oth")

    def validate_pid(passport_id: str) -> bool:
        return bool(re.fullmatch(r"\d{9}", passport_id))

    def validate_cid(_) -> bool:
        return True

    validators_context = locals()

    if len(REQUIRED_FIELDS.difference(passport.keys())) == 0:
        return all(map(
            lambda field: validators_context[f"validate_{field}"](passport[field]),
            REQUIRED_FIELDS
        ))


if __name__ == '__main__':
    passports = []
    # Read all passports
    with open("in.txt", "r") as f:
        passport = ""
        for line in f:
            line_stripped = line.strip()
            if len(line_stripped) > 0:
                passport += line_stripped + " "
            else:
                passports.append(
                    transform_passport(passport.strip())
                )
                passport = ""
        passports.append(
            transform_passport(passport.strip())
        )

    # Make sure all required fields are present
    valid = 0
    for passport in passports:
        if len(REQUIRED_FIELDS.difference(passport.keys())) == 0:
            valid += 1
    print("Part 1:", valid)

    valid = 0
    for passport in passports:
        if validate_passport(passport):
            valid += 1
    print("Part 2:", valid)
