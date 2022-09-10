package models

type Game struct {
	gameId  string
	Players []string
}

func NewGame(gameId string) *Game {
	return &Game{gameId: gameId}
}

func (g *Game) AddPlayer(player string) {
	g.Players = append(g.Players, player)
}

func (g *Game) RemovePlayer(username string) {
	for i, player := range g.Players {
		if player == username {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
		}
	}
}
