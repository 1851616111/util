package relationship

type relation struct {
	From_id string
	To_id   string
}

type RelationShip struct {
	UserID string `json:"user_id"`
	State  string `json:"state"`
}

type RelationShipList []*RelationShip
