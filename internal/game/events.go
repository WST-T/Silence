package game

import (
	"fmt"
	"math/rand"
)

// executeMurders handles the killing of characters during the game
func (g *Game) executeMurders() {
	fmt.Println("\n----------------------------")
	fmt.Println("Suddenly, the lights in the manor flicker and go out completely!")
	fmt.Println("Panicked voices cry out in the darkness.")
	fmt.Println("When the lights return moments later, there are screams of horror...")

	// Select two random living non-killer characters to kill
	var livingCharacters []string
	for name, character := range g.Characters {
		if character.IsAlive && !character.IsKiller {
			livingCharacters = append(livingCharacters, name)
		}
	}

	// If there are fewer than 2 living non-killer characters, game over
	if len(livingCharacters) < 2 {
		fmt.Println("\nWith almost everyone dead, the killer makes their escape into the stormy night.")
		fmt.Println("You've failed to solve the mystery, and Sarah Williams' revenge remains incomplete.")
		fmt.Println("\nGAME OVER - The killer got away.")
		g.GameOver = true
		return
	}

	// Randomly select two victims
	rand.Shuffle(len(livingCharacters), func(i, j int) {
		livingCharacters[i], livingCharacters[j] = livingCharacters[j], livingCharacters[i]
	})

	victim1 := livingCharacters[0]
	victim2 := livingCharacters[1]

	// Kill the victims
	g.Characters[victim1].Kill()
	g.Characters[victim2].Kill()

	// Create dramatic death descriptions
	deathMethods := []string{
		"was found with a knife in their back",
		"appears to have been strangled with a curtain cord",
		"was discovered at the bottom of the stairs with a broken neck",
		"has a fatal blow to the head from a heavy object",
		"has been poisoned, their face frozen in a horrified expression",
	}

	deathLocation1 := getRandomLocation(g)
	deathLocation2 := getRandomLocation(g)

	deathMethod1 := deathMethods[rng.Intn(len(deathMethods))]
	deathMethod2 := deathMethods[rng.Intn(len(deathMethods))]

	// Remove the dead characters from their locations
	removeCharacterFromLocation(g, victim1)
	removeCharacterFromLocation(g, victim2)

	g.Characters[victim1].Location = deathLocation1
	g.Characters[victim2].Location = deathLocation2

	// Announce the deaths
	fmt.Printf("\n%s %s in the %s!\n", victim1, deathMethod1, deathLocation1)
	fmt.Printf("%s %s in the %s!\n", victim2, deathMethod2, deathLocation2)

	// Save the names of the last killed characters
	g.LastKilled = []string{victim1, victim2}

	// Increment murder count
	g.MurderCount++

	fmt.Println("\nThe remaining guests are terrified. You must solve this mystery quickly before more lives are lost!")
	fmt.Println("----------------------------")
}

// getRandomLocation returns the name of a random location
func getRandomLocation(g *Game) string {
	locations := make([]string, 0, len(g.Locations))
	for name := range g.Locations {
		locations = append(locations, name)
	}
	return locations[rng.Intn(len(locations))]
}

// removeCharacterFromLocation removes a character from their location
func removeCharacterFromLocation(g *Game, characterName string) {
	character := g.Characters[characterName]
	location := g.Locations[character.Location]
	location.RemoveCharacter(characterName)
}
