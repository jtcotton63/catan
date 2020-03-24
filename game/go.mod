module github.com/jtcotton63/catan/game

go 1.14

require (
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.4.0
	github.com/jtcotton63/catan/event v0.0.0
	github.com/jtcotton63/catan/eventtype v0.0.0
	github.com/jtcotton63/catan/model v0.0.0
	github.com/jtcotton63/catan/state v0.0.0
	github.com/jtcotton63/catan/statetype v0.0.0
	github.com/pkg/errors v0.9.1
)

replace github.com/jtcotton63/catan/event => ../event

replace github.com/jtcotton63/catan/eventtype => ../eventtype

replace github.com/jtcotton63/catan/model => ../model

replace github.com/jtcotton63/catan/state => ../state

replace github.com/jtcotton63/catan/statetype => ../statetype
