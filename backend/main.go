package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type question struct {
	ID string `json:"id"`
	Difficulty string `json:"diff"`
	Title string `json:"title"`
	AnswerA string `json:"a"`
	AnswerB string `json:"b"`
	AnswerC string `json:"c"`
	AnswerD string `json:"d"`
}

type correctAnswer struct {
	ID int
	CorrectAns string
}

type userAnswer struct {
	AnswerID string `json:"id" binding:"required"`
	Answer string `json:"answer" binding:"required"`
	Difficulty string `json:"diff" binding:"required"`
	CurrentScore int `json:"currScore" binding:"required"`
}

var eQuestions = []question{
	{
		ID: "1", 
		Difficulty: "easy", 
		Title: "Which of these roles do not exist in Unite?", 
		AnswerA: "Supporter", 
		AnswerB: "All-Rounder", 
		AnswerC: "Assassin", 
		AnswerD: "Defender",
	},
	{
		ID: "2", 
		Difficulty: "easy", 
		Title: "How many points can a Pokemon hold at maximum?", 
		AnswerA: "50", 
		AnswerB: "20", 
		AnswerC: "80", 
		AnswerD: "130",
	},
	{
		ID: "3", 
		Difficulty: "easy", 
		Title: "At what time does double point scoring start?",
		AnswerA: "5:00", 
		AnswerB: "1:00", 
		AnswerC: "0:30",
		AnswerD: "2:00",
	},
	{
		ID: "4", 
		Difficulty: "easy", 
		Title: "At what time do the Altaria and Swablu's in the middle of the map first spawn (Theria Sky Ruins)?", 
		AnswerA: "9:10",
		AnswerB: "8:00", 
		AnswerC: "5:00", 
		AnswerD: "2:00",
	},
	{
		ID: "5",
		Difficulty: "easy", 
		Title: "Which of these Pokemon would be considered a \"Sniper\"?", 
		AnswerA: "Espeon", 
		AnswerB: "Pikachu", 
		AnswerC: "Decidueye", 
		AnswerD: "Sylveon",
	},
	{
		ID: "6", 
		Difficulty: "easy", 
		Title: "What does the word \"Meta\" mean?", 
		AnswerA: "It isn't used as a term in this game.", 
		AnswerB: "It means a Pokemon is completely broken.", 
		AnswerC: "It means something is popular.", 
		AnswerD: "It stands for \"Most Effective Tactics Available\", meaning that the strategy is the most optimal way to win or has the best performance for a specific task.",
	},
	{
		ID: "7", 
		Difficulty: "easy", 
		Title: "Which Pokemon is NOT a Special Attacker?", 
		AnswerA: "Cramorant",
		AnswerB: "Cinderace",
		AnswerC: "Lapras", 
		AnswerD: "Mew", 
	},
	{
		ID: "8", 
		Difficulty: "easy", 
		Title: "What does 'CC' mean?", 
		AnswerA: "It stands for \"Crowd Control\" it is the general implication that when something is \"CC'd\", it is incapable of acting or moving.",
		AnswerB: "It doesn't mean anything.", 
		AnswerC: "It stands for \"Cringe Carries\", specifying pokemon that are extemely broken and that many people use to carry games easily.", 
		AnswerD: "It is something that has no relation within MOBA's.",
	},
	{
		ID: "9", 
		Difficulty: "easy", 
		Title: "What does \"fog of war\" describe?", 
		AnswerA: "It is the idea that you gain \"vision\" whenever an enemy is near one of your bases, or near you or one of your allies; otherwise you cannot see where your enemies are at.", 
		AnswerB: "It describes a literal fog within the game.", 
		AnswerC: "It describes the clouds in the background of the map that you can see near the edge.", 
		AnswerD: "It doesn't mean anything within MOBA's", 
	},
	{
		ID: "10",
		Difficulty: "easy", 
		Title: "Which one of these is NOT an Attacker?", 
		AnswerA: "Cinderace", 
		AnswerB: "Eldegoss", 
		AnswerC: "Inteleon", 
		AnswerD: "Gardevoir",
	},
}

var mQuestions = []question{
	{
		ID: "1", 
		Difficulty: "medium", 
		Title: "At what time do the first 2 Objectives spawn (Regieleki, Registeel, etc.)", 
		AnswerA: "7:30", 
		AnswerB: "7:00", 
		AnswerC: "8:00", 
		AnswerD: "5:00",
	},
	{
		ID: "2", 
		Difficulty: "medium", 
		Title: "How many times can you stack a stackable item (Attack Weight, Special Attack Specs, Aeos Cookie) at maximum?", 
		AnswerA: "10", 
		AnswerB: "5", 
		AnswerC: "6", 
		AnswerD: "9",
	},
	{
		ID: "3", 
		Difficulty: "medium", 
		Title: "Which Pokemon has the lowest base UNITE Move recharge time (Not counting Dragapult or Blaziken)? ", 
		AnswerA: "Delphox", 
		AnswerB: "Pikachu", 
		AnswerC: "Azumarill", 
		AnswerD: "Leafeon",
	},
	{
		ID: "4", 
		Difficulty: "medium", 
		Title: "Are all Ranged classified Pokemon Special Attackers?", 
		AnswerA: "Yes, there is no difference between Special Attack stats and Physical Attack stats.", 
		AnswerB: "Yes, all Ranged Pokemon only use Special Attack stats.", 
		AnswerC: "No, Pokemon like Cinderace, Decidueye, and Dragapult are Ranged Pokemon but use Physical Attack stats.", 
		AnswerD: "No, all Pokemon only have an Attack stat, there is no such thing as Special Attack or Physical Attack.",
	},
	{
		ID: "5", 
		Difficulty: "medium", 
		Title: "What does the Rayquaza Shield do when you successfuly taken down Rayquaza?", 
		AnswerA: "The shield does nothing. It is only aesthetic.", 
		AnswerB: "The shield is given to all living teammates, which increases scoring speed by 300%, prevents the enemy team from stopping you while scoring while the shield is active, gives a damage boost of 40%, and a shield equal to 30% of your HP + 3000 HP.", 
		AnswerC: "It makes your team slower, and does nothing to help you win the game.", 
		AnswerD: "It makes all the enemy goals defenseless for a certain amount of time, allowing you and your teammates to instantly score.",
	},
	{
		ID: "6", 
		Difficulty: "medium", 
		Title: "What are the ONLY 2 pokemon that can prevent a Pokemon from scoring with a Rayquaza shield by dealing damage?", 
		AnswerA: "Metagross is the only Pokemon who can prevent scoring, even with a Rayquaza shield.", 
		AnswerB: "Blaziken and Dragapult. Both do a lot of damage that break shields instantly.", 
		AnswerC: "Urshifu and Tyranitar. Both do true damage (damage that penetrates shields), completely bypassing the effect the shield has when preventing score stopping.", 
		AnswerD: "Mamoswine and Ninetales. Mamoswine slows scoring speed when using ice moves on its own goal, while Ninetales does a lot of CC that breaks through the shield.",
	},
	{
		ID: "7", 
		Difficulty: "medium", 
		Title: "What does Hoopa's Hyperspace Hole do?", 
		AnswerA: "It creates a large ring in a designated location which allies can use to return to base to heal, then warp back manually within a certain amount of alloted time.", 
		AnswerB: "It creates a death zone for allies that will insta-KO anything that touches it, it should be avoided at all costs.", 
		AnswerC: "It only does damage, nothing else.", 
		AnswerD: "It creates a massive hole which will send enemies back to base when they step onto it, forcing them to travel back to their previous location from base.",
	},
	{
		ID: "8", 
		Difficulty: "medium", 
		Title: "Which Pokemon has a UNITE Move that allows another ally to ride on their back?", 
		AnswerA: "Metagross", 
		AnswerB: "Garchomp", 
		AnswerC: "Machamp", 
		AnswerD: "Lapras",
	},
	{
		ID: "9", 
		Difficulty: "medium", 
		Title: "Why is it important to have a Defender on your team?", 
		AnswerA: "Defenders provide a lot of damage to your team but are very easy to KO. They don't provide a lot of CC but are very mobile, allowing them to easily infiltrate the enemy back line and KO the enemy carries.", 
		AnswerB: "Defenders provide nothing of value for your team, and you should never play one. it is a useless role that provides nothing but a free win for the enemy team.", 
		AnswerC: "Defenders provide a strong frontline that creates a wall between the enemy team and your Attackers. Defenders in general have high HP and Defenses, allowing them to take a lot of damage meanwhile providing meaningful CC and distance for your damage carries. They are a crucial part of any team.", 
		AnswerD: "It is not important to have a Defender. For the most part Defenders are optional since they provide nothing more than what an All-Rounder can provide. It is a throw to pick a Defender when you could instead play an All-rounder.",
	},
	{
		ID: "10", 
		Difficulty: "medium", 
		Title: "Which Pokemon is NOT considered an EX Pokemon?", 
		AnswerA: "Urshifu", 
		AnswerB: "Mewtwo X", 
		AnswerC: "Mewtwo Y", 
		AnswerD: "Zacian",
	},
}

var hQuestions = []question{
	{
		ID: "1", 
		Difficulty: "hard", 
		Title: "What is the most common and meta competitive setup for a team?", 
		AnswerA: "1 Jungler (Attacker, Speedster, or All-rounder), 1 Defender, 1 Support, 1 Top Laner (All-rounder, Speedster, or Defender), 1 Attacker ", 
		AnswerB: "5 Attackers", 
		AnswerC: "2 Junglers, 3 Bot laners", 
		AnswerD: "5 Supporters",
	},
	{
		ID: "2", 
		Difficulty: "hard", 
		Title: "Why is Exp Share so crucial on Supporters and Defenders?", 
		AnswerA: "It is not a crucial item at all. There is no reason to use the item, it does not provide damage therefore it is useless.", 
		AnswerB: "Exp Share provides a massive advantage to your team as it provides the damage laners on both Top and Bottom lane extra Exp that otherwise would not exist (130% total versus 100%). In the current meta, it is a crucial item needed to keep up with the Exp gain of the enemy team.", 
		AnswerC: "It is not a crucial item. Other items like Buddy Barrier or Resonance Guard provide much more value to your team versus Exp Share.", 
		AnswerD: "It is crucial because it takes your Pokemon's attack stat and gives it to nearby allies to deal more damage to the enemies. It is a crucial item in the current meta to keep up with the damage gain of the enemy team.",
	},
	{
		ID: "3", 
		Difficulty: "hard", 
		Title: "Why is the bottom objective in general more valuable at the beginning of the game?", 
		AnswerA: "Because it provides a large chunk of Exp to your entire team that can give you a strong early to mid game lead over your enemies, allowing you to snowball the game and keep a level advantage. It also provides a buff to all alive teammates based on the Objective that spawned, the best being Registeel.", 
		AnswerB: "Bottom objective is not at all important. It provides nothing of value and is best to ignore it as you waste a lot of resources trying to take it.", 
		AnswerC: "The top objective, Regieleki is more important because it allows you to push the enemies goal with essentially a 6th ally. It is an easy way to get a strong early to mid game lead that allows you to snowball your enemies.", 
		AnswerD: "Both objectives are equally important. You should split the team in between both objectives for the best chance to secure both.",
	},	{
		ID: "4", 
		Difficulty: "hard", 
		Title: "What does the Jungle Blue/Purple Buff do?", 
		AnswerA: "It provides a 30% slow when attacking enemies with a basic attack.", 
		AnswerB: "It provides a 10% cooldown reduction to your moves for 70 seconds. ", 
		AnswerC: "It provides a 20% cooldown reduction to your moves for 60 seconds.", 
		AnswerD: "It provides a Defense and Special Defense increase equal to 20% for 30 seconds.",
	},	{
		ID: "5", 
		Difficulty: "hard", 
		Title: "What is an \"Execute\"?", 
		AnswerA: "It is the term used to mean when you get knocked out by a piece of farm or objective.", 
		AnswerB: "An Execute is a massive amount of damage dealt to an enemy, usually a percentage, based on the enemies' missing HP. It is a substantially effective move meant to execute, as is the name, enemies that are at low HP thresholds.", 
		AnswerC: "It is damage dealt to an enemy based on their defenses. Executes deal more damage the lower the Defense and Special Defense stats of the enemy are.", 
		AnswerD: "Execute is a term used in other MOBAs, but not in Unite.",
	},	{
		ID: "6", 
		Difficulty: "hard", 
		Title: "Which Pokemon move has the highest Execute damage?", 
		AnswerA: "Buzzswole's Ultra Swole Slam UNITE Move", 
		AnswerB: "Urshifu's Wicked Blow Move", 
		AnswerC: "Metagross' Compute and Crush UNITE Move", 
		AnswerD: "Meowscarada's Flower Trick Move",
	},	{
		ID: "7", 
		Difficulty: "hard", 
		Title: "What is a \"secure\"?", 
		AnswerA: "The term secure does not exist within Unite. ", 
		AnswerB: "A secure describes generally a move that deals a large amount of damage within a relatively small amount of time, meant to be used as a \"last hit\" for objectives or farm. As is its name, it is mean to secure that farm or objective for your team, at a relative guarantee to prevent it from being stolen.", 
		AnswerC: "A secure is the idea of securing your bases by keeping at least one ally on each base (both and top) at all times to prevent enemies from scoring.", 
		AnswerD: "A secure is the idea of dealing large amounts of damage on an enemy Pokemon, so that you gain the most Exp from that KO.",
	},	{
		ID: "8", 
		Difficulty: "hard", 
		Title: "What does \"peeling\" mean?", 
		AnswerA: "Peeling doesn't mean anything.", 
		AnswerB: "Peeling is another way to refer to baiting. Essentially the ability for a front-liner to bait enemies within range of your back line to get easy KOs.", 
		AnswerC: "Peeling refers to the ability to \"peel\" the front line from the back line, separating them from each other to make it easier to KO the weaker enemies. Generally this is done by a Speester or fast All-rounder, who attacks from behind and casues the enemies to scramble and break their formation.", 
		AnswerD: "Peeling is the tactic of going in front of enemies or taking damage instead of your allies when they are at risk of being KO'd. The point is to take the damage for them, and prevent the enemies from being able to KO your ally, while they run to safety. Generally this can be done by CCing or slowing the enemies with your skills that prevent them from catching up to your ally.",
	},	{
		ID: "9", 
		Difficulty: "hard", 
		Title: "Why is pushing a goal at 2:30 or lower before taking Rayquaza in most cases considered a bad play?", 
		AnswerA: "It is not considered a bad play. It is the most common play in competitive that generally results in a team already snowballing, snowballing even more. Being able to successfully pull this off means you can ignore Rayquaza and just focusing on defending your own goals or Rayquaza.", 
		AnswerB: "Because the negative return generally outweighs any positives gained. Pushing a goal usually means using UNITE moves to push the enemies off of it, meaning you won't have them for Rayquaza. If your entire team is wiped, you are likely to lose Rayquaza as the enemies will likely immediately burn it uncontested. You are feeding the enemies a massive amount of comeback Exp, which could turn a substantial lead into becoming completely behind at the most crucial point in the game. At most, you would be able to break the single goal with a substantial overcap, but the odds of this occuring are much worse than everything you are likely to lose.", 
		AnswerC: "It is not a bad play because many Master players do it. If the top players in the game are doing it, it must mean there is a reason behind it. \"Back-capping\" is a popular method of scoring during double points (2:00) to give your team a substantial lead over your enemies. It just depends on how you do it and when. Generally the best time is when your team is taking a major team fight that can likely decide the game near Rayquaza.", 
		AnswerD: "Many players are just not skilled enough to pull off a 2:30 push. That is the only reason why it is a bad play, because the skill issues on your team mean it'll never work.",
	},	{
		ID: "10", 
		Difficulty: "hard", 
		Title: "Why is map awareness such a crucial skill within Unite?", 
		AnswerA: "What is a map? I actively ignore anything happening on the top left corner of my screen, it is unimportant to my farming simulator.", 
		AnswerB: "Map awareness is a crucial skill because of how much information you can get from paying attention to the vision you have around the entire map. It allows you to keep track of all ally positions, enemy positions, farm spawned, and objectives up, allowing you to make informed and greater decisions based on the information you can acquire from the map. Generally, it is recommended that you should look at your map about half the time you are playing so that you are completely aware of all your surroundings and can make informed decisions.", 
		AnswerC: "Map awareness is not a crucial skill. Generally there is too much happening in the game for you to keep track of the map on top of everything else, so its generally better to only look at the map when you aren't actively doing something.", 
		AnswerD: "Looking at your map is important, but it is not everything. You should be paying attention more to the fights occuring on your screen rather than constantly looking at your map. The information gained is not worth the risk when you have to actively look away from what is in your immediate vecinity.",
	},
}

var eAnswers = []correctAnswer{
	{
		ID: 1,
		CorrectAns: "C",
	},
	{
		ID: 2,
		CorrectAns: "A",
	},
	{
		ID: 3,
		CorrectAns: "D",
	},
	{
		ID: 4,
		CorrectAns: "B",
	},
	{
		ID: 5,
		CorrectAns: "C",
	},
	{
		ID: 6,
		CorrectAns: "D",
	},
	{
		ID: 7,
		CorrectAns: "B",
	},
	{
		ID: 8,
		CorrectAns: "A",
	},
	{
		ID: 9,
		CorrectAns: "A",
	},
	{
		ID: 10,
		CorrectAns: "B",
	},
}

var mAnswers = []correctAnswer{
	{
		ID: 1,
		CorrectAns: "B",
	},
	{
		ID: 2,
		CorrectAns: "C",
	},
	{
		ID: 3,
		CorrectAns: "A",
	},
	{
		ID: 4,
		CorrectAns: "C",
	},
	{
		ID: 5,
		CorrectAns: "B",
	},
	{
		ID: 6,
		CorrectAns: "C",
	},
	{
		ID: 7,
		CorrectAns: "A",
	},
	{
		ID: 8,
		CorrectAns: "D",
	},
	{
		ID: 9,
		CorrectAns: "C",
	},
	{
		ID: 10,
		CorrectAns: "A",
	},
}

var hAnswers = []correctAnswer{
	{
		ID: 1,
		CorrectAns: "A",
	},
	{
		ID: 2,
		CorrectAns: "B",
	},
	{
		ID: 3,
		CorrectAns: "A",
	},
	{
		ID: 4,
		CorrectAns: "B",
	},
	{
		ID: 5,
		CorrectAns: "B",
	},
	{
		ID: 6,
		CorrectAns: "A",
	},
	{
		ID: 7,
		CorrectAns: "B",
	},
	{
		ID: 8,
		CorrectAns: "D",
	},
	{
		ID: 9,
		CorrectAns: "B",
	},
	{
		ID: 10,
		CorrectAns: "B",
	},
}

func getEQuestions(q *gin.Context) {
	q.IndentedJSON(http.StatusOK, eQuestions)
}

func getCurrentQuestion(q *gin.Context) {
	diff := q.Param("diff")

	id := q.Param("id")

	question, err := getQuestionByID(id, diff)
	
	if err != nil {
		q.IndentedJSON(http.StatusNotFound, gin.H{"message": "question not found."})
	}

	q.IndentedJSON(http.StatusOK, question)
}

func getTotalQuestions(q *gin.Context) {
	diff := q.Param("diff")
	println(diff)
	var totalQuestions int;
	switch diff {
	case "easy":
		totalQuestions = len(eQuestions)
		break
	case "medium":
		totalQuestions = len(mQuestions)
		break
	case "hard":
		totalQuestions = len(hQuestions)
		break
	}
	println(totalQuestions)
	q.IndentedJSON(http.StatusOK, totalQuestions)
}

func getScore(q *gin.Context) {
	var answer userAnswer
	q.BindJSON(&answer)
	
	answerID, err := strconv.Atoi(answer.AnswerID)
	if err != nil {
		println("Cannot convert id to Int.")
	}

	newScore := validateScore(answerID, answer.Answer, answer.Difficulty, answer.CurrentScore)

	q.JSON(http.StatusOK, newScore)
}

func validateScore(ansID int, ans string, diff string, scr int) int {
	switch diff {
	case "easy":
		for _, a := range eAnswers {
			if ansID == a.ID {
				if ans == a.CorrectAns {
					scr++
					break
				}
			}
		}
		break

	case "medium":
		for _, a := range mAnswers {
			if ansID == a.ID {
				if ans == a.CorrectAns {
					scr++
					break
				}
			}
		}
		break

	case "hard":
		for _, a := range hAnswers {
			if ansID == a.ID {
				if ans == a.CorrectAns {
					scr++
					break
				}
			}
		}
		break
	}

	return scr
}

func getQuestionByID(id string, diff string) (*question, error) {
	switch diff {
	case "easy":
		for i, q := range eQuestions {
			if q.ID == id {
				return &eQuestions[i], nil
			}
		}
		break
	case "medium":
		for i, q := range mQuestions {
			if q.ID == id {
				return &mQuestions[i], nil
			}
		}
		break
	case "hard":
		for i, q := range hQuestions {
			if q.ID == id {
				return &hQuestions[i], nil
			}
		}
		break
	}

	return nil, errors.New("question not found")
}

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "https://unite-quiz-dhza.vercel.app")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
	
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	
		c.Next()
	})
	
	router.GET("/questions/:diff/:id", getCurrentQuestion)
	router.GET("/totalquestions/:diff", getTotalQuestions)
	router.POST("/questions/score", getScore)

	port := ":" + os.Getenv("PORT")
	router.Run(port)
}