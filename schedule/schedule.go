package schedule

type Time struct {
	Hour    string
	Minutes string
}

type Schedule struct {
	Title string
	Table []Time
}

type Ferry struct {
	Title     string
	Schedules []Schedule
}
