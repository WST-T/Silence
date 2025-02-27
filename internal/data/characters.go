package data

import (
	"github.com/WST-T/Silence/internal/models"
)

// InitializeCharacters creates all the characters for the game
func InitializeCharacters() map[string]*models.Character {
	characters := make(map[string]*models.Character)

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
		characters[couple.name1] = &models.Character{
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
		characters[couple.name2] = &models.Character{
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

	return characters
}
