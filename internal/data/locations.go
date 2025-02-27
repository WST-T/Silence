package data

import (
	"github.com/WST-T/Silence/internal/models"
)

// InitializeLocations creates all the locations for the game
func InitializeLocations() map[string]*models.Location {
	locations := make(map[string]*models.Location)

	locationData := []struct {
		name        string
		description string
	}{
		{
			name:        "Dining Room",
			description: "A grand room with a long oak table that seats twenty. Portraits of previous owners hang on the walls, their eyes seeming to follow you. The chandelier casts eerie shadows on the faded wallpaper.",
		},
		{
			name:        "Library",
			description: "Floor-to-ceiling bookshelves line the walls, filled with ancient tomes. A ladder on rails provides access to the highest shelves. The room smells of old paper and leather. A single desk lamp provides dim illumination.",
		},
		{
			name:        "Conservatory",
			description: "Tropical plants crowd this glass-walled room. The rain patters against the roof, and occasional flashes of lightning illuminate the verdant space. Some of the plants appear unusual and possibly toxic.",
		},
		{
			name:        "Study",
			description: "A cozy room with a large desk and leather chair. Hunting trophies adorn the walls, their glass eyes reflecting the light from the crackling fireplace. Maps and documents are spread across the desk.",
		},
		{
			name:        "Drawing Room",
			description: "Elegant furniture arranged around a central fireplace. The mantelpiece holds silver-framed photographs of young people in graduation attire. One photo has a black ribbon across the corner - Sarah Williams.",
		},
		{
			name:        "Kitchen",
			description: "A spacious kitchen with modern appliances contrasting with the old-fashioned architecture. Copper pots hang above a large island. Several knives are missing from the knife block on the counter.",
		},
		{
			name:        "Master Bedroom",
			description: "A four-poster bed dominates this lavish room. Heavy curtains block most of the moonlight. The antique vanity mirror seems to distort reflections slightly. A half-open closet reveals elegant clothing.",
		},
		{
			name:        "Hallway",
			description: "A long corridor connecting the manor's many rooms. The runner carpet muffles footsteps, and sconces cast pools of light at intervals. Class photos from various years hang on the walls, including your graduation year.",
		},
	}

	// Create all locations
	for _, loc := range locationData {
		locations[loc.name] = &models.Location{
			Name:        loc.name,
			Description: loc.description,
			Characters:  []string{},
		}
	}

	return locations
}

// GetStartingLocations returns a list of locations characters can start in
func GetStartingLocations() []string {
	return []string{
		"Drawing Room", "Library", "Conservatory", "Study", "Hallway",
	}
}
