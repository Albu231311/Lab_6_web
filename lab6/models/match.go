package models

// Formato principal de cada partido
type Match struct {
	ID          uint   `json:"id" gorm:"primaryKey"` // Identificador único del partido (clave primaria)
	HomeTeam    string `json:"homeTeam"`             // Nombre del equipo local
	AwayTeam    string `json:"awayTeam"`             // Nombre del equipo visitante
	MatchDate   string `json:"matchDate"`            // Fecha del partido
	TotalGoals  int    `json:"totalGoals"`           // Goles totales
	YellowCards int    `json:"yellowCards"`          // Cantidad de tarjetas amarillas
	RedCards    int    `json:"redCards"`             // Cantidad de tarjetas rojas
	ExtraTime   bool   `json:"extraTime"`            // Hubo no hubo tiempo extra
}
