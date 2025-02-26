package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

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

// Location represents a room in the manor
type Location struct {
	Name        string
	Description string
	Characters  []string // Names of characters in this location
}

// Game holds the game state
type Game struct {
	Characters        map[string]*Character
	Locations         map[string]*Location
	KillerName        string
	CurrentTurn       int
	GameOver          bool
	PlayerLocation    string
	LastKilled        []string // Names of the last characters killed
	MurderCount       int      // Number of murder events
	TurnsSinceKilling int      // Turns elapsed since last murder event
}

// Initialize the game
func NewGame() *Game {
	game := &Game{
		Characters:        make(map[string]*Character),
		Locations:         make(map[string]*Location),
		CurrentTurn:       1,
		GameOver:          false,
		MurderCount:       0,
		TurnsSinceKilling: 0,
	}

	// Initialize characters, locations, and other game elements
	game.initializeCharacters()
	game.initializeLocations()
	game.selectKiller()
	game.distributeCharacters()

	// Player starts in the dining room
	game.PlayerLocation = "Dining Room"

	return game
}

// Initialize all characters in the game
func (g *Game) initializeCharacters() {
	// Create pairs of characters (couples)
	couples := []struct {
		name1       string
		name2       string
		background1 string
		background2 string
		secret1     string
		secret2     string
		clue1       string
		clue2       string
		suspicion1  string
		suspicion2  string
	}{
		{
			name1:       "Robert Thompson",
			name2:       "Elizabeth Thompson",
			background1: "Class treasurer who managed the funds for senior events. Had an unrequited crush on the deceased classmate, Sarah Williams.",
			background2: "Met Robert after college. Never knew Sarah Williams personally but heard many stories.",
			secret1:     "Had an argument with Sarah the night she died about missing funds from the class account.",
			secret2:     "Suspects her husband might have been involved with Sarah more deeply than he admits.",
			clue1:       "I noticed James acting strangely when Sarah's name was mentioned during dinner. His face went pale.",
			clue2:       "Margaret kept glancing at the old class photo on the wall, specifically at Sarah. There seemed to be anger in her eyes.",
			suspicion1:  "Me? The killer? That's absurd! I was devastated by Sarah's death. We all were...",
			suspicion2:  "I wasn't even part of your class. Why would I have any motive to kill anyone here?",
		},
		{
			name1:       "James Wilson",
			name2:       "Patricia Wilson",
			background1: "Star quarterback who dated Sarah Williams briefly in senior year before she broke it off.",
			background2: "College sweetheart of James, always jealous of his past relationship with Sarah.",
			secret1:     "Still keeps a locket with Sarah's picture hidden in his wallet.",
			secret2:     "Anonymously sent threatening letters to Sarah in senior year to stay away from James.",
			clue1:       "Thomas and Catherine arrived separately to the dinner, even though they're married. Something seemed off between them.",
			clue2:       "I overheard William on the phone earlier, he sounded nervous and said something about 'keeping quiet about the past.'",
			suspicion1:  "This is ridiculous! Sarah's death was ruled an accident. Why would I harm anyone here?",
			suspicion2:  "You think I'd wait all these years for revenge? I've moved on from petty high school jealousy!",
		},
		{
			name1:       "Thomas Clark",
			name2:       "Catherine Clark",
			background1: "Class president who worked closely with Sarah on the yearbook committee.",
			background2: "Former lab partner of Sarah in chemistry class. They had a falling out over a failed experiment.",
			secret1:     "Was the last person to see Sarah alive according to the original investigation, but never told anyone.",
			secret2:     "Found Sarah's diary after her death and destroyed pages that implicated several classmates in bullying her.",
			clue1:       "Robert keeps checking his watch and seems anxious, like he's waiting for something to happen.",
			clue2:       "Jennifer avoided walking past the memorial photo of Sarah and made an excuse to leave the room when stories about Sarah came up.",
			suspicion1:  "I had nothing but respect for Sarah! We worked well together on the committee!",
			suspicion2:  "This is absurd! Why would I kill people at a reunion dinner? That makes no sense!",
		},
		{
			name1:       "William Davis",
			name2:       "Jennifer Davis",
			background1: "School newspaper editor who wrote the memorial piece after Sarah's death.",
			background2: "Transfer student who only knew Sarah briefly but became friends with her quickly.",
			secret1:     "Changed details in his article about Sarah's death to protect certain classmates from scrutiny.",
			secret2:     "Was with Sarah the night she died but left her alone because she wanted to meet someone secretly.",
			clue1:       "I noticed Margaret and David arguing quietly in the corner before the lights went out.",
			clue2:       "Michael seemed particularly disturbed by the mention of Sarah during dinner. He nearly dropped his glass.",
			suspicion1:  "I honored Sarah's memory with my article! I would never dishonor her by harming our classmates!",
			suspicion2:  "I barely knew Sarah! I have no motive! This is insane!",
		},
		{
			name1:       "Michael Brown",
			name2:       "Margaret Brown",
			background1: "Science genius who tutored Sarah in physics. They were close friends.",
			background2: "Cheer captain who had a rivalry with Sarah over a position on the squad.",
			secret1:     "Developed feelings for Sarah and confessed to her the week before she died. She rejected him.",
			secret2:     "Pushed Sarah during an argument at a party, causing her to hit her head - may have contributed to her later accident.",
			clue1:       "Patricia was in the bathroom for an unusually long time right before the lights went out.",
			clue2:       "I saw Thomas slip something into his pocket from the mantelpiece right before dinner. It looked like an old key or small object.",
			suspicion1:  "Sarah was my friend! I was devastated when she died! This is outrageous!",
			suspicion2:  "Our rivalry was just typical high school drama! I never wished any real harm on Sarah!",
		},
		{
			name1:       "David Anderson",
			name2:       "Rebecca Anderson",
			background1: "Track team captain who trained with Sarah for cross-country events.",
			background2: "Drama club president who competed with Sarah for the lead in the senior play.",
			secret1:     "Supplied Sarah with performance-enhancing drugs that may have affected her health.",
			secret2:     "Sabotaged Sarah's audition by tampering with the stage props, causing an accident that injured Sarah.",
			clue1:       "Elizabeth was looking through the host's desk drawer when she thought no one was watching.",
			clue2:       "James seemed to recognize something in the study - an old book or document. He looked disturbed.",
			suspicion1:  "Sarah was a friend and teammate! I would never hurt anyone from our class!",
			suspicion2:  "That was theatrical rivalry, not murder motive! This accusation is absurd!",
		},
		{
			name1:       "Richard Miller",
			name2:       "Susan Miller",
			background1: "Computer club president who helped Sarah recover deleted emails containing threats from an anonymous classmate.",
			background2: "Art student who painted Sarah's portrait for the memorial service after her death.",
			secret1:     "Discovered who was threatening Sarah but kept it secret when that person paid him to stay quiet.",
			secret2:     "Included hidden symbols in the memorial portrait hinting at her suspicions about Sarah's death.",
			clue1:       "William went to the garden alone shortly after dinner, even though it was raining. He seemed to be searching for something.",
			clue2:       "Catherine spent a long time staring at Sarah's old locker combination that's still preserved in a frame in the hallway.",
			suspicion1:  "I tried to help Sarah! Why would I hurt anyone connected to her case?",
			suspicion2:  "I honored Sarah through my art! This accusation is both hurtful and senseless!",
		},
	}

	// Create all characters from the couples
	for _, couple := range couples {
		// First partner
		g.Characters[couple.name1] = &Character{
			Name:       couple.name1,
			Partner:    couple.name2,
			Background: couple.background1,
			Secret:     couple.secret1,
			IsKiller:   false,
			IsAlive:    true,
			Clue:       couple.clue1,
			Suspicion:  couple.suspicion1,
		}

		// Second partner
		g.Characters[couple.name2] = &Character{
			Name:       couple.name2,
			Partner:    couple.name1,
			Background: couple.background2,
			Secret:     couple.secret2,
			IsKiller:   false,
			IsAlive:    true,
			Clue:       couple.clue2,
			Suspicion:  couple.suspicion2,
		}
	}
}

// Initialize all locations in the manor
func (g *Game) initializeLocations() {
	locations := []struct {
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
	for _, loc := range locations {
		g.Locations[loc.name] = &Location{
			Name:        loc.name,
			Description: loc.description,
			Characters:  []string{},
		}
	}
}

// Randomly select a killer from among the characters
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
	g.Characters[killerName].IsKiller = true
	g.KillerName = killerName

	// Add some special dialogue for the killer
	g.Characters[killerName].Clue = "I think we should all stay together. It's safer that way. I don't trust anyone wandering alone in this old manor."
	g.Characters[killerName].Suspicion = "Me? The killer? *laughs nervously* That's absurd. I've been trying to keep everyone safe! You have no proof!"
}

// Distribute characters across different locations
func (g *Game) distributeCharacters() {
	// Define possible starting locations (excluding some locations for variety)
	startingLocations := []string{
		"Drawing Room", "Library", "Conservatory", "Study", "Hallway",
	}

	// Place each character in a random location
	for name, character := range g.Characters {
		randLocIndex := rng.Intn(len(startingLocations))
		location := startingLocations[randLocIndex]

		character.Location = location
		g.Locations[location].Characters = append(g.Locations[location].Characters, name)
	}
}

// Show the current location and characters present
func (g *Game) describeCurrentLocation() {
	location := g.Locations[g.PlayerLocation]

	fmt.Println("\n===", g.PlayerLocation, "===")
	fmt.Println(location.Description)

	// List characters in the room
	if len(location.Characters) == 0 {
		fmt.Println("There is no one else in this room.")
	} else {
		fmt.Println("\nIn this room:")
		for _, name := range location.Characters {
			if g.Characters[name].IsAlive {
				fmt.Println("  -", name)
			}
		}
	}
}

// List available commands
func (g *Game) showCommands() {
	fmt.Println("\nWhat would you like to do?")
	fmt.Println("  [look] - Look around the room again")
	fmt.Println("  [go <location>] - Go to another room")
	fmt.Println("  [talk <person>] - Talk to someone in the room")
	fmt.Println("  [accuse <person>] - Accuse someone of being the killer")
	fmt.Println("  [clues] - Review collected clues and information")
	fmt.Println("  [locations] - See all locations in the manor")
	fmt.Println("  [status] - Check the status of all guests")
	fmt.Println("  [help] - Show this help menu")
	fmt.Println("  [quit] - Exit the game")
}

// List all locations in the manor
func (g *Game) showLocations() {
	fmt.Println("\nLocations in the manor:")
	for name := range g.Locations {
		fmt.Println("  -", name)
	}
}

// Show status of all guests (alive or dead)
func (g *Game) showStatus() {
	fmt.Println("\nGuest Status:")
	for name, character := range g.Characters {
		status := "Alive"
		if !character.IsAlive {
			status = "Dead"
		}
		fmt.Printf("  - %s: %s (Location: %s)\n", name, status, character.Location)
	}

	// Show murder count
	fmt.Printf("\nMurder Events: %d\n", g.MurderCount)
	fmt.Printf("People Murdered: %d\n", g.MurderCount*2)
	fmt.Printf("People Remaining Alive: %d\n", len(g.Characters)-g.MurderCount*2)
}

// Move to a different location
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

// Talk to a character in the current location
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

	// If this is their spouse, they share a concerned look
	if spouse, exists := g.Characters[character.Partner]; exists && spouse.IsAlive {
		fmt.Printf("\n%s glances toward %s with a concerned expression.\n", characterName, character.Partner)
	}

	// If their spouse is dead, they show grief
	if spouse, exists := g.Characters[character.Partner]; exists && !spouse.IsAlive {
		fmt.Printf("\n%s's eyes fill with tears. \"I still can't believe %s is gone. Who would do this?\"\n",
			characterName, character.Partner)
	}
}

// Accuse a character of being the killer
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

// Execute a murder event
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
	g.Characters[victim1].IsAlive = false
	g.Characters[victim2].IsAlive = false

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

	// Announce the deaths
	fmt.Printf("\n%s %s in the %s!\n", victim1, deathMethod1, deathLocation1)
	fmt.Printf("%s %s in the %s!\n", victim2, deathMethod2, deathLocation2)

	// Save the names of the last killed characters
	g.LastKilled = []string{victim1, victim2}

	// Increment murder count
	g.MurderCount++

	// Remove the dead characters from their locations
	removeCharacterFromLocation(g, victim1)
	removeCharacterFromLocation(g, victim2)

	fmt.Println("\nThe remaining guests are terrified. You must solve this mystery quickly before more lives are lost!")
	fmt.Println("----------------------------")
}

// Helper function to get a random location name
func getRandomLocation(g *Game) string {
	locations := make([]string, 0, len(g.Locations))
	for name := range g.Locations {
		locations = append(locations, name)
	}
	return locations[rng.Intn(len(locations))]
}

// Helper function to remove a character from their location
func removeCharacterFromLocation(g *Game, characterName string) {
	character := g.Characters[characterName]
	location := g.Locations[character.Location]

	// Create a new slice without the character
	newCharacters := []string{}
	for _, name := range location.Characters {
		if name != characterName {
			newCharacters = append(newCharacters, name)
		}
	}

	location.Characters = newCharacters
}

// Show collected clues and information
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

	// Display information about Sarah Williams
	fmt.Println("\nAbout Sarah Williams (Deceased Classmate):")
	fmt.Println("  - Sarah died under mysterious circumstances shortly after graduation")
	fmt.Println("  - Her death was officially ruled an accident")
	fmt.Println("  - Many of the guests had complicated relationships with her")
	fmt.Println("  - Some guests seem uncomfortable discussing her")
}

// Main game loop
func (g *Game) play() {
	reader := bufio.NewReader(os.Stdin)

	// Introduction
	fmt.Println("======================================")
	fmt.Println("       		SILENCE                ")
	fmt.Println("======================================")
	fmt.Println("\nIt's been 15 years since graduation. Your old classmates have gathered at")
	fmt.Println("Blackwood Manor for a reunion dinner. The mood is tense as conversations")
	fmt.Println("turn to Sarah Williams, a classmate who died under mysterious circumstances")
	fmt.Println("shortly after graduation.")
	fmt.Println("\nAs a detective, you notice the strange undercurrents among the guests...")
	fmt.Println("Then the lights go out, and the killing begins.")
	fmt.Println("\nCan you find the killer before everyone is dead?")
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

func main() {
	// Create and start a new game
	game := NewGame()
	game.play()
}
