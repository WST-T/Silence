package models

// Location represents a room in the manor
type Location struct {
	Name        string
	Description string
	Characters  []string // Names of characters in this location
}

// AddCharacter adds a character to this location
func (l *Location) AddCharacter(characterName string) {
	l.Characters = append(l.Characters, characterName)
}

// RemoveCharacter removes a character from this location
func (l *Location) RemoveCharacter(characterName string) {
	newCharacters := []string{}
	for _, name := range l.Characters {
		if name != characterName {
			newCharacters = append(newCharacters, name)
		}
	}
	l.Characters = newCharacters
}

// GetLivingCharacters returns the names of characters that are alive in this location
func (l *Location) GetLivingCharacters(characters map[string]*Character) []string {
	livingCharacters := []string{}
	for _, name := range l.Characters {
		if character, exists := characters[name]; exists && character.IsAlive {
			livingCharacters = append(livingCharacters, name)
		}
	}
	return livingCharacters
}
