package vote

import (
	"fmt"
	"github.com/1851616111/util/validator/mobile"
	"github.com/1851616111/util/validator/name"
	"github.com/1851616111/util/validator/tel"
)

func ParamNotFoundError(param string) error {
	return fmt.Errorf("param %s not found", param)
}

type Voter struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Company     string `json:"company"`
	Mobile      string `json:"mobile"`
	Declaration string `json:"declaration"`
	VotedCount  int    `json:"voteCount"`
}

func (v Voter) Validate() error {
	if err := name.Validate(v.Name); err != nil {
		return ParamNotFoundError("name")
	}

	if len(v.Image) == 0 {
		return ParamNotFoundError("image")
	}

	err1, err2 := mobile.Validate(v.Mobile), tel.Validate(v.Mobile)
	if err1 != nil && err2 != nil {
		return ParamNotFoundError("mobile")
	}

	if len(v.Company) == 0 {
		return ParamNotFoundError("company")
	}

	if len(v.Declaration) == 0 {
		return ParamNotFoundError("decalration")
	}

	return nil
}

func (v *Voter) Complete() {
	v.VotedCount = 0
	v.ID = ""
}
