package flights

// space complexity O(n)
// time complexity O(n)
func GetSrcDestFromFlights(list [][]string) []string {
	// bad input
	if len(list) == 0 {
		return []string{}
	}

	// create a map to store the source and destination
	srcDestMap := make(map[string]string)
	// create a map to store the destination and source
	destSrcMap := make(map[string]string)
	for _, arr := range list {
		k := arr[0]
		v := arr[1]
		srcDestMap[k] = v
		destSrcMap[v] = k
	}

	ret := list[0]
	// find source before current position
	for {
		dest := ret[0]
		src, ok := destSrcMap[dest]
		if !ok {
			break
		}
		// update src to current position
		ret[0] = src
		// avoid cycles
		if ret[0] == ret[1] {
			return []string{}
		}
	}
	// find destination after current position
	for {
		src := ret[1]
		dest, ok := srcDestMap[src]
		if !ok {
			break
		}
		// update dest to current position
		ret[1] = dest
		// avoid cycles
		if ret[0] == ret[1] {
			return []string{}
		}
	}

	return ret
}
