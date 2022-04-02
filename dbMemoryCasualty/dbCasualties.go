package dbMemoryCasualty

import (
	"fmt"
	"sync"
	"time"
	"wecare/entities"
)

//Represents a read-write lock
type Casualty struct {
	Mu   sync.Mutex
	Data map[int]entities.Casualty
}

type ScanHistory struct {
	Mu   sync.Mutex
	Data map[int][]entities.ScanLogs
}

var CasualtyDB = Casualty{Data: map[int]entities.Casualty{}}

// initializes

var CasualtyScanHistory = ScanHistory{
	Data: map[int][]entities.ScanLogs{}, // initializes
}

func UpsertCasualty(c *entities.CasualtyProcessing) {

	tNow := time.Now()
	t := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		tNow.Year(), tNow.Month(), tNow.Day(),
		tNow.Hour(), tNow.Minute(), tNow.Second())

	CasualtyDB.Mu.Lock()
	defer CasualtyDB.Mu.Unlock()

	CasualtyScanHistory.Mu.Lock()
	defer CasualtyScanHistory.Mu.Unlock()

	if _, ok := CasualtyDB.Data[c.ID]; !ok {
		// means Casualty not found
		CasualtyDB.Data[c.ID] = entities.Casualty{
			Name:                    c.Name,
			Nric:                    c.NRIC,
			ID:                      c.ID,
			QRCodeDataCreatedDate:   c.Date,
			MortuaryDataCreatedDate: t,
			LastAction:              "InPro",
			LastActionTime:          t,
			LastActionTeam:          c.TeamMember,
			InPro:                   t,
			InProTeam:               c.TeamMember,
		}

		// means Casualty confirm exists
		fmt.Printf("New Record Created: %+v\n", CasualtyDB.Data[c.ID])

		CasualtyScanHistory.Data[c.ID] = append(CasualtyScanHistory.Data[c.ID], entities.ScanLogs{ActionName: "InPro", ActionTime: t})
		fmt.Println("this is casualty history scan", CasualtyScanHistory.Data[c.ID])

	} else {
		fmt.Printf("Case %v received a new scan\n", c.ID)
		oldData := CasualtyDB.Data[c.ID]

		switch c.LastAction {
		case "InPro":
			InPro(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "BodyCleanIn":
			OutPro(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "BodyCleanOut":
			BodyCleanOut(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "BodyDressIn":
			BodyDressIn(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "BodyDressOut":
			BodyDressOut(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "EcoffinIn":
			EcoffinIn(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "EcoffinOut":
			EcoffinOut(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "PeGoStorage":
			PeGoStorage(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "PeInPro":
			PeInPro(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "PeIntoStorage":
			PeIntoStorage(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "PeOutPro":
			PeOutPro(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "PeOutofStorage":
			PeOutofStorage(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "ProcessedCleanIn":
			ProcessedCleanIn(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "ProcessedCleanOut":
			ProcessedCleanOut(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "TransportIn":
			TransportIn(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "TransportOut":
			TransportOut(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "UnprocessedOut":
			UnprocessedOut(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "UnprocessedIn":
			UnprocessedIn(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		case "OutPro":
			OutPro(&oldData, c.TeamMember, t)
			CasualtyDB.Data[c.ID] = oldData
			fmt.Printf("new updated casualty information %+v \n", CasualtyDB.Data[c.ID])

		default:
			fmt.Println("unknown case in casualty update switch detected")

		} // end case

		CasualtyScanHistory.Data[c.ID] = append(CasualtyScanHistory.Data[c.ID], entities.ScanLogs{ActionName: c.LastAction, ActionTime: t})
		fmt.Println("this is casualty history scan", CasualtyScanHistory.Data[c.ID])
	}
	return
}

func generateStructFields(role string) string {
	return role + "Team"
}

func InPro(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "InPro"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.InPro = t
	c.InProTeam = teamMember
}

func OutPro(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "OutPro"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.OutPro = t
	c.OutProTeam = teamMember
}

func BodyCleanIn(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "BodyCleanIn"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.BodyCleanIn = t
	c.BodyCleanInTeam = teamMember
}

func BodyCleanOut(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "BodyCleanOut"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.BodyCleanOut = t
	c.BodyCleanOutTeam = teamMember
}

func BodyDressIn(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "BodyDressIn"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.BodyDressIn = t
	c.BodyDressInTeam = teamMember
}

func BodyDressOut(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "BodyDressOut"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.BodyDressOut = t
	c.BodyDressOutTeam = teamMember
}

func EcoffinIn(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "EcoffinIn"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.EcoffinIn = t
	c.EcoffinInTeam = teamMember
}

func EcoffinOut(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "EcoffinOut"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.EcoffinOut = t
	c.EcoffinOutTeam = teamMember
}

func PeGoStorage(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "PeGoStorage"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.PeGoStorage = t
	c.PeGoStorageTeam = teamMember
}
func PeIntoStorage(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "PeIntoStorage"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.PeIntoStorage = t
	c.PeIntoStorageTeam = teamMember
}

func PeOutofStorage(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "PeOutofStorage"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.PeOutofStorage = t
	c.PeOutofStorageTeam = teamMember
}

func PeInPro(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "PeInPro"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.PeInPro = t
	c.PeInProTeam = teamMember
}

func PeOutPro(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "PeOutPro"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.PeOutPro = t
	c.PeOutProTeam = teamMember
}

func ProcessedCleanIn(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "ProcessedCleanIn"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.ProcessedCleanIn = t
	c.ProcessedCleanInTeam = teamMember
}

func ProcessedCleanOut(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "ProcessedCleanOut"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.ProcessedCleanOut = t
	c.ProcessedCleanOutTeam = teamMember
}

func TransportIn(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "TransportIn"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.TransportIn = t
	c.TransportInTeam = teamMember
}

func TransportOut(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "TransportOut"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.TransportOut = t
	c.TransportOutTeam = teamMember
}

func UnprocessedIn(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "UnprocessedIn"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.UnprocessedIn = t
	c.UnprocessedInTeam = teamMember
}

func UnprocessedOut(c *entities.Casualty, teamMember, t string) {
	c.LastAction = "UnprocessedOut"
	c.LastActionTime = t
	c.LastActionTeam = teamMember
	c.UnprocessedOut = t
	c.UnprocessedOutTeam = teamMember
}

func SearchCasualty(caseID int) *entities.Casualty {
	CasualtyDB.Mu.Lock()
	defer CasualtyDB.Mu.Unlock()
	if _, ok := CasualtyDB.Data[caseID]; !ok {
		return nil
	}

	data := CasualtyDB.Data[caseID]
	return &data
} // end func

func SearchCasualtyScanLogs(caseID int) []entities.ScanLogs {
	CasualtyDB.Mu.Lock()
	defer CasualtyDB.Mu.Unlock()
	if _, ok := CasualtyDB.Data[caseID]; !ok {
		return nil
	}

	return CasualtyScanHistory.Data[caseID]
}

func AllCasualtyLogs() map[int]entities.Casualty {
	CasualtyDB.Mu.Lock()
	defer CasualtyDB.Mu.Unlock()
	return CasualtyDB.Data
} // end func

// function clears all casualty logs
func DeleteAllCasualtyLogs() {
	CasualtyDB.Mu.Lock()
	defer CasualtyDB.Mu.Unlock()
	CasualtyDB.Data = map[int]entities.Casualty{} // initializes

} // end func
