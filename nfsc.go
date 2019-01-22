package beat_utils

import (
	//"bufio"
	"bytes"
	//"fmt"
	"io/ioutil"
	"log" // how do I do logging from beats? May be their logp?
	//"path/filepath"
	"strconv"
	//"strings"
)

/*
License go here. The code will  more or less converted from python-Diamond 
*/

PROC = '/proc/net/rpc/nfs'

func GetNFSClientStats(s string) (result map[string]uint64, err error) {
	// Reads /proc/net/rpc/nfs and return 

	buf, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(buf))
	// first line is something ZFS specific with a timestamp, second is names, third is valus and fourth empty
	// using splits for just two lines?
	bb := bytes.Split(buf, []byte("\n"))
	
	for n, line := bb {
		fmt.Println(n, string(line)
	}
	result = make(map[string]uint64)		    
	/*if len(bb) > 2 {
		line := bytes.Fields(bb[1])
		svals := bytes.Fields(bb[2])
		l := len(keys)
		if len(svals) != l {
			log.Fatal("Number of fields mismatch", keys, svals)
		}
		result = make(map[string]uint64)
		for i, k := range keys {
			v, err := strconv.ParseUint(string(svals[i]), 10, 64)
			if err != nil {
				log.Fatal("Parse failure in zfs io", err)
			}
			result[string(k)] = v
		}
		fmt.Println(result)
	} else {
		log.Fatal("strange io file, len", len(bb))
	}
	*/
	return result, nil
}
