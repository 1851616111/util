package relationship

import (
	"github.com/1851616111/tantan-test/db"
)

const (
	key_dislike = "relationship_dislike"
	key_match   = "relationship_match"
)

type RelationShipInterface interface {
	LikesInterface
	DislikesInterface
	MatchesInterface
}

func NewRelationShipClient(db *db.DB) RelationShipInterface {
	return &impl{
		db,
	}
}

type impl struct {
	*db.DB
}

func (i *impl) Like() LikeInterface {
	return &likeImpl{i}
}

func (i *impl) Dislike() DislikeInterface {
	return &disLikeImpl{i}
}

func (i *impl) Match() MatchInterface {
	return &matchImpl{i}
}
