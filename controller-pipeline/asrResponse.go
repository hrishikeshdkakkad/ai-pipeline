package main

//--------------------------------------------------------------------- Speaker Diarization-----------------------------------------------------------------------------------------------------------------------------------

type DiarizationAudioInputData struct {
	AudioFile    string `json:"audio_file"`
	SampleRate   int    `json:"sample_rate"`
	LanguageCode string `json:"language_code"`
	SpeakerCount int    `json:"speaker_count"`
	Diarization  struct {
		EnableSpeakerDiarization bool    `json:"enable_speaker_diarization"`
		MaxSpeakerTurns          int     `json:"max_speaker_turns"`
		MinSpeakerTurns          int     `json:"min_speaker_turns"`
		SpeechTurnsThreshold     float64 `json:"speech_turns_threshold"`
		SilenceTurnsThreshold    float64 `json:"silence_turns_threshold"`
		Model                    string  `json:"model"`
		UseEnhanced              bool    `json:"use_enhanced"`
	} `json:"diarization_config"`
}

type SpeakerSegment struct {
	SpeakerID int     `json:"speaker_id"`
	StartTime float64 `json:"start_time"`
	EndTime   float64 `json:"end_time"`
}

type DiarizationResults struct {
	SpeakerSegments []SpeakerSegment `json:"speaker_segments"`
}

//type DiarizationAudioOutputData struct {
//	AudioFile         string             `json:"audio_file"`
//	SpeakerCount      int                `json:"speaker_count"`
//	DiarizationResult DiarizationResults `json:"diarization_results"`
//}

type DiarizationAudioOutputData struct {
	AudioFile          string `json:"audio_file"`
	SpeakerCount       int    `json:"speaker_count"`
	DiarizationResults struct {
		SpeakerSegments []struct {
			SpeakerId int     `json:"speaker_id"`
			StartTime float64 `json:"start_time"`
			EndTime   float64 `json:"end_time"`
		} `json:"speaker_segments"`
	} `json:"diarization_results"`
}

// ---------------------------------------------------------------------END OF SPEAKER DIARIZATION-----------------------------------------------------------------

//--------------------------------------------------------------------- Speaker CLASSIFIER-----------------------------------------------------------------------------------------------------------------------------------

type DiarizationResult struct {
	SpeakerSegments []struct {
		SpeakerID   int     `json:"speaker_id"`
		StartTime   float64 `json:"start_time"`
		EndTime     float64 `json:"end_time"`
		SpeakerType string  `json:"speaker_type"`
	} `json:"speaker_segments"`
}

type SpeakerClassifierOutputData struct {
	AudioFile          string            `json:"audio_file"`
	SpeakerCount       int               `json:"speaker_count"`
	DiarizationResults DiarizationResult `json:"diarization_results"`
}

// ---------------------------------------------------------------------END OF SPEAKER CLASSIFIER-----------------------------------------------------------------

//-------------------------------------------------------------------------------ASR-----------------------------------------------------------------------------------------------------------------------------------

type SpeechSegment struct {
	StartTime     float64 `json:"start_time"`
	EndTime       float64 `json:"end_time"`
	Transcription string  `json:"transcription"`
}

type SpeechSegmentResults struct {
	SpeakerID      int             `json:"speaker_id"`
	StartTime      float64         `json:"start_time"`
	EndTime        float64         `json:"end_time"`
	SpeakerType    string          `json:"speaker_type"`
	SpeechSegments []SpeechSegment `json:"speech_segments"`
}

type SpeakerSegmentResults struct {
	SpeakerSegments []SpeechSegmentResults `json:"speaker_segments"`
}

type ASROutputResults struct {
	SpeakerCount       int                   `json:"speaker_count"`
	DiarizationResults SpeakerSegmentResults `json:"diarization_results"`
}

// ---------------------------------------------------------------------END OF ASR-----------------------------------------------------------------------

//----------------------------------------------------------------SPEECH ACTIVITY DETECTION ALGO-----------------------------------------------------------------------------------------------------------------------------------

type SpeechActivitySegmentResults struct {
	SpeakerSegments []struct {
		SpeakerID       int     `json:"speaker_id"`
		StartTime       float64 `json:"start_time"`
		EndTime         float64 `json:"end_time"`
		SpeakerType     string  `json:"speaker_type"`
		SilenceSegments []struct {
			StartTime float64 `json:"start_time"`
			EndTime   float64 `json:"end_time"`
		} `json:"silence_segments"`
	} `json:"speaker_segments"`
}

type SpeechActivityDetectionOutputData struct {
	SpeakerCount       int                          `json:"speaker_count"`
	DiarizationResults SpeechActivitySegmentResults `json:"diarization_results"`
}

// ---------------------------------------------------------------------END OF SPEECH ACTIVITY DETECTION ALGO-----------------------------------------------------------------------

type AsrPipelineResponse struct {
	Asr                     ASROutputResults
	SpeechActivityDetection SpeechActivityDetectionOutputData
	Diarization             DiarizationAudioOutputData
}
