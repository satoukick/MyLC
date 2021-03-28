package main

func numDifferentIntegers(word string) int {
	nums := []string{}
	cur := ""
loop:
	for _, s := range word {
		if s >= '0' && s <= '9' {
			if len(cur) > 0 && cur[0] == '0' {
				cur = ""
			}
			cur += string(s)
		} else {
			if len(cur) > 0 {
				str := cur
				cur = ""
				for _, v := range nums {
					if v == str {
						continue loop
					}
				}

				nums = append(nums, str)
			}
		}
	}

	if cur != "" {
		flag := true
		for _, v := range nums {
			if v == cur {
				flag = false
				break
			}
		}
		if flag {
			nums = append(nums, cur)
		}
	}
	// fmt.Println(nums)
	return len(nums)
}

func main() {
	numDifferentIntegers("uu717761265362523668772526716127260222200144937319826j717761265362523668772526716127260222200144937319826k717761265362523668772526716127260222200144937319826b7177612653625236687725267161272602222001449373198262hgb9utohfvfrxprqva3y5cdfdudf7zuh64mobfr6mhy17")
}
