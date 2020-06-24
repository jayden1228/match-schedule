package core

// CombinationSum 从数组中选中N个数求和，满足 candidates[m] + ... candidates[n] = target
// @candidates 元素数组
// @target     求和
// @curDep     当前搜索深度
// @tarDep     目标搜索深度
func CombinationSum(candidates []int32, target int32, curDeep int32, tarDeep int32) [][]int32 {
	comb := make([][]int32, 0)
	if curDeep <= tarDeep {
		for i := range candidates {
			if candidates[i] == target {
				if curDeep == tarDeep {
					comb = append(comb, []int32{candidates[i]})
					break
				}
			} else if candidates[i] < target {
				rtn := CombinationSum(candidates[i:], target-candidates[i], curDeep+1, tarDeep)
				for j := range rtn {
					if len(rtn[j]) == 0 {
						continue
					}
					comb = append(comb, append([]int32{candidates[i]}, rtn[j]...))
				}
			} else {
				break
			}
		}
	}

	return comb
}
