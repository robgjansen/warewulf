package warewulfd

import (
	"net/http"

	"github.com/hpcng/warewulf/internal/pkg/kernel"
)

func KmodsSend(w http.ResponseWriter, req *http.Request) {

	node, err := getSanity(req)
	if err != nil {
		w.WriteHeader(404)
		daemonLogf("ERROR: %s\n", err)
		return
	}

	if node.KernelVersion.Defined() {
		fileName := kernel.KmodsImage(node.KernelVersion.Get())

		err := sendFile(w, fileName, node.Id.Get())
		if err != nil {
			daemonLogf("ERROR: %s\n", err)
		}

	} else {
		w.WriteHeader(503)
		daemonLogf("WARNING: No 'kernel version' set for node %s\n", node.Id.Get())
	}
}
