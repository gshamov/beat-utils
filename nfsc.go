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
			fmt.Println(string( line[0]), string(line[1]), mylen)
			if ( len(line) != (mylen + 2) ) {
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
			fmt.Println(string(line[0]), string(line[1]), mylen)
			if len(line) != (mylen + 2) {
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
			mylen := int( atoi(line[1]) )
			fmt.Println(string(line[0]), string(line[1]), mylen)
/*			if len(ll) != (mylen + 2) {
				// something wrong
				// mylen should be 48 for proc4
				// or at least 30
				break
			} */
		    result["v4.null"] = atoi(line[2])
                    result["v4.read"] = atoi(line[3])
                    result["v4.write"] = atoi(line[4])
                    result["v4.commit"] = atoi(line[5])
                    result["v4.open"] = atoi(line[6])
                    result["v4.open_conf"] = atoi(line[7])
                    result["v4.open_noat"] = atoi(line[8])
                    result["v4.open_dgrd"] = atoi(line[9])
                    result["v4.close"] = atoi(line[10])
                    result["v4.setattr"] = atoi(line[11])
                    result["v4.fsinfo"] = atoi(line[12])
                    result["v4.renew"] = atoi(line[13])
                    result["v4.setclntid"] = atoi(line[14])
                    result["v4.confirm"] = atoi(line[15])
                    result["v4.lock"] = atoi(line[16])
                    result["v4.lockt"] = atoi(line[17])
                    result["v4.locku"] = atoi(line[18])
                    result["v4.access"] = atoi(line[19])
                    result["v4.getattr"] = atoi(line[20])
                    result["v4.lookup"] = atoi(line[21])
                    result["v4.lookup_root"] = atoi(line[22])
                    result["v4.remove"] = atoi(line[23])
                    result["v4.rename"] = atoi(line[24])
                    result["v4.link"] = atoi(line[25])
                    result["v4.symlink"] = atoi(line[26])
                    result["v4.create"] = atoi(line[27])
                    result["v4.pathconf"] = atoi(line[28])
                    result["v4.statfs"] = atoi(line[29])
                    result["v4.readlink"] = atoi(line[30])
		    result["v4.readdir"] = atoi(line[31])
		    if mylen > 34 {
			result["v4.getacl"] = atoi(line[35])
			result["v4.setacl"] = atoi(line[36])
		    }
		    // do we need the rest of them?
		default:
			// NFS5?
			break
		}
	}

	return result, nil
}

func GetNFSClientStats1(s string) (result map[string]map[string]uint64, err error) {
	// Reads /proc/net/rpc/nfs and returns map of maps, each inner map corresponds to a line
	// this is to avoid dots in metric names?
	
	s = "/proc/net/rpc/nfs"

	buf, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(buf))

	bb := bytes.Split(buf, []byte("\n"))

	result = make(map[string]map[string]uint64)

	for n, ll := range bb {
		fmt.Println(n, string(ll))
		line := bytes.Fields(ll)
		if len(line) < 1 {
			break
		}
		switch string(line[0]) {
		case "net":
			mnet := map[string]uint64 {
			"packets": atoi(line[1]),
			"udpcnt":  atoi(line[2]),
			"tcpcnt":  atoi(line[3]),
			"tcpconn":  atoi(line[4]),
			}
			result["net"] = mnet
		case "rpc":
			mrpc := map[string]uint64 {
			"calls":  atoi(line[1]),
			"retrans":  atoi(line[2]),
			"authrefrsh":  atoi(line[3]),
			}
			result["rpc"] = mrpc
			
		case "proc2":
			mylen := int(atoi(line[1]))
			fmt.Println(string( line[0]), string(line[1]), mylen)
			if ( len(line) != (mylen + 2) ) {
				// something wrong
				// mylen should be 18 for proc2
				break
			}
			mproc2 := map[string]uint64 {
			"null": atoi(line[2]),
			"getattr": atoi(line[3]),
			"setattr": atoi(line[4]),
			"root": atoi(line[5]),
			"lookup": atoi(line[6]),
			"readlink": atoi(line[7]),
			"read": atoi(line[8]),
			"wrcache": atoi(line[9]),
			"write": atoi(line[10]),
			"create": atoi(line[11]),
			"remove": atoi(line[12]),
			"rename": atoi(line[13]),
			"link": atoi(line[14]),
			"symlink": atoi(line[15]),
			"mkdir": atoi(line[16]),
			"rmdir": atoi(line[17]),
			"readdir": atoi(line[18]),
			"fsstat": atoi(line[19]),
			}
			result["v2"] = mproc2
		case "proc3":
			mylen := int( atoi(line[1]) )
			fmt.Println(string(line[0]), string(line[1]), mylen)
			if len(line) != (mylen + 2) {
				// something wrong
				// mylen should be 22 for proc3
				break
			}
			mproc3 := map[string]uint64 {
			"null": atoi(line[2]),
			"getattr": atoi(line[3]),
			"setattr": atoi(line[4]),
			"lookup": atoi(line[5]),
			"access": atoi(line[6]),
			"readlink": atoi(line[7]),
			"read": atoi(line[8]),
			"write": atoi(line[9]),
			"create": atoi(line[10]),
			"mkdir": atoi(line[11]),
			"symlink": atoi(line[12]),
			"mknod": atoi(line[13]),
			"remove": atoi(line[14]),
			"rmdir": atoi(line[15]),
			"rename": atoi(line[16]),
			"link": atoi(line[17]),
			"readdir": atoi(line[18]),
			"readdirplus": atoi(line[19]),
			"fsstat": atoi(line[20]),
			"fsinfo": atoi(line[21]),
			"pathconf": atoi(line[22]),
			"commit": atoi(line[23]),
			}
			result["v3"] = mproc3
		case "proc4":
			mylen := int( atoi(line[1]) )
			fmt.Println(string(line[0]), string(line[1]), mylen)
			if len(ll) != (mylen + 2) {
				// something wrong
				// mylen should be 48 for proc4
				// or at least 30
				break
			}
			mproc4 := map[string]uint64{
		    "null": atoi(line[2]),
                    "read": atoi(line[3]),
                    "write": atoi(line[4]),
                    "commit": atoi(line[5]),
                    "open": atoi(line[6]),
                    "open_conf": atoi(line[7]),
                    "open_noat": atoi(line[8]),
                    "open_dgrd": atoi(line[9]),
                    "close": atoi(line[10]),
                    "setattr": atoi(line[11]),
                    "fsinfo": atoi(line[12]),
                    "renew": atoi(line[13]),
                    "setclntid": atoi(line[14]),
                    "confirm": atoi(line[15]),
                    "lock": atoi(line[16]),
                    "lockt": atoi(line[17]),
                    "locku": atoi(line[18]),
                    "access": atoi(line[19]),
                    "getattr": atoi(line[20]),
                    "lookup": atoi(line[21]),
                    "lookup_root": atoi(line[22]),
                    "remove": atoi(line[23]),
                    "rename": atoi(line[24]),
                    "link": atoi(line[25]),
                    "symlink": atoi(line[26]),
                    "create": atoi(line[27]),
                    "pathconf": atoi(line[28]),
                    "statfs": atoi(line[29]),
                    "readlink": atoi(line[30]),
		    "readdir": atoi(line[31]),
		    //if mylen > 34 {
			"getacl": atoi(line[35]),
			"setacl": atoi(line[36]),
		    // do we need the rest of them?
		    }
		    result["v4"] = mproc4
		default:
			// NFSv5???
			break
		}
	}

	return result, nil
}
