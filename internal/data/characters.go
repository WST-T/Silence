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
			name1:       "Zoro Roronoa",
			name2:       "Penny Roronoa",
			background1: "Former sport leader and exceptional swordsman of the class. Despite his impressive athletic abilities, he was known for his terrible sense of direction and straightforward thinking. Was Sarah Williams' training partner.",
			background2: "Talented artist who created all illustrations for the school yearbook. Extremely introverted and rarely shared her thoughts, making her hard to read. Designed a memorial artwork after Sarah's passing.",
			secret1:     "Was practicing sword techniques in the woods near where Sarah died. Heard a scream but got lost trying to find the source and arrived too late to help.",
			secret2:     "Created a series of sketches depicting what she believed happened to Sarah, hidden in her art portfolio. Never showed them to anyone, fearing the implications.",
			clue1:       "I saw Mathieu Lemaire examining the old training equipment in the garden. He seemed fixated on the fencing foils, especially the one with the damaged tip.",
			clue2:       "When I was sketching in the library, I noticed Siqi Wang removing something from behind the old class photo - it looked like a folded note.",
			suspicion1:  "You think I killed them? Ridiculous! If I wanted to challenge someone, I'd do it face to face, not from the shadows like a coward!",
			suspicion2:  "Me? I barely speak to anyone here. Why would I suddenly start killing people? I observe life, I don't take it.",
		},
		{
			name1:       "Mathieu Lemaire",
			name2:       "Aria Lemaire",
			background1: "Passionate financial advisor known for his hot-headed nature but genuine charm. Helped organize the class fundraiser for Sarah's memorial but lost his temper during a disagreement about fund allocation.",
			background2: "Calm and calculating investment analyst who balances out her husband's explosive temperament. Managed the accounting for Sarah's memorial fund with meticulous precision.",
			secret1:     "Discovered Sarah was investigating financial irregularities in the student council just before her death. Lost his temper and confronted her, though their argument was witnessed by several students.",
			secret2:     "Quietly paid off a security guard to destroy footage showing her husband's heated argument with Sarah the day before she died, believing it would only cause unnecessary suspicion.",
			clue1:       "I noticed Ed Wang checking the value of an unusual antique in the library. When he saw me watching, he quickly put it back and claimed it was just a family heirloom.",
			clue2:       "While everyone was distracted at dinner, I saw Stephane Pizeuil slip what looked like financial documents into his jacket. Something about investment opportunities.",
			suspicion1:  "You're accusing ME? That's ridiculous! Yes, I have a temper, but I'd never KILL anyone over it! I make money LEGALLY, not through murder!",
			suspicion2:  "What a preposterous suggestion. If I wanted someone eliminated, I'd simply destroy their finances, not resort to something as messy and unprofitable as murder.",
		},
		{
			name1:       "Ed Wang",
			name2:       "Siqi Wang",
			background1: "Famous guitarist and singer who performed at school events. Wrote and dedicated a song to Sarah's memory that became his first hit single after graduation.",
			background2: "Talented robotics engineer who created the automated lighting system for the school theater. Worked with Sarah on a science project combining art and technology.",
			secret1:     "Was recording music in the studio near where Sarah died and heard her arguing with someone. Never came forward as a witness to avoid publicity that could affect his rising music career.",
			secret2:     "Discovered a malfunction in the stage equipment that may have contributed to Sarah's accident. Fixed it quickly and never told anyone, fearing professional repercussions.",
			clue1:       "Guillaume Mardrus was handling Sarah's old songbook last night. When I approached, he quickly put it away and acted strangely defensive.",
			clue2:       "I noticed Minseong Pizeuil examining the electrical panel in the hallway before dinner. She seemed to know exactly what she was looking for.",
			suspicion1:  "Me, a killer? I've spent my life bringing joy through music! Why would I destroy what I create? My hands are for playing guitars, not taking lives!",
			suspicion2:  "This accusation is illogical. I build systems that follow precise programming. Murder is chaotic and irrational - completely against my nature.",
		},
		{
			name1:       "Stephane Pizeuil",
			name2:       "Minseong Pizeuil",
			background1: "Cloud and DevOps Engineer at Rockstar Games currently working on the new GTA title. During high school, he created a digital memorial website for Sarah that became an unexpected viral tribute among gaming communities.",
			background2: "Talented and famous pianist who studied at Juilliard. Performed a moving musical tribute at Sarah's memorial service despite having transferred to the school only months before the tragedy.",
			secret1:     "While archiving the school's digital records for the memorial website, discovered suspicious messages on Sarah's school account but kept them private to protect the school's reputation and his future career prospects.",
			secret2:     "Was practicing piano late in the music room the night Sarah died and heard unusual noises from the theater next door. Never mentioned it because as a new student, she feared getting involved in a scandal.",
			clue1:       "I noticed Vincent Mardrus trying to access the old security system logs on his phone. When I offered technical help, he quickly put the phone away and changed the subject.",
			clue2:       "During my performance at dinner, I saw Maya Vu reaction when I played Sarah's favorite piece. She seemed physically pained by the melody, as if it triggered something traumatic.",
			suspicion1:  "You think I'm responsible? That's absurd! I build systems to protect data and people. My code has security protocols and audit trails - if I were a killer, I'd leave digital evidence everywhere!",
			suspicion2:  "Me? A killer? My entire career is built on bringing beauty and emotion to audiences. Violence destroys harmony - it goes against everything I stand for as an artist.",
		},
		{
			name1:       "Guillaume Mardrus",
			name2:       "Vincent Mardrus",
			background1: "Witty accountant who managed the class finances with precise humor. Organized the budget for Sarah's memorial event, adding personal touches like custom bookmarks with her favorite quotes.",
			background2: "Creative software engineer known for his outlandish app ideas. Developed a commemorative digital yearbook featuring Sarah that included interactive memories and photos.",
			secret1:     "Found financial discrepancies in the school's memorial fund but kept quiet when he traced them back to a faculty member he deeply respected. Still carries the evidence in an encrypted spreadsheet.",
			secret2:     "Created tracking software to help find Sarah when she went missing, but the data revealed her location too late. Deleted the logs out of guilt but kept the original algorithm.",
			clue1:       "I noticed Souhib Trabelsi taking photos of the old yearbook pages with his phone when he thought no one was watching. Specifically the treasurer's notes section.",
			clue2:       "While setting up the WiFi, I saw Ed Wang's laptop had multiple tabs open researching 'undetectable poisons' and 'famous unsolved murders.' He closed them when I offered technical help.",
			suspicion1:  "Me? A murderer? Please! My spreadsheets are the only things I kill - with efficiency! Besides, can you imagine the tax implications of homicide? Absolutely dreadful!",
			suspicion2:  "That's the most absurd bug report I've ever received! If I designed a murder algorithm, it would be far more sophisticated than whatever crude method was used here. This is amateur work!",
		},
		{
			name1:       "Minh-Duy Vu",
			name2:       "Maya Vu",
			background1: "Ambitious finance executive known for transforming struggling companies into profitable ventures. Organized a lavish fundraiser for Sarah's memorial that became the social event of the year.",
			background2: "Former high-level financial strategist from Lebanon who met her husband during an international merger. Now enjoys life as a pampered socialite who donated generously to Sarah's memorial fund.",
			secret1:     "Discovered Sarah was investigating unethical investment practices that could have implicated some of his biggest clients. Paid her to stay quiet about her findings.",
			secret2:     "Used her connections to ensure certain documents related to Sarah's investigations disappeared from the school archives, protecting both her husband's reputation and her luxurious lifestyle.",
			clue1:       "Guillaume was acting strangely when I mentioned old financial records. He practically ran from the room when I suggested we look through the reunion archives.",
			clue2:       "I noticed Penny Roronoa carrying what looked like old sketches from her portfolio. She seemed startled when I complimented her artistic talent and quickly hid them away.",
			suspicion1:  "Murder? Please! I build financial empires, not criminal records! Besides, eliminating problems with money is far more effective than... messier solutions.",
			suspicion2:  "What an absurd accusation! I left that cutthroat corporate world behind. Why would I risk my perfect life for something so... distasteful?",
		},
		{
			name1:       "Souhib Trabelsi",
			name2:       "Aaliyah Trabelsi",
			background1: "Brilliant software engineer known for his Python expertise, though infamous for falling asleep during meetings and taking suspiciously long bathroom breaks. Created a memorial software that analyzed Sarah's digital footprint after her death.",
			background2: "Renowned chef from a famous Dubai restaurant who catered the memorial service for Sarah. Her Middle Eastern fusion cuisine became a comforting tradition at reunion events despite her limited connection to the original class.",
			secret1:     "Discovered suspicious patterns in Sarah's online communications while analyzing her digital footprint, but accepted payment to manipulate the data and remove certain conversations from the final report.",
			secret2:     "Overheard a heated argument about Sarah on the phone while preparing food for the memorial service, recognized the voice but never revealed who it was, fearing it would damage her prestigious culinary reputation.",
			clue1:       "I noticed Mathieu Lemaire deleting something from his phone when I offered to help with his banking app. He seemed unusually defensive about his call history.",
			clue2:       "While preparing dinner, I saw Vincent Mardrus checking the ingredients in the kitchen with unusual scrutiny, asking specifically about which dishes contained certain herbs and spices.",
			suspicion1:  "You think I'm the killer? That's absurd! I build software that connects people! Besides, I was probably asleep in a meeting or...elsewhere...when these murders happened!",
			suspicion2:  "Me? A killer? I create food that brings joy! My hands craft delicacies, not death! This accusation is as distasteful as pairing caviar with ketchup!",
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
