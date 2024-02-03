package api

type Player struct {
	Id           int    `json:"id"`
	Number       int    `json:"number"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Trnsfermarkt int    `json:"market_value"`
	DateOfBirth  string `json:"date_of_birth"`
	Nation       string `json:"nation"`
	Position     string `json:"position"`
}

var Squad = []Player{
	{1, 1, "Thibaut", "Courtois", 35000000, "11-05-1992", "Belgium", "GK"},
	{2, 2, "Daniel", "Carvajal", 12000000, "11-01-1992", "Spain", "RB"},
	{3, 3, "Eder", "Militao", 70000000, "18-01-1998", "Brazil", "CB"},
	{4, 4, "David", "Alaba", 30000000, "24-06-1992", "Austria", "CB"},
	{5, 5, "Jude", "Bellingham", 180000000, "29-06-2003", "England", "CAM"},
	{6, 6, "Nacho", "Fernandez", 5000000, "18-01-1990", "Spain", "CB"},
	{7, 7, "Vinicius", "Junior", 150000000, "12-07-2000", "Brazil", "LW"},
	{8, 8, "Toni", "Kroos", 14000000, "04-01-1990", "Germany", "CM"},
	{9, 10, "Luka", "Modric", 10000000, "09-09-1985", "Croatia", "CM"},
	{10, 11, "Rodrigo", "de Goes", 100000000, "09-01-2001", "Brazil", "RW"},
	{11, 12, "Eduardo", "Camavinga", 90000000, "10-10-2002", "France", "CM"},
	{12, 13, "Andriy", "Lunin", 8000000, "11-02-1999", "Ukraine", "GK"},
	{13, 14, "Joselu", "Mato", 5000000, "27-03-1990", "Spain", "ST"},
	{14, 15, "Federico", "Valverde", 100000000, "22-07-1998", "Uruguay", "CM"},
	{15, 17, "Lucas", "Vasquez", 5000000, "01-06-1991", "Spain", "RB"},
	{16, 18, "Aurelien", "Tchouameni", 90000000, "27-01-2000", "France", "CDM"},
	{17, 19, "Dani", "Ceballos", 9000000, "07-08-1996", "Spain", "CM"},
	{18, 20, "Fran", "Garcia", 15000000, "14-08-1999", "Spain", "LB"},
	{19, 21, "Brahim", "Diaz", 20000000, "03-08-1999", "Spain", "CAM"},
	{20, 22, "Antonio", "Rudiger", 25000000, "03-03-1993", "Germany", "CB"},
	{21, 23, "Ferland", "Mendy", 20000000, "08-06-1995", "France", "LB"},
	{22, 24, "Arda", "Guler", 15000000, "25-02-2005", "Spain", "RB"},
	{23, 25, "Kepa", "Arrizabalaga", 20000000, "03-10-1994", "Spain", "GK"},
	{24, 32, "Nico", "Paz", 10000000, "08-09-2004", "Argentina", "CM"},
}
