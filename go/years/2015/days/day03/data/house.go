package data

// House represents a house
type House struct {
	Longitude int // negative int are west of reference and positive int are east
	Latitude  int // negative int are south of ref point and positive int are north
}

// Matches returns a Boolean if the current house matches the house passed in.
func (h House) Matches(house House) bool {
	return h.Longitude == house.Longitude && h.Latitude == house.Latitude
}
