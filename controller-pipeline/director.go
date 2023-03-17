package main

import (
	"log"
)

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildResponse() Response {
	keywordsCh := make(chan error)
	audioToTextCh := make(chan error)
	emotionsCh := make(chan error)

	go d.builder.setEmotions(emotionsCh)
	go d.builder.setKeywordsType(keywordsCh)
	go d.builder.setAudioToText(audioToTextCh)

	for i := 0; i < 3; i++ {
		select {
		case err := <-keywordsCh:
			if err != nil {
				log.Fatal(err)
			}
		case err := <-audioToTextCh:
			if err != nil {
				log.Fatal(err)
			}
		case err := <-emotionsCh:
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	x := d.builder.getResponse()
	return x

}

//func (d *Director) buildEmotionsAndSLPResponse() Response {
//	d.builder.setEmotions()
//	d.builder.setAudioToText()
//	return d.builder.getResponse()
//}
