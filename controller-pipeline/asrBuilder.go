package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AsrPipelineBuilder struct {
	asr               *ASROutputResults
	speechActivity    *SpeechActivityDetectionOutputData
	diarization       *DiarizationAudioOutputData
	speakerClassifier *SpeakerClassifierOutputData
}

func NewAsrPipelineBuilder() *AsrPipelineBuilder {
	return &AsrPipelineBuilder{
		asr:               &ASROutputResults{},
		speechActivity:    &SpeechActivityDetectionOutputData{},
		diarization:       &DiarizationAudioOutputData{},
		speakerClassifier: &SpeakerClassifierOutputData{},
	}
}

func (b *AsrPipelineBuilder) setAsrResults() {
	//ch <- nil
}

func (b *AsrPipelineBuilder) setSpeechActivityDetectionResults() {

}

func (b *AsrPipelineBuilder) setDiarization() {
	postBody, err := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	response, httperr := http.Post("http://localhost:3006/speakerDiarization", "application/json", responseBody)
	if httperr != nil {
		b.diarization = &DiarizationAudioOutputData{}
		//ch <- nil
		return
	}
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}
	if err := json.Unmarshal(body, &b.diarization); err != nil {
		//ch <- fmt.Errorf("error unmarshaling JSON: %v", err)
		return
	}

	fmt.Println("Finally here")
	//ch <- nil
}

func (b *AsrPipelineBuilder) setSpeakerClassifier(inp DiarizationAudioOutputData) {
	postBody, err := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	response, httperr := http.Post("http://localhost:3007/speaker-classifier", "application/json", responseBody)
	if httperr != nil {
		b.speakerClassifier = &SpeakerClassifierOutputData{}
		//ch <- nil
		return
	}
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}
	if err := json.Unmarshal(body, &b.speakerClassifier); err != nil {
		//ch <- fmt.Errorf("error unmarshaling JSON: %v", err)
		return
	}
	fmt.Println(b.speakerClassifier, "ddd")
}

func (b *AsrPipelineBuilder) getResponse() AsrPipelineResponse {
	return AsrPipelineResponse{
		Asr:                     *b.asr,
		SpeechActivityDetection: *b.speechActivity,
		Diarization:             *b.diarization,
	}
}
