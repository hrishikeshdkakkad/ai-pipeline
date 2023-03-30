package main

type Keywords struct {
	Features []struct {
		Text      string  `json:"text"`
		Relevance float64 `json:"relevance"`
	} `json:"features"`
}

type Segments struct {
	Text       string  `json:"text"`
	StartTime  float64 `json:"start_time"`
	EndTime    float64 `json:"end_time"`
	Confidence float64 `json:"confidence"`
	Speaker    string  `json:"speaker"`
}

type Emotions struct {
	Segments []Segments `json:"segments"`
}

type AudioToText struct {
	Features struct {
		Transcript string  `json:"transcript"`
		Confidence float64 `json:"confidence"`
	} `json:"features"`
}

type NormalResponse struct {
	Keywords    Keywords
	Emotions    Emotions
	AudioToText AudioToText
}
