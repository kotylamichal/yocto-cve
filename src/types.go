package main

type Issue struct {
	Id      string `json:"id"`
	Summary string `json:"sumary"`
	ScoreV2 string `json:"scorev2"`
	ScoreV3 string `json:"scorev3"`
	Vector  string `json:"vector"`
	Status  string `json:"status"`
	Link    string `json:"link"`
}

type Packages struct {
	Packages []Package `json:"package"`
}

type Package struct {
	Name    string  `json:"name"`
	Layer   string  `json:"layer"`
	Version string  `json:"version"`
	Issues  []Issue `json:"issue"`
}
