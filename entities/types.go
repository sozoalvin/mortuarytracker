package entities

type InPro struct {
	Name string
}

type Casualty struct {
	BodyCleanIn             string
	BodyCleanInTeam         string
	BodyCleanOut            string
	BodyCleanOutTeam        string
	BodyDressIn             string
	BodyDressInTeam         string
	BodyDressOut            string
	BodyDressOutTeam        string
	EcoffinIn               string
	EcoffinInTeam           string
	EcoffinOut              string
	EcoffinOutTeam          string
	ID                      int
	InPro                   string
	InProTeam               string
	LastAction              string
	LastActionTime          string
	LastActionTeam          string
	MortuaryDataCreatedDate string
	Name                    string
	Nric                    string
	OutPro                  string
	OutProTeam              string
	OutproInPECheck         string // confirm functionality again
	OutproInPECheckTeam     string // confirm functionality again
	PeGoStorage             string
	PeGoStorageTeam         string
	PeInPro                 string
	PeInProTeam             string
	PeIntoStorage           string
	PeIntoStorageTeam       string
	PeOutPro                string
	PeOutProTeam            string
	PeOutofStorage          string
	PeOutofStorageTeam      string
	ProcessedCleanIn        string
	ProcessedCleanInTeam    string
	ProcessedCleanOut       string
	ProcessedCleanOutTeam   string
	QRCodeDataCreatedDate   string
	TransportIn             string
	TransportInTeam         string
	TransportOut            string
	TransportOutTeam        string
	UnprocessedInTeam       string
	UnprocessedOut          string
	UnprocessedOutTeam      string
	UnprocessedIn           string
}

type ScanLogs struct {
	ActionName string
	ActionTime string
}

// type Casualty struct {
// 	Name                string
// 	Nric                string
// 	InPro               time.Time
// 	InProTeam           string
// 	UnprocessedIn       time.Time
// 	UnprocessedInTeam   string
// 	UnprocessedOut      time.Time
// 	UnprocessedOutTeam  string
// 	TransportIn         time.Time
// 	TransportInTeam     string
// 	TransportOut        time.Time
// 	TransportOutTeam    string
// 	BodyCleanIn         time.Time
// 	BodyCleanInTeam     string
// 	BodyCleanOut        time.Time
// 	BodyCleanOutTeam    string
// 	PeInPro             time.Time
// 	PeInProTeam         string
// 	peOutPro            time.Time
// 	PeOutProTeam        string
// 	PeGoStorage         time.Time
// 	PeGoStorageTeam     string
// 	PeIntoStorage       time.Time
// 	PeIntoStorageTeam   string
// 	PeOutofStorage      time.Time
// 	BodyDressIn         time.Time
// 	BodyDressInTeam     string
// 	BodyDressOut        time.Time
// 	BodyDressOutTeam    string
// 	ProcessedClean      time.Time
// 	ProcessedCleanTeam  string
// 	EcoffinIn           time.Time
// 	EcoffinInTeam       string
// 	EcoffinOut          time.Time
// 	EcoffinOutTeam      string
// 	OutproInPECheck     time.Time
// 	OutproInPECheckTeam string
// 	OutPro              time.Time
// 	OutProTeam          string
// 	LastAction          string
// 	LastActionTime      time.Time
// }

type CasualtyProcessing struct {
	// ID         string
	ID         int
	NRIC       string
	Name       string
	LastAction string `json:"Role"`
	PE         string
	Date       string
	TeamMember string
}

type UpdateCasualtyInterface interface {
	UpdateCasualty(Casualty) error
}

// func (i InPro) UpdateCasualty(c *Casualty) error {
// 	c.LastAction = i.Name
// 	c.

// }
