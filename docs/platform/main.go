//	 To parse and unparse do;
//    Interchange, err := UnmarshalInterchange(bytes)
//    bytes, err = Interchange.Marshal()

package main

import "encoding/json"

type Interchange []InterchangeElement

func UnmarshalInterchange(data []byte) (Interchange, error) {
	var r Interchange
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Interchange) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InterchangeElement struct {
	DataElementDelimiter string       `json:"DataElementDelimiter"`
	SegmentDelimiter     string       `json:"SegmentDelimiter"`
	Ta1                  Ta1          `json:"TA1"`
	ISA                  ISA          `json:"ISA"`
	Groups               []Group      `json:"Groups"`
	IEATrailers          []IEATrailer `json:"IEATrailers"`
	Result               Result       `json:"Result"`
}

type Group struct {
	Transactions []Transaction `json:"Transactions"`
	GETrailers   []GETrailer   `json:"GETrailers"`
	Gs           Gs            `json:"GS"`
}

type GETrailer struct {
	NumberOfIncludedSets1 string `json:"NumberOfIncludedSets_1"`
	GroupControlNumber2   string `json:"GroupControlNumber_2"`
}

type Gs struct {
	Time5                           string `json:"Time_5"`
	ReceiverIDCode3                 string `json:"ReceiverIDCode_3"`
	CodeIdentifyingInformationType1 string `json:"CodeIdentifyingInformationType_1"`
	VersionAndRelease8              string `json:"VersionAndRelease_8"`
	SenderIDCode2                   string `json:"SenderIDCode_2"`
	Date4                           string `json:"Date_4"`
	TransactionTypeCode7            string `json:"TransactionTypeCode_7"`
	GroupControlNumber6             string `json:"GroupControlNumber_6"`
}

type Transaction struct {
}

type IEATrailer struct {
	InterchangeControlNumber2 string `json:"InterchangeControlNumber_2"`
	NumberOfIncludedGroups1   string `json:"NumberOfIncludedGroups_1"`
}

type ISA struct {
	InterchangeReceiverID8                  string `json:"InterchangeReceiverID_8"`
	SenderIDQualifier5                      string `json:"SenderIDQualifier_5"`
	AcknowledgementRequested14              string `json:"AcknowledgementRequested_14"`
	SecurityInformation4                    string `json:"SecurityInformation_4"`
	InterchangeDate9                        string `json:"InterchangeDate_9"`
	InterchangeControlStandardsIdentifier11 string `json:"InterchangeControlStandardsIdentifier_11"`
	ComponentElementSeparator16             string `json:"ComponentElementSeparator_16"`
	SecurityInformationQualifier3           string `json:"SecurityInformationQualifier_3"`
	AuthorizationInformationQualifier1      string `json:"AuthorizationInformationQualifier_1"`
	InterchangeControlNumber13              string `json:"InterchangeControlNumber_13"`
	InterchangeSenderID6                    string `json:"InterchangeSenderID_6"`
	UsageIndicator15                        string `json:"UsageIndicator_15"`
	ReceiverIDQualifier7                    string `json:"ReceiverIDQualifier_7"`
	InterchangeTime10                       string `json:"InterchangeTime_10"`
	AuthorizationInformation2               string `json:"AuthorizationInformation_2"`
	InterchangeControlVersionNumber12       string `json:"InterchangeControlVersionNumber_12"`
}

type Result struct {
	Status    string   `json:"Status"`
	Details   []Detail `json:"Details"`
	LastIndex int64    `json:"LastIndex"`
}

type Detail struct {
	Status           string `json:"Status"`
	Message          string `json:"Message"`
	SegmentID        string `json:"SegmentId"`
	TransactionIndex int64  `json:"TransactionIndex"`
	TransactionRef   string `json:"TransactionRef"`
	Value            string `json:"Value"`
	Index            int64  `json:"Index"`
	DataElementIndex int64  `json:"DataElementIndex"`
}

type Ta1 struct {
	InterchangeDate2               string `json:"InterchangeDate_2"`
	InterchangeAcknowledgmentCode4 string `json:"InterchangeAcknowledgmentCode_4"`
	InterchangeTime3               string `json:"InterchangeTime_3"`
	InterchangeControlNumber1      string `json:"InterchangeControlNumber_1"`
	InterchangeNoteCode5           string `json:"InterchangeNoteCode_5"`
}
