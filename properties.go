package azuretexttospeech

// AudioOutput types represent the supported audio encoding formats for the text-to-speech endpoint.
// This type is required when requesting to azuretexttospeech.Synthesize text-to-speed request.
// Each incorporates a bitrate and encoding type. The Speech service supports 24 kHz, 16 kHz, and 8 kHz audio outputs.
// See: https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/rest-text-to-speech#audio-outputs

//go:generate enumer -type=AudioOutput
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
	LocaleafZA  Locale = iota //af-ZA
	LocaleamET                //am-ET
	LocalearAE                //ar-AE
	LocalearBH                //ar-BH
	LocalearDZ                //ar-DZ
	LocalearEG                //ar-EG
	LocalearIQ                //ar-IQ
	LocalearJO                //ar-JO
	LocalearKW                //ar-KW
	LocalearLY                //ar-LY
	LocalearMA                //ar-MA
	LocalearQA                //ar-QA
	LocalearSA                //ar-SA
	LocalearSY                //ar-SY
	LocalearTN                //ar-TN
	LocalearYE                //ar-YE
	LocalebgBG                //bg-BG
	LocalebnBD                //bn-BD
	LocalecaES                //ca-ES
	LocalecsCZ                //cs-CZ
	LocalecyGB                //cy-GB
	LocaledaDK                //da-DK
	LocaledeAT                //de-AT
	LocaledeCH                // de-CH
	LocaledeDE                // de-DE
	LocaleelGR                // el-GR
	LocaleenAU                // en-AU
	LocaleenCA                // en-CA
	LocaleenGB                // en-GB
	LocaleenHK                // en-HK
	LocaleenIE                // en-IE
	LocaleenIN                // en-IN
	LocaleenKE                // en-KE
	LocaleenNG                // en-NG
	LocaleenNZ                // en-NZ
	LocaleenPH                // en-PH
	LocaleenSG                // en-SG
	LocaleenTZ                // en-TZ
	LocaleenUS                // en-US
	LocaleenZA                // en-ZA
	LocaleesAR                // es-AR
	LocaleesBO                // es-BO
	LocaleesCL                // es-CL
	LocaleesCO                // es-CO
	LocaleesCR                // es-CR
	LocaleesCU                // es-CU
	LocaleesDO                // es-DO
	LocaleesEC                // es-EC
	LocaleesES                // es-ES
	LocaleesGQ                // es-GQ
	LocaleesGT                // es-GT
	LocaleesHN                // es-HN
	LocaleesMX                // es-MX
	LocaleesNI                // es-NI
	LocaleesPA                // es-PA
	LocaleesPE                // es-PE
	LocaleesPR                // es-PR
	LocaleesPY                // es-PY
	LocaleesSV                // es-SV
	LocaleesUS                // es-US
	LocaleesUY                // es-UY
	LocaleesVE                // es-VE
	LocaleetEE                // et-EE
	LocalefaIR                // fa-IR
	LocalefiFI                // fi-FI
	LocalefilPH               // fil-PH
	LocalefrBE                // fr-BE
	LocalefrCA                // fr-CA
	LocalefrCH                // fr-CH
	LocalefrFR                // fr-FR
	LocalegaIE                // ga-IE
	LocaleglES                // gl-ES
	LocaleguIN                // gu-IN
	LocaleheIL                // he-IL
	LocalehiIN                // hi-IN
	LocalehrHR                // hr-HR
	LocalehuHU                // hu-HU
	LocaleidID                // id-ID
	LocaleitIT                // it-IT
	LocalejaJP                // ja-JP
	LocalejvID                // jv-ID
	LocalekmKH                // km-KH
	LocalekoKR                // ko-KR
	LocaleltLT                // lt-LT
	LocalelvLV                // lv-LV
	LocalemrIN                // mr-IN
	LocalemsMY                // ms-MY
	LocalemtMT                // mt-MT
	LocalemyMM                // my-MM
	LocalenbNO                // nb-NO
	LocalenlBE                // nl-BE
	LocalenlNL                // nl-NL
	LocaleplPL                // pl-PL
	LocaleptBR                // pt-BR
	LocaleptPT                // pt-PT
	LocaleroRO                // ro-RO
	LocaleruRU                // ru-RU
	LocaleskSK                // sk-SK
	LocaleslSI                // sl-SI
	LocalesoSO                // so-SO
	LocalesuID                // su-ID
	LocalesvSE                // sv-SE
	LocaleswKE                // sw-KE
	LocaleswTZ                // sw-TZ
	LocaletaIN                // ta-IN
	LocaletaLK                // ta-LK
	LocaletaSG                // ta-SG
	LocaleteIN                // te-IN
	LocalethTH                // th-TH
	LocaletrTR                // tr-TR
	LocaleukUA                // uk-UA
	LocaleurIN                // ur-IN
	LocaleurPK                // ur-PK
	LocaleuzUZ                // uz-UZ
	LocaleviVN                // vi-VN
	LocalezhCN                // zh-CN
	LocalezhHK                // zh-HK
	LocalezhTW                // zh-TW
	LocalezuZA                // zu-ZA
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
