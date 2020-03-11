package routes

import (
	"github.com/gorilla/mux"
)

// SetUpRoutes sets up routes for jorm
func Routes() *mux.Router {
	r := mux.NewRouter()

	r.Path("/").HandlerFunc(hnd.IndexHandler).Name("index")

	//e := r.PathPrefix("/e").Subrouter()
	//rtse.ExplorerRoutes(e)
	//
	//c := r.PathPrefix("/c").Subrouter()
	//rtsc.ComUsRoutes(c)
	//
	//s := r.PathPrefix("/s").Subrouter()
	//rtss.StatusRoutes(s)
	//
	//pif := r.PathPrefix("/pif").Subrouter()
	//rtspif.ParallelCoinInfoRoutes(pif)
	//
	//pe := r.PathPrefix("/pe").Subrouter()
	//rtspe.ParallelExplorerRoutes(pe)
	//
	//pio := r.PathPrefix("/pio").Subrouter()
	//rtspio.ParallelCoinIoRoutes(pio)

	v := r.PathPrefix("/v").Subrouter()
	rtsv.VesicaRoutes(v)

	// b := r.PathPrefix("/b").Subrouter()
	// rtsc.ComUsRoutes(b)

	// XXX := r.PathPrefix("/XXX").Subrouter()
	// XXX := r.PathPrefix("/XXX").Subrouter()
	// XXX := r.PathPrefix("/XXX").Subrouter()
	// XXX := r.PathPrefix("/XXX").Subrouter()
	// XXX := r.PathPrefix("/XXX").Subrouter()
	// XXX := r.PathPrefix("/XXX").Subrouter()
	// XXX := r.PathPrefix("/XXX").Subrouter()
	// XXX := r.PathPrefix("/XXX").Subrouter()
	// XXX := r.PathPrefix("/XXX").Subrouter()

	r.Schemes("https")
	return r
}
