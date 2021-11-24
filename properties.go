package azuretexttospeech

// AudioOutput types represent the supported audio encoding formats for the text-to-speech endpoint.
// This type is required when requesting to azuretexttospeech.Synthesize text-to-speed request.
// Each incorporates a bitrate and encoding type. The Speech service supports 24 kHz, 16 kHz, and 8 kHz audio outputs.
// See: https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/rest-text-to-speech#audio-outputs
type AudioOutput int

const (
	AudioRIFF8Bit8kHzMonoPCM AudioOutput = iota
	AudioRIFF16Bit16kHzMonoPCM
	AudioRIFF16khz16kbpsMonoSiren
	AudioRIFF24khz16bitMonoPcm
	AudioRAW8Bit8kHzMonoMulaw
	AudioRAW16Bit16kHzMonoMulaw
	AudioRAW24khz16bitMonoPcm
	AudioSsml16khz16bitMonoTts
	Audio16khz16kbpsMonoSiren
	Audio16khz32kbitrateMonoMp3
	Audio6khz64kbitrateMonoMp3
	Audio16khz128kbitrateMonoMp3
	Audio24khz48kbitrateMonoMp3
	Audio24khz96kbitrateMonoMp3
)

func (a AudioOutput) String() string {
	return []string{"riff-8khz-8bit-mono-mulaw",
		"riff-16khz-16bit-mono-pcm",
		"riff-16khz-16kbps-mono-siren",
		"riff-24khz-16bit-mono-pcm",
		"raw-8khz-8bit-mono-mulaw",
		"raw-16khz-16bit-mono-pcm",
		"raw-24khz-16bit-mono-pcm",
		"ssml-16khz-16bit-mono-tts",
		"audio-16khz-16kbps-mono-siren",
		"audio-16khz-32kbitrate-mono-mp3",
		"audio-16khz-64kbitrate-mono-mp3",
		"audio-16khz-128kbitrate-mono-mp3",
		"audio-24khz-48kbitrate-mono-mp3",
		"audio-24khz-96kbitrate-mono-mp3",
	}[a]
}

// Gender type for the digitized language
//go:generate enumer -type=Gender -linecomment -json
type Gender int

const (
	// GenderMale , GenderFemale are the static Gender constants for digitized voices.
	// See Gender in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices for breakdown
	GenderMale   Gender = iota // Male
	GenderFemale               // Female
)

// Locale references the language or locale for text-to-speech.
// See "locale" in https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/language-support#standard-voices
//go:generate enumer -type=Locale -linecomment -json
type Locale int

const (
	LocaleafZA Locale = iota
	LocaleamET
	LocalearAE
	LocalearBH
	LocalearDZ
	LocalearEG
	LocalearIQ
	LocalearJO
	LocalearKW
	LocalearLY
	LocalearMA
	LocalearQA
	LocalearSA
	LocalearSY
	LocalearTN
	LocalearYE
	LocalebgBG
	LocalebnBD
	LocalecaES
	LocalecsCZ
	LocalecyGB
	LocaledaDK
	LocaledeAT
	LocaledeCH
	LocaledeDE
	LocaleelGR
	LocaleenAU
	LocaleenCA
	LocaleenGB
	LocaleenHK
	LocaleenIE
	LocaleenIN
	LocaleenKE
	LocaleenNG
	LocaleenNZ
	LocaleenPH
	LocaleenSG
	LocaleenTZ
	LocaleenUS
	LocaleenZA
	LocaleesAR
	LocaleesBO
	LocaleesCL
	LocaleesCO
	LocaleesCR
	LocaleesCU
	LocaleesDO
	LocaleesEC
	LocaleesES
	LocaleesGQ
	LocaleesGT
	LocaleesHN
	LocaleesMX
	LocaleesNI
	LocaleesPA
	LocaleesPE
	LocaleesPR
	LocaleesPY
	LocaleesSV
	LocaleesUS
	LocaleesUY
	LocaleesVE
	LocaleetEE
	LocalefaIR
	LocalefiFI
	LocalefilPH
	LocalefrBE
	LocalefrCA
	LocalefrCH
	LocalefrFR
	LocalegaIE
	LocaleglES
	LocaleguIN
	LocaleheIL
	LocalehiIN
	LocalehrHR
	LocalehuHU
	LocaleidID
	LocaleitIT
	LocalejaJP
	LocalejvID
	LocalekmKH
	LocalekoKR
	LocaleltLT
	LocalelvLV
	LocalemrIN
	LocalemsMY
	LocalemtMT
	LocalemyMM
	LocalenbNO
	LocalenlBE
	LocalenlNL
	LocaleplPL
	LocaleptBR
	LocaleptPT
	LocaleroRO
	LocaleruRU
	LocaleskSK
	LocaleslSI
	LocalesoSO
	LocalesuID
	LocalesvSE
	LocaleswKE
	LocaleswTZ
	LocaletaIN
	LocaletaLK
	LocaletaSG
	LocaleteIN
	LocalethTH
	LocaletrTR
	LocaleukUA
	LocaleurIN
	LocaleurPK
	LocaleuzUZ
	LocaleviVN
	LocalezhCN
	LocalezhHK
	LocalezhTW
	LocalezuZA
)

// Region references the locations of the availability of standard voices.
// See https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/regions#standard-voices
type Region int

const (
	// Azure regions and their endpoints that support the Text To Speech service.
	RegionAustraliaEast Region = iota
	RegionBrazilSouth
	RegionCanadaCentral
	RegionCentralUS
	RegionEastAsia
	RegionEastUS
	RegionEastUS2
	RegionFranceCentral
	RegionIndiaCentral
	RegionJapanEast
	RegionJapanWest
	RegionKoreaCentral
	RegionNorthCentralUS
	RegionNorthEurope
	RegionSouthCentralUS
	RegionSoutheastAsia
	RegionUKSouth
	RegionWestEurope
	RegionWestUS
	RegionWestUS2
)

func (t Region) String() string {
	return [...]string{
		"australiaeast",
		"brazilsouth",
		"canadacentral",
		"centralus",
		"eastasia",
		"eastus",
		"eastus2",
		"francecentral",
		"indiacentral",
		"japaneast",
		"japanwest",
		"koreacentral",
		"northcentralus",
		"northeurope",
		"southcentralus",
		"southeastasia",
		"uksouth",
		"westeurope",
		"westus",
		"westus2",
	}[t]

}
