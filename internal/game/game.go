package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/WST-T/Silence/internal/data"
	"github.com/WST-T/Silence/internal/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Game holds the game state
type Game struct {
	Characters        map[string]*models.Character
	Locations         map[string]*models.Location
	KillerName        string
	CurrentTurn       int
	GameOver          bool
	PlayerLocation    string
	LastKilled        []string // Names of the last characters killed
	MurderCount       int      // Number of murder events
	TurnsSinceKilling int      // Turns elapsed since last murder event
}

// NewGame initializes a new game
func NewGame() *Game {
	game := &Game{
		CurrentTurn:       1,
		GameOver:          false,
		MurderCount:       0,
		TurnsSinceKilling: 0,
	}

	// Initialize characters, locations, and other game elements
	game.Characters = data.InitializeCharacters()
	game.Locations = data.InitializeLocations()
	game.selectKiller()
	game.distributeCharacters()

	// Player starts in the dining room
	game.PlayerLocation = "Dining Room"

	return game
}

// selectKiller randomly selects a killer from among the characters
func (g *Game) selectKiller() {
	// Create a slice of character names
	var names []string
	for name := range g.Characters {
		names = append(names, name)
	}

	// Randomly select a killer
	killerIndex := rng.Intn(len(names))
	killerName := names[killerIndex]

	// Set the selected character as the killer
	g.Characters[killerName].SetAsKiller()
	g.KillerName = killerName
}

// distributeCharacters places characters in different starting locations
func (g *Game) distributeCharacters() {
	// Get possible starting locations
	startingLocations := data.GetStartingLocations()

	// Place each character in a random location
	for name, character := range g.Characters {
		randLocIndex := rng.Intn(len(startingLocations))
		location := startingLocations[randLocIndex]

		character.Location = location
		g.Locations[location].AddCharacter(name)
	}
}

// Play starts the main game loop
func (g *Game) Play() {
	reader := bufio.NewReader(os.Stdin)

	// Show introduction
	showIntroduction()

	fmt.Println("\nPress Enter to begin...")
	reader.ReadString('\n')

	// Initial location description
	g.describeCurrentLocation()
	g.showCommands()

	// Turn counter for timed events
	turnCount := 0

	// Main game loop
	for !g.GameOver {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		command := args[0]

		switch command {
		case "look":
			g.describeCurrentLocation()

		case "go":
			if len(args) < 2 {
				fmt.Println("Go where? Use 'go <location>'.")
				continue
			}
			// Combine all remaining words as the location name
			locationName := strings.Join(args[1:], " ")
			// Capitalize each word in the location name
			words := strings.Fields(locationName)
			caser := cases.Title(language.English)
			for i, word := range words {
				words[i] = caser.String(word)
			}
			locationName = strings.Join(words, " ")
			g.goToLocation(locationName)
			turnCount++
			g.TurnsSinceKilling++ // Increment turns since killing

		case "talk":
			if len(args) < 2 {
				fmt.Println("Talk to whom? Use 'talk <person>'.")
				continue
			}
			// Combine all remaining words as the person's name
			personName := strings.Join(args[1:], " ")
			// Capitalize each word in the person's name
			words := strings.Fields(personName)
			caser := cases.Title(language.English)
			for i, word := range words {
				words[i] = caser.String(word)
			}
			personName = strings.Join(words, " ")
			g.talkToCharacter(personName)
			turnCount++
			g.TurnsSinceKilling++ // Increment turns since killing

		case "accuse":
			if len(args) < 2 {
				fmt.Println("Accuse whom? Use 'accuse <person>'.")
				continue
			}
			// Combine all remaining words as the person's name
			personName := strings.Join(args[1:], " ")
			// Capitalize each word in the person's name
			words := strings.Fields(personName)
			caser := cases.Title(language.English)
			for i, word := range words {
				words[i] = caser.String(word)
			}
			personName = strings.Join(words, " ")
			g.accuseCharacter(personName)

		case "clues":
			g.showClues()

		case "locations":
			g.showLocations()

		case "status":
			g.showStatus()

		case "help":
			g.showCommands()

		case "quit":
			fmt.Println("Are you sure you want to quit? (y/n)")
			confirm, _ := reader.ReadString('\n')
			confirm = strings.TrimSpace(strings.ToLower(confirm))
			if confirm == "y" || confirm == "yes" {
				fmt.Println("Thanks for playing!")
				return
			}

		default:
			fmt.Println("I don't understand that command. Type 'help' for a list of commands.")
		}

		// Check if the game is over
		if g.GameOver {
			break
		}

		// Execute murders after every 5 turns have passed since the last murder
		if turnCount > 0 && g.TurnsSinceKilling >= 5 {
			g.executeMurders()
			g.TurnsSinceKilling = 0 // Reset counter after murder
		}
	}

	// End of game
	fmt.Println("\nOver..For now..")
	fmt.Println("Thank you for playing Silence.")
}
