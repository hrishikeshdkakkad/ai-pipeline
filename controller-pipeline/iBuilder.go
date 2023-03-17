package main

type IBuilder interface {
	setKeywordsType(ch chan error)
	setAudioToText(ch chan error)
	setEmotions(ch chan error)
	getResponse() Response
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newSlpBuilder()
	}

	//if builderType == "slpAndEmotions" {
	//	return newSlpBuilder()
	//}

	//if builderType == "igloo" {
	//	return newIglooBuilder()
	//}
	return nil
}
