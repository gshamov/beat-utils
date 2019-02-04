package beat_utils

import (
  "strings"
  "path/filepath"
)

func MakeCgroupTags(s string) (uid, job, step string, ok bool) {
        // takes /slurm/uid_3052932/job_11354355/step_batch and returns the components
        // match already checked upstream? can be defensive here though..
        
        pattern := "/slurm/uid_*/job_*/step_*"
        match, _ := filepath.Match(pattern, s)
        if !match {
                return "", "", "", false
        }
        uid = ""
        job = ""
        step = ""
     
        ll := strings.Split(s, "/")
        //fmt.Println(ll)
   
        uid = strings.Replace(ll[2], "uid_", "", 1)
        job = ll[3][4:]
        step = ll[4][4:] // this will keep underscore like _7, _batch
 
      return uid, job, step, true
}
