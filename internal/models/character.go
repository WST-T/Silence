package models

// Character represents a person at the dinner party
type Character struct {
	Name       string
	Partner    string // Name of the spouse
	Background string // Brief background and relation to the deceased classmate
	Secret     string // Something they're hiding
	IsKiller   bool   // Whether this character is the killer
	IsAlive    bool   // Whether this character is alive
	Location   string // Current location in the manor
	Clue       string // Clue they provide when questioned
	Suspicion  string // What they say when accused
}

// IsInLocation checks if the character is in the specified location
func (c *Character) IsInLocation(locationName string) bool {
	return c.Location == locationName
}

// MoveToLocation moves the character to a new location
func (c *Character) MoveToLocation(locationName string) {
	c.Location = locationName
}

// Kill marks the character as dead
func (c *Character) Kill() {
	c.IsAlive = false
}

// SetAsKiller marks this character as the killer
func (c *Character) SetAsKiller() {
	c.IsKiller = true
	// Special dialogue for the killer
	c.Clue = "I think we should all stay together. It's safer that way. I don't trust anyone wandering alone in this old manor."
	c.Suspicion = "Me? The killer? *laughs nervously* That's absurd. I've been trying to keep everyone safe! You have no proof!"
}
