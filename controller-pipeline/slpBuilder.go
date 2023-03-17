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

func newSlpBuilder() *SlpBuilder {
	return &SlpBuilder{
		keywords:    &Keywords{},
		emotions:    &Emotions{},
		audioToText: &AudioToText{},
	}
}

func (b *SlpBuilder) setKeywordsType(ch chan error) {
	postBody, err := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	response, _ := http.Post("http://localhost:3003/slp/keyword", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}

	fmt.Println(string(body))

	if err := json.Unmarshal(body, &b.keywords); err != nil {
		ch <- fmt.Errorf("error unmarshaling JSON: %v", err)
		return
	}
	ch <- nil
	fmt.Println(response.StatusCode)
	fmt.Println("In setKeywordsType", *b.keywords)
}

func (b *SlpBuilder) setAudioToText(ch chan error) {
	postBody, err := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	response, _ := http.Post("http://localhost:3004/slp/audio-to-text", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
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
	fmt.Println(response.StatusCode)
	fmt.Println("In setAudioToText", *b.audioToText)
}

func (b *SlpBuilder) setEmotions(ch chan error) {
	postBody, err := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	response, _ := http.Post("http://localhost:3002/cv/emotions", "application/json", responseBody)
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
	fmt.Println(response.StatusCode)
	fmt.Println("In setEmotions")
}

func (b *SlpBuilder) getResponse() Response {
	return Response{
		keywords:    *b.keywords,
		audioToText: *b.audioToText,
		emotions:    *b.emotions,
	}
}
