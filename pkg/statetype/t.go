package statetype

type T int

const (
	Rolling T = iota
	DiscardingResources
	MovingRobber
	RobbingNeighboringCommunity
	NormalGameplay
)
