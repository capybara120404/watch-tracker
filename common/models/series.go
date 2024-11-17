package models

type Series struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Link            string  `json:"link"`
	IMDB            float32 `json:"imdb"`
	StartYear       int     `json:"start_year"`
	EndYear         *int    `json:"end_year,omitempty"`
	PosterLink      string  `json:"poster_link"`
	Country         string  `json:"country"`
	NumberOfEpisode uint    `json:"number_of_episode"`
	EpisodeDuration uint    `json:"episode_duration"`
}

func (series *Series) CalculateTotalDuration() uint {
	return series.NumberOfEpisode * series.EpisodeDuration
}
