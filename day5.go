package main

import (
    "fmt"
    "strings"
    "strconv"
    "os"
    "bufio"
)

type rule struct {
    first int;
    second int
}

func day5() {
    f, err := os.Open("input5")
    if err != nil {
        panic("Can't open file")
    }

    scanner := bufio.NewScanner(bufio.NewReader(f))
    parse_rules := true

    rules := make([]rule, 0)

    sum := 0
    incorr_sum := 0

    for scanner.Scan() {
        text := scanner.Text()
        if text == "" {
            parse_rules = false
            continue
        }

        if parse_rules {
            parts := strings.Split(text, "|")
            num_first,err := strconv.Atoi(parts[0])
            num_second,err2 := strconv.Atoi(parts[1])

            if err != nil && err2 != nil {
                panic("Error parsing numbers")
            }

            rules = append(rules, rule{first:num_first, second:num_second})
        } else {
            parts := strings.Split(text, ",")
            sequence := make([]int, 0)
            for i := 0; i < len(parts); i++ {
                conv,err := strconv.Atoi(parts[i])
                if err != nil {
                    panic("Parse failed")
                }
                sequence = append(sequence, conv)
            }
            filtered := filterRules(sequence,rules)
            if satisfiesRules(sequence, filtered) {
                num := sequence[(len(sequence)-1)/2]
                sum += num
            } else {
                fixedSequence := make([]int, len(sequence))
                copy(fixedSequence, sequence)
                for !satisfiesRules(fixedSequence, filtered) {
                    fmt.Println(fixedSequence)
                    fixOrdering(fixedSequence, filtered)
                }
                num := fixedSequence[(len(fixedSequence)-1)/2]
                incorr_sum += num
            }
        }

    }
    fmt.Println("Correct pages: ", sum)
    fmt.Println("Incorrect pages: ", incorr_sum)
}

func filterRules(sequence []int, rules []rule) (applicable_rules []rule) {
    for i := 0; i < len(rules); i++ {
        for j := 0; j < len(sequence); j++ {
            if sequence[j] == rules[i].first {
                for k := j+1; k < len(sequence); k++ {
                    if sequence[k] == rules[i].second {
                        applicable_rules = append(applicable_rules, rules[i])
                        break
                    }
                }
            } else if sequence[j] == rules[i].second {
                for k := j+1; k < len(sequence); k++ {
                    if sequence[k] == rules[i].first {
                        applicable_rules = append(applicable_rules, rules[i])
                        break
                    }
                }
            }
        }
    }
    return applicable_rules
}

func satisfiesRules(sequence []int, rules []rule) bool {
    satisfied := true
    for i := 0; i < len(rules); i++ {
        satisfied = satisfied && satisfiesRule(sequence, rules[i])
    }
    return satisfied
}

func satisfiesRule(sequence []int, condition rule) bool {
    idx_first, idx_second := -1, -1
    for i := 0; i < len(sequence); i++ {
        if sequence[i] == condition.first {
            idx_first = i
        } else if sequence[i] == condition.second {
            idx_second = i
        }
    }
    assert(idx_first != -1, "Number not found")
    assert(idx_second != -1, "Second number not found")
    return idx_first < idx_second
}

func fixOrdering(sequence []int, rules []rule) {
    unsatisfied := -1
    for i := 0; i < len(rules); i++ {
        if !satisfiesRule(sequence, rules[i]) {
            unsatisfied = i
            break
        }
    }
    assert(unsatisfied != -1, "Unsatisfied index out of range")
    rule_to_fix := rules[unsatisfied]
    
    idx_first, idx_second := -1, -1

    for i := 0; i < len(sequence); i++ {
        if rule_to_fix.first == sequence[i] {
            idx_first = i
        }
        if rule_to_fix.second == sequence[i] {
            idx_second = i
        }
    }

    assert(idx_first != -1, "Number not found, fixing")
    assert(idx_second != -1, "Number 2 not found, fixing")

    sequence[idx_first] = rule_to_fix.second
    sequence[idx_second] = rule_to_fix.first
    //TODO fix the rule (move bad one right in front of where it sould be.. or swap)
}
