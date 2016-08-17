package main

import "errors"

type state struct {
	State string `json:"state"`
}

func (s *state) validate() error {
	if s.State != relationShip_Request_State_Like && s.State != relationShip_Request_State_Dislike {
		return ErrInvalidateParamState
	}

	return nil
}

var (
	relationShip_Request_State_Like    string = "liked"
	relationShip_Request_State_Dislike string = "disliked"
	relationShip_Request_State_Match   string = "match"
)

var ErrInvalidateParamState = errors.New("invalid param state")
