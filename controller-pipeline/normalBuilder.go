package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SlpBuilder struct {
	keywords    *Keywords
	emotions    *Emotions
	audioToText *AudioToText
}

func NewSimpleBuilder() *SlpBuilder {
	return &SlpBuilder{
		keywords:    &Keywords{},
		emotions:    &Emotions{},
		audioToText: &AudioToText{},
	}
}

func (b *SlpBuilder) setKeywordsType(ch chan error) {
	postBody, err := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	response, httperr := http.Post("http://localhost:3003/slp/keyword", "application/json", responseBody)
	if httperr != nil {
		b.keywords = &Keywords{}
		ch <- nil
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
	if err := json.Unmarshal(body, &b.keywords); err != nil {
		ch <- fmt.Errorf("error unmarshaling JSON: %v", err)
		return
	}
	ch <- nil
}

func (b *SlpBuilder) setAudioToText(ch chan error) {
	postBody, err := json.Marshal(map[string]string{})
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	responseBody := bytes.NewBuffer(postBody)
	response, httperr := http.Post("http://localhost:3004/slp/audio-to-text", "application/json", responseBody)
	if httperr != nil {
		b.audioToText = &AudioToText{}
		ch <- nil
		return
	}
	defer response.Body.Close()
	body, error := io.ReadAll(response.Body)
	if error != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	if err := json.Unmarshal(body, &b.audioToText); err != nil {
		ch <- fmt.Errorf("error unmarshaling JSON: %v", err)
		return
	}
	ch <- nil
}

func (b *SlpBuilder) setEmotions(ch chan error) {
	postBody, err := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	response, httperr := http.Post("http://localhost:3002/cv/emotions", "application/json", responseBody)
	if httperr != nil {
		b.emotions = &Emotions{}
		ch <- nil
		return
	}
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer response.Body.Close()
	body, error := io.ReadAll(response.Body)
	if error != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	if err := json.Unmarshal(body, &b.emotions); err != nil {
		ch <- fmt.Errorf("error unmarshaling JSON: %v", err)
		return
	}
	ch <- nil
}

func (b *SlpBuilder) getResponse() NormalResponse {
	return NormalResponse{
		Keywords:    *b.keywords,
		AudioToText: *b.audioToText,
		Emotions:    *b.emotions,
	}
}
