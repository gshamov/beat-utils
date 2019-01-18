package beat_utils

func List2Set (ss []string) ( mm map[string]struct{}) {
    // converts lists to sort of sets
    // main use is to avoid having set literals in text configs
    
    mm = make ( map[string]struct{}, len(ss) )
    for _, s := range ss {
        mm[s] = struct{}{}
    }
    return mm
}

func DropKVs(list []string, mm map[string]interface{} ) {
    // drops key-value pairs from map mm if key is not present in the list 
    // the use of this function is to prune unwanted metrics without much thinking
        metrics := List2Set(list)
        var keys []string
    // making list of keys to iterate over (deletion while iterating looks dangerous)
        for k := range (mm) {
               keys = append (keys, k)
        }
        for _, k := range(keys) {
                        if  _, found := metrics[k] ; ! found {
                            delete(mm, k)
                }
         }
}
