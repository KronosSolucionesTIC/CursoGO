package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 //Piedra vence a tijera (tijeras + 1) % 3 = 0
	PAPER    = 1 //Papel vence a la piedra (piedra + 1) % 3 = 1
	SCISSORS = 2 //Tijeras vence a papel (papel + 1) % 3 = 2
)

// Estructura para dar resultado a cada ronda
type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt string `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// Mensajes para cuando gana
var winMessages = []string{
	"Bien hecho",
	"Buen trabajo",
	"Deberias comprar un boleto de loteria",
}

// Mensajes para cuando pierde
var loseMessages = []string{
	"Que lastima",
	"Intentalo de nuevo",
	"Hoy no es tu dia",
}

// Mensajes para cuando empata
var drawMessages = []string{
	"Las grandes mentes piensan igual",
	"Empate",
	"Nadie gana, pero puedes intentarlo de nuevo",
}

// Variables para el puntaje
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	//Mensaje de lo que eligio la computadora
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligio TIJERA"
	}

	//Numero aleatorio para seleccionar mensaje
	messageInt := rand.Intn(3)

	//Variable para mensaje
	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana"
		message = loseMessages[messageInt]
	}

	//Retornar resultado
	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: strconv.Itoa(computerChoiceInt),
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
