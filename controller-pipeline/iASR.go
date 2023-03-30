package main

type ISimpleAsrBuilder interface {
	setAsrResults()
	setSpeechActivityDetectionResults()
	setDiarization()
	setSpeakerClassifier(DiarizationAudioOutputData)
	getResponse() AsrPipelineResponse
}

func getSimpleAsrBuilder(builderType string) ISimpleAsrBuilder {
	if builderType == "simpleAsr" {
		return NewAsrPipelineBuilder()
	}
	return nil
}
