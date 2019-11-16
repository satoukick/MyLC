package main

func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	regionMap := make(map[string]string)
	for _, sli := range regions {
		root := sli[0]
		for i := 1; i < len(sli); i++ {
			regionMap[sli[i]] = root
		}
	}

	chain1 := []string{}
	chain2 := []string{}

	for region1 != "" {
		chain1 = append(chain1, region1)
		region1 = regionMap[region1]
	}
	for region2 != "" {
		chain2 = append(chain2, region2)
		region2 = regionMap[region2]
	}

	for _, v1 := range chain1 {
		for _, v2 := range chain2 {
			if v1 == v2 {
				return v1
			}
		}
	}
	return ""
}
