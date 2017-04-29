package vote

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fmt"
	httputil "github.com/1851616111/util/http"
	"github.com/julienschmidt/httprouter"
	"path/filepath"
)

var dbI DBInterface

const DEFAULT_PAGE_SIZE = "20"

func NewRouter() *httprouter.Router {
	r := httprouter.New()

	r.GET("/api/activity/voters", ListVotersHandler)
	r.POST("/api/activity/voter", RegisterVoterHandler)
	r.POST("/api/activity/voter/:id/vote", VoteHandler)

	dist, err := filepath.Abs("./activity_imgs")
	if err != nil {
		panic(err)
	}
	r.ServeFiles("/api/activity/file/*filepath", http.Dir(dist))
	return r
}

func RegisterVoterHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	openid := ps.ByName("openid")
	v := &Voter{}
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	if err := v.Validate(); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	v.Complete()

	if err := dbI.Register(openid, v); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.Response(w, 200, "ok")
}

func VoteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	openid := ps.ByName("openid")
	voteID := ps.ByName("id")
	if err := dbI.Vote(openid, voteID); err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.Response(w, 200, "ok")
}

func ListVotersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	indexS, sizeS := r.FormValue("index"), r.FormValue("size")
	if len(indexS) == 0 {
		indexS = "1"
	}

	if len(sizeS) == 0 {
		sizeS = DEFAULT_PAGE_SIZE
	}

	index, err := strconv.Atoi(indexS)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	size, err := strconv.Atoi(sizeS)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	l, err := dbI.ListVoters(index, size)
	if err != nil {
		httputil.Response(w, 400, err)
		return
	}

	httputil.ResponseJson(w, 200, l)
	return
}
