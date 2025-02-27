package game

import (
	"fmt"
)

// goToLocation moves the player to a different location
func (g *Game) goToLocation(locationName string) {
	// Check if the location exists
	_, exists := g.Locations[locationName]
	if !exists {
		fmt.Println("That location doesn't exist in the manor. Try again.")
		return
	}

	// Move to the new location
	g.PlayerLocation = locationName
	fmt.Println("You move to the", locationName)

	// Describe the new location
	g.describeCurrentLocation()
}

// talkToCharacter initiates a conversation with a character
func (g *Game) talkToCharacter(characterName string) {
	// Check if the character exists
	character, exists := g.Characters[characterName]
	if !exists {
		fmt.Println("There is no one here by that name.")
		return
	}

	// Check if the character is in the current location
	if character.Location != g.PlayerLocation {
		fmt.Println(characterName, "is not in this room.")
		return
	}

	// Check if the character is alive
	if !character.IsAlive {
		fmt.Println("You cannot talk to", characterName, "because they are dead.")
		return
	}

	// Display dialogue with the character
	fmt.Println("\nYou approach", characterName, "and start a conversation.")

	// Basic information about their background
	fmt.Println("\n"+characterName+":", "\"Hello, Detective. Terrible situation we have here, isn't it?\"")

	// Their relationship to Sarah
	fmt.Println("\nYou ask about their connection to Sarah Williams.")
	fmt.Println("\n"+characterName+":", "\""+character.Background+"\"")

	// Their thoughts on the current murders
	fmt.Println("\nYou ask if they noticed anything suspicious before the murders.")
	fmt.Println("\n"+characterName+":", "\""+character.Clue+"\"")

	// NEW: Store the clue in collected clues
	g.CollectedClues[characterName] = character.Clue

	// Even more nuanced version
	if spouse, exists := g.Characters[character.Partner]; exists && spouse.IsAlive {
		if spouse.Location == g.PlayerLocation {
			// Spouse is in the same room
			fmt.Printf("\n%s glances toward %s with a concerned expression.\n", characterName, character.Partner)
		} else {
			// Spouse is alive but in a different room
			fmt.Printf("\n%s seems distracted, occasionally glancing toward the door as if worried about %s.\n",
				characterName, character.Partner)
		}
	}

	// If their spouse is dead, they show grief
	if spouse, exists := g.Characters[character.Partner]; exists && !spouse.IsAlive {
		fmt.Printf("\n%s's eyes fill with tears. \"I still can't believe %s is gone. Who would do this?\"\n",
			characterName, character.Partner)
	}
}

// accuseCharacter accuses a character of being the killer
func (g *Game) accuseCharacter(characterName string) {
	// Check if the character exists
	character, exists := g.Characters[characterName]
	if !exists {
		fmt.Println("There is no one here by that name.")
		return
	}

	// Check if the character is alive
	if !character.IsAlive {
		fmt.Println("You cannot accuse", characterName, "because they are dead.")
		return
	}

	// Display accusation dialogue
	fmt.Println("\nYou point at", characterName, "dramatically.")
	fmt.Println("\"I accuse you of being the killer!\"")

	// Character's reaction to accusation
	fmt.Println("\n"+characterName+":", "\""+character.Suspicion+"\"")

	// Check if the accusation is correct
	if character.IsKiller {
		fmt.Println("\nThe room falls silent as", characterName, "looks around frantically.")
		fmt.Println("Suddenly, their expression changes to a cold, calculating stare.")
		fmt.Println("\n" + characterName + ": \"Well, Detective, you've caught me. But how did you know?\"")
		fmt.Println("\nAs security is called, you explain the clues that led to your conclusion.")
		fmt.Println("\nCONGRATULATIONS! You've solved the mystery and caught the killer!")
		fmt.Println("Sarah Williams can finally rest in peace, and her classmates are safe from her killer's revenge.")
		g.GameOver = true
	} else {
		fmt.Println("\nOther guests look uncomfortable. Your accusation seems to have fallen flat.")
		fmt.Println("Perhaps you should gather more evidence before making accusations.")
		fmt.Println("\nYou apologize for the misunderstanding, but keep", characterName, "in mind.")
	}
}

// showClues displays all collected information
func (g *Game) showClues() {
	fmt.Println("\nCollected Information:")

	// Display information about the deaths
	if g.MurderCount > 0 {
		fmt.Println("\nMurders:")
		fmt.Printf("  Murder Events: %d\n", g.MurderCount)
		fmt.Printf("  People Murdered: %d\n", g.MurderCount*2)

		fmt.Println("\nVictims:")
		for name, character := range g.Characters {
			if !character.IsAlive {
				fmt.Printf("  - %s (Partner: %s)\n", name, character.Partner)
			}
		}
	}

	// NEW: Display collected clues from characters
	if len(g.CollectedClues) > 0 {
		fmt.Println("\nClues from Conversations:")
		for name, clue := range g.CollectedClues {
			fmt.Printf("  %s: \"%s\"\n", name, clue)
		}
	}

	// Display information about Sarah Williams
	fmt.Println("\nAbout Sarah Williams (Deceased Classmate):")
	fmt.Println("  - Sarah died under mysterious circumstances shortly after graduation")
	fmt.Println("  - Her death was officially ruled an accident")
	fmt.Println("  - Many of the guests had complicated relationships with her")
	fmt.Println("  - Some guests seem uncomfortable discussing her")
}
