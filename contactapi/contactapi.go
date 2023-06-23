package contactapi

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

const (
	ContactRelationshipNone           string = "none"
	ContactRelationshipIAddedHim      string = "iAddedHim"
	ContactRelationshipHeAddedMe      string = "heAddedMe"
	ContactRelationshipAddedEachOther string = "addedEachOther"
)

type Relation struct {
	UserID                 string `json:"userId"`
	ContactID              string `json:"contactID"`
	ContactRemarkName      string `json:"contactRemarkName"`
	Relationship           string `json:"relationship"`
	MyRemarkNameForContact string `json:"myRemarkNameForContact"`
}

type ContactApiClient struct {
	client *graphql.Client
}

func (client *ContactApiClient) Init(endpoint string) error {
	if endpoint == "" {
		return fmt.Errorf("contact api client endpoint is empty")
	}
	client.client = graphql.NewClient(endpoint)
	return nil
}

func (client *ContactApiClient) Relation(ctx context.Context, contactID string, token string) (*Relation, error) {
	if client.client == nil {
		return nil, fmt.Errorf("ContactApiClient not init")
	}

	req := graphql.NewRequest(`
	query relation($contactId: ID!){
		relation(contactId: $contactId){
		  userId
		  contactId
			   contactRemarkName
		  myRemarkNameForContact
		  relationship
		}
	  }
	
`)
	req.Var("contactId", contactID)
	req.Header.Set("authorization", "Bearer "+token)
	var response map[string]interface{}
	if err := client.client.Run(ctx, req, &response); err != nil {
		return nil, err
	}
	responseMap, ok := response["relation"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("response parse error")
	}
	relation := &Relation{}
	relation.UserID, ok = responseMap["userId"].(string)
	if !ok {
		return nil, fmt.Errorf("response parse(userId) error")
	}
	relation.ContactID, ok = responseMap["contactId"].(string)
	if !ok {
		return nil, fmt.Errorf("response parse(contactId) error")
	}
	relation.ContactRemarkName, ok = responseMap["contactRemarkName"].(string)
	if !ok {
		return nil, fmt.Errorf("response parse(contactRemarkName) error")
	}
	relation.Relationship, ok = responseMap["relationship"].(string)
	if !ok {
		return nil, fmt.Errorf("response parse(relationship) error")
	}
	relation.MyRemarkNameForContact, ok = responseMap["myRemarkNameForContact"].(string)
	if !ok {
		return nil, fmt.Errorf("response parse(myRemarkNameForContact) error")
	}
	return relation, nil
}

// func (client *ContactApiClient) Contact(ctx context.Context, contactId string, token string) (*Contact, error) {
// 	if client.client == nil {
// 		return nil, fmt.Errorf("ContactApiClient not init")
// 	}

// 	req := graphql.NewRequest(`
// 	query contact($contactId: ID!){
// 		contact(contactId: $contactId){
// 		  id
// 		  remarkName
// 		  relationship
// 		}
// 	  }
// `)
// 	req.Var("contactId", contactId)
// 	req.Header.Set("authorization", "Bearer "+token)
// 	var response map[string]interface{}
// 	if err := client.client.Run(ctx, req, &response); err != nil {
// 		return nil, err
// 	}

// 	contact := &Contact{
// 		ID: contactId,
// 		RemarkName: "",
// 		Relationship: "",
// 	}
// 	contactMap, ok := response["contact"].(map[string]interface{})
// 	if !ok {
// 		return nil, fmt.Errorf("response parse error")
// 	}
// 	contact.RemarkName, ok = contactMap["remarkName"].(string)
// 	if !ok {
// 		return nil, fmt.Errorf("response parse error")
// 	}
// 	contact.Relationship = contactMap["relationship"].(string)
// 	if !ok {
// 		return nil, fmt.Errorf("response parse error")
// 	}
// 	return contact, nil
// }
