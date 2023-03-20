package concurrence

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func RunConcurrenceExample() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Voce é o visitante número: %d\n", number)))
	})

	http.ListenAndServe(":3000", nil)
}
