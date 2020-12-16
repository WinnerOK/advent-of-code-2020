package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input []string) (map[string][][2]int, []int, [][]int) {
	rangeRegexp := regexp.MustCompile(`(\d+)-(\d+)`)

	rules := map[string][][2]int{}
	myTicket := []int{}
	otherTickets := [][]int{}

	idx := 0
	for len(input[idx]) > 0 {
		lineSplit := strings.Split(input[idx], ":")
		ruleName := lineSplit[0]
		ruleRanges := lineSplit[1]
		ranges := rangeRegexp.FindAllStringSubmatch(ruleRanges, -1)
		parsedRanges := [][2]int{}
		for _, rangeDescriptor := range ranges {
			parsedMin, _ := strconv.Atoi(rangeDescriptor[1])
			parsedMax, _ := strconv.Atoi(rangeDescriptor[2])
			parsedRanges = append(parsedRanges, [2]int{parsedMin, parsedMax})
		}
		rules[ruleName] = parsedRanges
		idx += 1
	}
	idx += 2
	myTicket = stringSliceToIntSlice(strings.Split(input[idx], ","))

	idx += 3
	for idx < len(input) {
		otherTicket := stringSliceToIntSlice(strings.Split(input[idx], ","))
		otherTickets = append(otherTickets, otherTicket)
		idx += 1
	}
	return rules, myTicket, otherTickets
}

func main() {
	input := readInput("./in.txt")
	rules, myTicket, otherTickets := parseInput(input)
	part1Answer, validTickets := part1(rules, otherTickets)
	println("Part 1:", part1Answer)
	part2Answer := part2(rules, myTicket, validTickets)
	println("Part 2:", part2Answer)

}

func inRuleRange(number int, ruleRanges [][2]int) bool {
	passCurrentRule := []bool{}
	for _, ruleRange := range ruleRanges {
		min := ruleRange[0]
		max := ruleRange[1]
		passCurrentRule = append(
			passCurrentRule,
			min <= number && number <= max,
		)
	}
	return any(passCurrentRule, true)
}

func part1(rules map[string][][2]int, otherTickets [][]int) (int, [][]int) {
	answer := 0
	validTickets := [][]int{}
TicketLoop:
	for _, ticket := range otherTickets {
	NumberLoop:
		for _, number := range ticket {
			for _, ruleRanges := range rules {
				if inRuleRange(number, ruleRanges) {
					continue NumberLoop
				}
			}
			answer += number
			continue TicketLoop
		}
		validTickets = append(validTickets, ticket)
	}
	return answer, validTickets
}

type IdxLen struct {
	idx, length int
}

func getMapKeys(m map[string]bool) []string {
	answer := []string{}
	for k := range m {
		answer = append(answer, k)
	}
	return answer
}

func setRemove(from map[string]bool, rem map[string]int) map[string]bool {
	for k := range rem {
		if _, ok := from[k]; ok {
			delete(from, k)
		}
	}
	return from
}

func part2(rules map[string][][2]int, myTicket []int, validTickets [][]int) int {
	// map of [index in ticket]possible fields at given index
	rulesOrder := map[int]map[string]bool{}
	// map of all rules
	allRules := map[string]bool{}
	// fill rulesOrder map
	for ruleName, _ := range rules {
		allRules[ruleName] = true
	}
	for i := 0; i < len(rules); i++ {
		rulesOrder[i] = map[string]bool{}
		for k, v := range allRules {
			rulesOrder[i][k] = v
		}
	}

	// For each ticket's number
	for _, ticket := range validTickets {
		for numberIdx, number := range ticket {
			// if current position is impossible to follow a rule
			// remove the rule from possibilities on this position
			for ruleName, ruleRanges := range rules {
				if !inRuleRange(number, ruleRanges) {
					delete(rulesOrder[numberIdx], ruleName)
				}
			}
		}
	}

	// Get indexes of possible maps: first - most certain, last - least certain
	// so lengths[0] after the sort will point to map with the only element
	lengths := []IdxLen{}
	for mapIdx, possibleFieldsMap := range rulesOrder {
		lengths = append(lengths, IdxLen{
			idx:    mapIdx,
			length: len(possibleFieldsMap),
		})
	}
	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i].length < lengths[j].length
	})

	// fields that are surely resolved
	resolvedFields := map[string]int{}
	resolvedFields[getMapKeys(rulesOrder[lengths[0].idx])[0]] = lengths[0].idx

	for i := 1; i < len(lengths); i++ {
		// from each following map remove all surely resolved keys
		// only 1 key should be left then
		remaining := setRemove(rulesOrder[lengths[i].idx], resolvedFields)
		if len(remaining) > 1 {
			panic("could not determine length")
		}
		resolvedFields[getMapKeys(remaining)[0]] = lengths[i].idx
	}

	answer := 1
	for ruleName := range rules{
		if strings.HasPrefix(ruleName, "departure"){
			answer *= myTicket[resolvedFields[ruleName]]
		}
	}
	return answer
}
