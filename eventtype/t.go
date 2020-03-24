package eventtype

type T int

const (
	// Turn state
	RolledNumber T = iota
	EndedTurn

	// The robber
	ResourcesDiscarded
	PlacedRobber
	RobbedNeighboringCommunity

	// Building
	BuiltRoad
	BuiltSettlement
	BuiltCity

	// Trading
	AcceptedTrade
	OfferedTrade
	PerformedMaritimeTrade

	// Development cards
	PurchasedDevelopmentCard
	PlayedMonopolyCard
	PlayedMonumentCard
	PlayedSoldierCard
	PlayedYearOfPlentyCard
)
