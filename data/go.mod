module github.com/jtcotton63/catan/data

go 1.14

require (
	github.com/google/uuid v1.1.1
	github.com/jtcotton63/catan/event v0.0.0
	github.com/jtcotton63/catan/eventtype v0.0.0
	github.com/jtcotton63/catan/model v0.0.0
)

replace github.com/jtcotton63/catan/event => ../event

replace github.com/jtcotton63/catan/eventtype => ../eventtype

replace github.com/jtcotton63/catan/model => ../model
