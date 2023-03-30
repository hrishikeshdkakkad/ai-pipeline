package main

import "fmt"

type AsrDirector struct {
	builder ISimpleAsrBuilder
}

func newAsrDirector(b ISimpleAsrBuilder) *AsrDirector {
	return &AsrDirector{
		builder: b,
	}
}

func (d *AsrDirector) setAsrBuilder(b ISimpleAsrBuilder) {
	d.builder = b
}

func (d *AsrDirector) buildAsrResponse() AsrPipelineResponse {
	//diarizationCh := make(chan error)
	//asrCh := make(chan error)
	//speakerClassifierCh := make(chan error)
	//setSpeechActivityDetectionCh := make(chan error)

	d.builder.setDiarization()
	fmt.Println(d.builder.getResponse().Diarization, "lala")
	d.builder.setSpeakerClassifier(d.builder.getResponse().Diarization)
	//d.builder.setAsrResults(asrCh)
	//d.builder.setSpeechActivityDetectionResults(setSpeechActivityDetectionCh)

	//for i := 0; i < 1; i++ {
	//	select {
	//	case err := <-diarizationCh:
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//case err := <-asrCh:
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//case err := <-setSpeechActivityDetectionCh:
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//case err := <-speakerClassifierCh:
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	}
	//}

	x := d.builder.getResponse()
	return x

}
