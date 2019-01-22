package beat_utils

import (
	//"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log" // how do I do logging from beats? May be their logp?
	//"path/filepath"
	"strconv"
	//"strings"
)

/*
License go here. The code will  more or less converted from python-Diamond
*/

func atoi(bb []byte) uint64 {
	ii, err := strconv.ParseUint(string(bb),10,64)
	// unsafe function that returns zero on parse failure
	if err != nil {
		return 0
	}
	return ii
}

func GetNFSClientStats(s string) (result map[string]uint64, err error) {
	// Reads /proc/net/rpc/nfs and return
	s = "/proc/net/rpc/nfs"

	buf, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(buf))

	bb := bytes.Split(buf, []byte("\n"))

	result = make(map[string]uint64)

	for n, ll := range bb {
		fmt.Println(n, string(ll))
		line := bytes.Fields(ll)
		if len(line) < 1 {
			break
		}
		switch string(line[0]) {
		case "net":
			result["net.packets"] = atoi(line[1])
			result["net.udpcnt"] = atoi(line[2])
			result["net.tcpcnt"] = atoi(line[3])
			result["net.tcpconn"] = atoi(line[4])
		case "rpc":
			result["rpc.calls"] = atoi(line[1])
			result["rpc.retrans"] = atoi(line[2])
			result["rpc.authrefrsh"] = atoi(line[3])
		case "proc2":
			mylen := int(atoi(line[1]))
			if ( len(ll) != (mylen + 2) ) || (mylen != 18) {
				// something wrong
				// mylen should be 18 for proc2
				break
			}
			result["v2.null"] = atoi(line[2])
			result["v2.getattr"] = atoi(line[3])
			result["v2.setattr"] = atoi(line[4])
			result["v2.root"] = atoi(line[5])
			result["v2.lookup"] = atoi(line[6])
			result["v2.readlink"] = atoi(line[7])
			result["v2.read"] = atoi(line[8])
			result["v2.wrcache"] = atoi(line[9])
			result["v2.write"] = atoi(line[10])
			result["v2.create"] = atoi(line[11])
			result["v2.remove"] = atoi(line[12])
			result["v2.rename"] = atoi(line[13])
			result["v2.link"] = atoi(line[14])
			result["v2.symlink"] = atoi(line[15])
			result["v2.mkdir"] = atoi(line[16])
			result["v2.rmdir"] = atoi(line[17])
			result["v2.readdir"] = atoi(line[18])
			result["v2.fsstat"] = atoi(line[19])

		case "proc3":
			mylen := int( atoi(line[1]) )
			if len(ll) != (mylen + 2) {
				// something wrong
				// mylen should be 22 for proc3
				break
			}
			result["v3.null"] = atoi(line[2])
			result["v3.getattr"] = atoi(line[3])
			result["v3.setattr"] = atoi(line[4])
			result["v3.lookup"] = atoi(line[5])
			result["v3.access"] = atoi(line[6])
			result["v3.readlink"] = atoi(line[7])
			result["v3.read"] = atoi(line[8])
			result["v3.write"] = atoi(line[9])
			result["v3.create"] = atoi(line[10])
			result["v3.mkdir"] = atoi(line[11])
			result["v3.symlink"] = atoi(line[12])
			result["v3.mknod"] = atoi(line[13])
			result["v3.remove"] = atoi(line[14])
			result["v3.rmdir"] = atoi(line[15])
			result["v3.rename"] = atoi(line[16])
			result["v3.link"] = atoi(line[17])
			result["v3.readdir"] = atoi(line[18])
			result["v3.readdirplus"] = atoi(line[19])
			result["v3.fsstat"] = atoi(line[20])
			result["v3.fsinfo"] = atoi(line[21])
			result["v3.pathconf"] = atoi(line[22])
			result["v3.commit"] = atoi(line[23])
		case "proc4":

		default:
			// NFS5?
			break
		}
	}

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
