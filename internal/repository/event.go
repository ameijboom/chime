package repository

func (e *Event) MaxSize() int {
	switch *e.PartyType {
	case Full:
		return 8
	case Light:
		return 4
	default:
		return 0
	}
}
