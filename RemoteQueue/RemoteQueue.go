package RemoteQueue

import (  
    "fmt"
)

type RemoteQueue struct {  
    JOB_DONE 		string
    JOB_RUNNING 	string
    JOB_INCOMPLETE 	string
	JOB_CREATED 	string
	
	destination		string
	setup_runtime	string
	need_runtime 	string
    queue 			string
    _queue 			string
    _ssh 			string

	// JOB_DONE = 'finished'
    // JOB_RUNNING = 'running'
    // JOB_INCOMPLETE = 'incomplete'
    // JOB_CREATED = 'created'

	// PROTOCOL_VERSION = 0, 2
}

func init() {

	

}