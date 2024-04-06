package model

type Event struct {
	Id           string `json:"id"`
	EventName    string `json:"event_name"`
	FromDate     string `json:"from_date"`
	ToDate       string `json:"to_date"`
	EventLocation     string `json:"event_location"`
	Descriptions string `json:"descriptions"`
	// PersonAdded     []*User `json:"person_wish_list"`
}

type EventPersonConfirmed struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	IsConfirmed bool   `json:"is_confirmed"`
}

type EveryPerson struct {
	Username string `json:"username"`
}

// CREATE TABLE IF NOT EXISTS events(
//     id VARCHAR(255) NOT NULL,
//     event_name VARCHAR(255) NOT NULL,
//     from_date VARCHAR(255) NOT NULL,
//     to_date VARCHAR(255) NOT NULL,
//     event_location VARCHAR(255) NOT NULL,
//     descriptions VARCHAR(255),
//     PRIMARY KEY (id)
// )

// CREATE TABLE IF NOT EXISTS event_person_confirmed(
//     id VARCHAR(255) NOT NULL,
//     username VARCHAR(255) NOT NULL,
//     is_confirmed BOOLEAN NOT NULL,
//     PRIMARY KEY (event_id, username)
// )