package main

import (
	"encoding/json"
	"errors"
	idapi "github.com/1851616111/tantan-test/idcontroller/client"
	relationshipapi "github.com/1851616111/tantan-test/relationship"
	"github.com/1851616111/tantan-test/user"
	"github.com/pivotal-golang/lager"
	"net/http"
)

//curl -XPOST -d '{"name":"Alice"}' "http://localhost:8080/users"
func createUserHandler(w http.ResponseWriter, r *http.Request, log lager.Logger, vars map[string]string) {
	defer r.Body.Close()

	user := new(user.User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	id, err := idapi.GenUserID(IDControllerCfg.Host, IDControllerCfg.Port)
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}
	user.Id = id

	if err := userClient.Create(user); err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	respond(w, http.StatusOK, user)
}

//curl -XGET "http://localhost:8080/users"
func listUsersHandler(w http.ResponseWriter, r *http.Request, log lager.Logger, vars map[string]string) {
	l, err := userClient.List(nil)
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	respond(w, http.StatusOK, l)
}

//curl -XPUT -d '{"state":"liked"}' "http://localhost:8080/users/10000000000/relationships/10000000001"
//{
//"user_id": "21341231231",
//"state": "liked" ,
//"type": "relationship"
//}

//$curl -XPUT -d '{"state":"liked"}' "http://localhost:8080/users/10000000001/relationships/10000000000"
//{
//"user_id": "11231244213",
//"state": "matched" ,
//"type": "relationship"
//}

//$curl -XPUT -d '{"state":"disliked"}' "http://localhost:8080/users/10000000001/relationships/10000000000"
//{
//"user_id": "11231244213",
//"state": "disliked" ,
//"type": "relationship"
//}

//curl -XPUT -d '{"state":"disliked"}' "http://localhost:8080/users/10000000000/relationships/10000000001"
func createRelationShipHandler(w http.ResponseWriter, r *http.Request, log lager.Logger, vars map[string]string) {
	user_id, other_user_id := vars["user_id"], vars["other_user_id"]
	if len(other_user_id) == 0 {
		respond(w, http.StatusBadRequest, errors.New("param other_user_id must not be nil."))
		return
	}

	state := new(state)
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(state); err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}
	if err := state.validate(); err != nil {
		respond(w, http.StatusBadRequest, err)
		return
	}

	ret := &relationshipapi.RelationShip{
		UserID: other_user_id,
		State:  state.State,
	}

	switch state.State {
	case relationShip_Request_State_Dislike:
		if err := disLikeLogic(user_id, other_user_id, relationShipClient, log); err != nil {
			respond(w, http.StatusInternalServerError, err)
			return
		}

	case relationShip_Request_State_Like:
		ifMatch, err := relationShipClient.Match().IfMatch(user_id, other_user_id)
		if err != nil {
			respond(w, http.StatusInternalServerError, err)
			return
		}

		if ifMatch {
			ret.State = relationShip_Request_State_Match
			respond(w, http.StatusOK, ret)
			return
		}

		ifLike, err := relationShipClient.Like().IfLike(other_user_id, user_id)
		if err != nil {
			respond(w, http.StatusInternalServerError, err)
			return
		}
		switch ifLike {
		case true:
			ret.State = relationShip_Request_State_Match
			relationShipClient.Match().Create(user_id, other_user_id)
			relationShipClient.Match().Create(other_user_id, user_id)
			relationShipClient.Like().Delete(other_user_id, user_id)

		case false:
			relationShipClient.Dislike().Delete(user_id, other_user_id) //maybe like
			relationShipClient.Like().Create(user_id, other_user_id)
		}

	}

	respond(w, http.StatusOK, ret)
}

//curl -XGET "http://localhost:8080/users/10000000001/relationships"
func listRelationShipsHandler(w http.ResponseWriter, r *http.Request, log lager.Logger, vars map[string]string) {
	user_id := vars["user_id"]

	var err error
	var likeList, dislikeList, matchList *relationshipapi.RelationShipList

	likeList, err = relationShipClient.Like().List(user_id)
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	dislikeList, err = relationShipClient.Dislike().List(user_id)
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	matchList, err = relationShipClient.Match().List(user_id)
	if err != nil {
		respond(w, http.StatusInternalServerError, err)
		return
	}

	respond(w, http.StatusOK, relationshipapi.Merge(likeList, dislikeList, matchList))
}

func disLikeLogic(user_id, other_user_id string, cli relationshipapi.RelationShipInterface, log lager.Logger) error {
	ifMatch, err := cli.Match().IfMatch(user_id, other_user_id)
	if err != nil {
		return err
	}

	switch ifMatch {
	case true:
		cli.Match().Delete(user_id, other_user_id)
		cli.Match().Delete(other_user_id, user_id)
		cli.Dislike().Create(user_id, other_user_id)
		cli.Like().Create(other_user_id, user_id)

	case false:
		cli.Like().Delete(user_id, other_user_id) //maybe like
		cli.Dislike().Create(user_id, other_user_id)
	}

	return nil
}
