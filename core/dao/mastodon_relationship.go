package dao

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// activitypubID.
func ReadRelationship(object, actor string) (*entities.Relationship, error) {

	account := &entities.Account{Uri: object}
	if err := Read(db, account); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	lra2o := &model.LocalRelation{
		Actor:  actor,
		Object: object,
	}
	if err := Read(DB(), lra2o); err != nil {
		// log.Printf("%s", err.Error())
		// return nil, err
	}
	// object to actor
	lro2a := &model.LocalRelation{
		Actor:  object,
		Object: actor,
	}
	if err := Read(DB(), lro2a); err != nil {
		// log.Printf("%s", err.Error())
		// return nil, err
	}

	relationship := &entities.Relationship{
		// Type: String (cast from an integer, but not guaranteed to be a number)
		// Description: The account ID.
		Attributes: (account.Id),
		// Type: Boolean
		// Description: Are you following this user?
		Following: lra2o.Type == model.RelationTypeFollow && lra2o.Status == model.RelationStatusAccepted,
		// Type: Boolean
		// Description: Are you receiving this user’s boosts in your home timeline?
		ShowingReblogs: true,
		// Type: Boolean
		// Description: Have you enabled notifications for this user?
		Notifying: false,
		// Type: Array of String (ISO 639-1 language two-letter code)
		// Description: Which languages are you following from this user?
		Languages: []string{},
		// Type: Boolean
		// Description: Are you followed by this user?
		FollowedBy: lro2a.Type == model.RelationTypeFollow && lro2a.Status == model.RelationStatusAccepted,
		// Type: Boolean
		// Description: Are you blocking this user?
		Blocking: lra2o.Type == model.RelationTypeBlock,
		// Type: Boolean
		// Description: Is this user blocking you?
		BlockedBy: lro2a.Type == model.RelationTypeBlock,
		// Type: Boolean
		// Description: Are you muting this user?
		Muting: false,
		// Type: Boolean
		// Description: Are you muting notifications from this user?
		MutingNotifications: false,
		// Type: Boolean
		// Description: Do you have a pending follow request for this user?
		Requested: lra2o.Type == model.RelationTypeFollow && lra2o.Status == model.RelationStatusPadding,
		// Type: Boolean
		// Description: Has this user requested to follow you?
		RequestedBy: lro2a.Type == model.RelationTypeFollow && lro2a.Status == model.RelationStatusPadding,
		// Type: Boolean
		// Description: Are you blocking this user’s domain?
		DomainBlocking: false,
		// Type: Boolean
		// Description: Are you featuring this user on your profile?
		Endorsed: false,
		// Type: String
		// Description: This user’s profile bio
		Note: account.Note,
	}
	return relationship, nil
}
