package beat-utils

func List2Set (ss []string) ( mm map[string]struct{}) {
    // converts lists to sort of sets
    // main use is to avoid having set literals in text configs
    
    mm = make ( map[string]struct{}, len(ss) )
    for _, s := range ss {
        mm[s] = struct{}{}
    }
    return mm
}
