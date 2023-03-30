package main

type INormalBuilder interface {
	setKeywordsType(ch chan error)
	setAudioToText(ch chan error)
	setEmotions(ch chan error)
	getResponse() NormalResponse
}

func getBuilder(builderType string) INormalBuilder {
	if builderType == "simple" {
		return NewSimpleBuilder()
	}
	//if builderType == "slpAndEmotions" {
	//	return newSlpBuilder()
	//}

	//if builderType == "igloo" {
	//	return newIglooBuilder()
	//}
	return nil
}
