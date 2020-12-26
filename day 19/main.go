package main

import (
	"strconv"
	"strings"
)

func main() {
	input := readInput("./in.txt")
	splitIdx := 0
	for len(input[splitIdx]) != 0 {
		splitIdx += 1
	}
	rules := parseRules(input[0:splitIdx])
	tests := input[splitIdx+1:]
	part1(rules, tests)
	part2(rules, tests)
}

type ExactRule struct {
	character string
}

type AndRule struct {
	refs []int
}

type OrRule struct {
	options []AndRule
}

func lexRule(rule string) interface{} {
	if strings.Contains(rule, "\"") {
		character := strings.Split(rule, "\"")[1]
		return ExactRule{character: character}
	} else if strings.Contains(rule, "|") {
		optionsStr := strings.Split(rule, " | ")
		options := []AndRule{}
		for idx := range optionsStr {
			optionsStr[idx] = strings.TrimSpace(optionsStr[idx])
			parsedAndRule := lexRule(optionsStr[idx]).(AndRule)
			options = append(options, parsedAndRule)
		}
		return OrRule{options: options}
	} else {
		ruleRefs := stringSliceToIntSlice(strings.Split(strings.TrimSpace(rule), " "))
		return AndRule{refs: ruleRefs}
	}
}

func parseRules(rules []string) map[int]interface{} {
	parsedRules := map[int]interface{}{}
	for _, rule := range rules {
		ruleSplit := strings.Split(rule, ":")
		ruleIdx, _ := strconv.Atoi(ruleSplit[0])
		ruleData := ruleSplit[1]
		parsedRules[ruleIdx] = lexRule(strings.TrimSpace(ruleData))
	}
	return parsedRules
}

func matchRule(rules map[int]interface{}, currentRule interface{}, test string) (bool, string) {
	switch currentRule.(type) {
	case ExactRule:
		pass := strings.HasPrefix(test, currentRule.(ExactRule).character)
		if pass {
			return pass, test[len(currentRule.(ExactRule).character):]
		} else {
			return pass, ""
		}
	case AndRule:
		testAgainst := test
		var ok bool
		for _, ruleRef := range currentRule.(AndRule).refs {
			if ok, testAgainst = matchRule(rules, rules[ruleRef], testAgainst);
				!ok {
				return false, ""
			}
		}
		return true, testAgainst
	case OrRule:
		for _, option := range currentRule.(OrRule).options {
			if ok, remaining := matchRule(rules, option, test); ok {
				return true, remaining
			}
		}
		return false, ""

	default:
		panic("Unknown type")
	}
}

func part1(rules map[int]interface{}, tests []string) {
	answer := 0
	for _, test := range tests {
		match, rem := matchRule(rules, rules[0], test)
		//fmt.Printf("%s -> (%t, %s)\n", test, match, rem)
		fullMatch := match && len(rem) == 0
		if fullMatch {
			answer += 1
		}
	}
	println("Part 1:",answer)
}

func part2Rules(rules map[int]interface{}) map[int]interface{} {
	rules[8] = OrRule{options: []AndRule{{refs: []int{42, 8}}, {refs: []int{42}}}}
	rules[11] = OrRule{options: []AndRule{{refs: []int{42, 11, 31}}, {refs: []int{42, 31}}}}
	return rules
}

func part2(rules map[int]interface{}, tests []string) {
	rules = part2Rules(rules)
	answer := 0

	for _, test := range tests {
		//	need to match rule 0 = 8 11
		//	rule 8: 42+
		//	rule 11: 42 (42 31)* 31
		intermediateRemainders := oneOrMore(rules, rules[42], test)
		for _, remainder := range intermediateRemainders {
			finalRemainders := middleExpand(rules, rules[42], rules[31], remainder)
			finalMatch := make([]bool, len(finalRemainders))
			for i := range finalRemainders {
				finalMatch[i] = len(finalRemainders[i]) == 0
			}
			if any(finalMatch, true) {
				answer++
			}
		}

	}
	println("Part 2:",answer)
}

func oneOrMore(rules map[int]interface{}, repetitiveRule interface{}, test string) []string {
	//	rule 8: 42+
	result := []string{}
	for {
		matched, remainder := matchRule(rules, repetitiveRule, test)
		if matched {
			result = append(result, remainder)
			test = remainder
		} else {
			break
		}
	}
	return result
}

func middleExpand(rules map[int]interface{}, leftRule, rightRule interface{}, test string) []string {
	//	rule 11: 42 (42 31)* 31
	result := []string{}
Attempts:
	for repetitions := 1; repetitions <= 30; repetitions++ {
		var matched bool
		current := test
		for i := 1; i <= repetitions; i++ {
			matched, current = matchRule(rules, leftRule, current)
			if !matched {
				continue Attempts
			}
		}
		print()
		for j := 1; j <= repetitions; j++ {
			matched, current = matchRule(rules, rightRule, current)
			if !matched {
				continue Attempts
			}
		}
		result = append(result, current)
	}
	return result
}
