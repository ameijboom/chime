package database

func (e *Event) MaxSize() int {
	switch *e.PartyType {
	case Full:
		return 8
	case Light:
		return 4
	}

	return 0
}
