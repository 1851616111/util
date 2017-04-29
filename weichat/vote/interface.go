package vote

type DBInterface interface {
	Init() error
	Register(openid string, voter *Voter) error
	Vote(openid, votedID string) error
	ListVoters(index, size int) ([]Voter, error)
}
