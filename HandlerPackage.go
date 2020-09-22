package libapi

import (
	"encoding/json"
	"github.com/phalpin/liberr"
	"github.com/phalpin/liberr/errortype"
	"net/http"
)

type HandlerPackage struct {
	Path    string
	Name    string
	Handler func(*http.Request) (*ActionResult, error)
	Methods []string
}

func HandlerRecord(name string, path string, handlerFunc func(*http.Request) (*ActionResult, error), methods ...string) *HandlerPackage {
	pkg := &HandlerPackage{
		Path:    path,
		Name:    name,
		Handler: handlerFunc,
		Methods: methods,
	}

	return pkg
}

func (hp *HandlerPackage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res, err := hp.Handler(r)
	if err != nil {
		WriteErrorResponse(w, err)
	}

	if res != nil {
		if res.Result != nil {
			encObj, marshalErr := json.Marshal(res.Result)
			if marshalErr != nil {
				WriteErrorResponse(w, liberr.NewBase("fatal error occurred, please try again later", liberr.WithErrorType(errortype.Unknown)))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(encObj)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
