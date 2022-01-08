package position

type PositionsRepository interface {
	BulkLoad() (*Positions, error)
}
