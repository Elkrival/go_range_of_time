package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Closings struct {
	ClosingId           string `json:"closingId, omitempty"`
	CreatedAt           int32  `json:"createdAt, omitempty"`
	UpdatedAt           int32  `json:"updatedAt, omitempty"`
	ContractGrossAmount int32  `json:"contractGrossAmount, omitempty"`
	CloseDate           string `json:"closeDate, omitempty"`
	Address             string `json:"address, omitempty"`
	ListingIDSHA        string `json:"listingIdSHA, omitempty"`
	SideRepresented     string `json:"sideRepresented, omitempty"`
	ListingType         int    `json:"listingType, omitempty"`
	Status              int    `json:"status, omitempty":`
	BrokerFeePercent    struct {
		Percent int64 `json:"percent"`
	} `json:"brokerFee"`
	BrokerFeeAmount              int64         `json:"brokerFeeAmount, omitempty"`
	IsCompassLead                bool          `json:"isCompassLead, omitempty"`
	OffMarketListing             bool          `json:"offMarketListing, omitempty"`
	PropertyType                 int           `json:"propertyType, omitempty"`
	Notes                        string        `json:"notes, omitempty, omitempty"`
	NotesForPrincipal            string        `json:"notesForPrincipal, omitempty, omitempty"`
	Files                        []Files       `json:"files, omitempty"`
	Fees                         []Fees        `json:"fees, omitempty"`
	Allocations                  []Allocations `json:"allocations, omitempty"`
	WithExternalBroker           bool          `json:"withExternalBroker, omitempty"`
	ExternalBrokerageName        string        `json:"externalBrokerageName, omitempty"`
	ExternalBrokerName           string        `json:"externalBrokerName, omitempty"`
	ExternalBrokerEmail          string        `json:"externalBrokerEmail, omitempty"`
	SubmittedAt                  int32         `json:"submittedAt, omitempty"`
	UserID                       string        `json:"userId, omitempty"`
	AgentMasterDataID            string        `json:"agentMasterDataID, omitempty"`
	AgentName                    string        `json:"agentName, omitempty"`
	AgentSplitOverridePercent    int32         `json:"agentSplitOverridePercent, omitempty"`
	AgentSplitOverrideNotes      string        `json:"agentSplitOverrideNotes, omitempty"`
	AgentSplitOverrideReason     string        `json:"agentSplitOverrideReason, omitempty"`
	ZipCode                      string        `json:"zipCode, omitempty"`
	TeamID                       string        `json:"teamId, omitempty"`
	SubmittedByID                string        `json:"submittedByID, omitempty"`
	SubmittedByName              string        `json:"submittedByName, omitempty"`
	SubmittedByAgentMasterDataID string        `json:"submittedByAgentMasterDataID, omitempty"`
	City                         string        `json:"city, omitempty"`
	State                        string        `json:"state, omitempty"`
}
type Fees struct {
	FeeId           int `json:"feeId`
	PercentOrAmount struct {
		Percent int `json:"percent, omitempty"`
	} `json:"percentOrAmount, omitempty"`
	FeeType      int `json:"feeType, omitempty"`
	CounterParty struct {
		Name          string `json:"name, omitempty"`
		Email         string `json:"email, omitempty"`
		BrokerageName string `json:"brokerageName, omitempty"`
	} `json:"counterParty, omitempty"`
	File struct {
		Filename string `json:"filename, omitempty"`
		Path     string `json:"path, omitempty"`
		Url      string `json:"url, omitempty"`
	} `json:"file, omitempty"`
	CalcType string `json:"calcType, omitempty"`
}
type Allocations struct {
	AllocationId       int    `json:"allocationId, omitempty"`
	AgentMasterDataID  string `json:"agentMasterId, omitempty"`
	AgentName          string `json:"agentName, omitempty"`
	PercentageOrAmount struct {
		Percent string `json:"percent, omitempty"`
	} `json:"percentOrAmount, omitempty"`
}
type Files struct {
	Role     string `json:"role, omitempty"`
	Filename string `json:"filename, omitempty"`
	Path     string `json:"path, omitempty"`
	Url      string `json:"url, omitempty"`
}
type ClosingsNodes struct {
	ClosingNodes []Closings `json:"closings, omitempty"`
}

func main() {
	const shortForm = "2006-01-02"
	ed := time.Now().Format(shortForm)
	dseparationDate, _ := time.Parse(shortForm, ed)
	cd, _ := time.Parse(shortForm, "2010-03-16")
	isBefore := cd.Before(dseparationDate)
	fmt.Printf("isBefore=%v\n", isBefore)
	const startDate = "2019-01-01"
	sd, _ := time.Parse(shortForm, startDate)
	isAfter := cd.After(sd)
	fmt.Printf("isAfter=%v\n", isAfter)
	file, _ := ioutil.ReadFile("testdata.json")
	data := ClosingsNodes{}
	newNodes := []Closings{}
	closingsHolder := ClosingsNodes{}
	_ = json.Unmarshal([]byte(file), &data)
	for _, v := range data.ClosingNodes {
		parseCd, _ := time.Parse(shortForm, v.CloseDate)
		if parseCd.Before(dseparationDate) && parseCd.After(sd) {
			newNodes = append(newNodes, v)
		}
	}
	fmt.Printf("len=%d\n", len(data.ClosingNodes))
	fmt.Printf("len=%d\n", len(newNodes))
	closingsHolder.ClosingNodes = newNodes
}
