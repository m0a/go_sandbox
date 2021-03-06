// Package models contains the types for schema 'feedbacks'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// Feedback represents a row from 'feedbacks.feedbacks'.
type Feedback struct {
	FeedbackID        int            `json:"feedback_id"`         // feedback_id
	FeedbacksStatusID int            `json:"feedbacks_status_id"` // feedbacks_status_id
	TantoName         sql.NullString `json:"tanto_name"`          // tanto_name
	URL               string         `json:"url"`                 // url
	HTML              string         `json:"html"`                // html
	Img               []byte         `json:"img"`                 // img
	Note              string         `json:"note"`                // note
	Created           *time.Time     `json:"created"`             // created
	Modified          *time.Time     `json:"modified"`            // modified

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Feedback exists in the database.
func (f *Feedback) Exists() bool {
	return f._exists
}

// Deleted provides information if the Feedback has been deleted from the database.
func (f *Feedback) Deleted() bool {
	return f._deleted
}

// Insert inserts the Feedback to the database.
func (f *Feedback) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if f._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO feedbacks.feedbacks (` +
		`feedbacks_status_id, tanto_name, url, html, img, note, created, modified` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, f.FeedbacksStatusID, f.TantoName, f.URL, f.HTML, f.Img, f.Note, f.Created, f.Modified)
	res, err := db.Exec(sqlstr, f.FeedbacksStatusID, f.TantoName, f.URL, f.HTML, f.Img, f.Note, f.Created, f.Modified)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	f.FeedbackID = int(id)
	f._exists = true

	return nil
}

// Update updates the Feedback in the database.
func (f *Feedback) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !f._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if f._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE feedbacks.feedbacks SET ` +
		`feedbacks_status_id = ?, tanto_name = ?, url = ?, html = ?, img = ?, note = ?, created = ?, modified = ?` +
		` WHERE feedback_id = ?`

	// run query
	XOLog(sqlstr, f.FeedbacksStatusID, f.TantoName, f.URL, f.HTML, f.Img, f.Note, f.Created, f.Modified, f.FeedbackID)
	_, err = db.Exec(sqlstr, f.FeedbacksStatusID, f.TantoName, f.URL, f.HTML, f.Img, f.Note, f.Created, f.Modified, f.FeedbackID)
	return err
}

// Save saves the Feedback to the database.
func (f *Feedback) Save(db XODB) error {
	if f.Exists() {
		return f.Update(db)
	}

	return f.Insert(db)
}

// Delete deletes the Feedback from the database.
func (f *Feedback) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !f._exists {
		return nil
	}

	// if deleted, bail
	if f._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM feedbacks.feedbacks WHERE feedback_id = ?`

	// run query
	XOLog(sqlstr, f.FeedbackID)
	_, err = db.Exec(sqlstr, f.FeedbackID)
	if err != nil {
		return err
	}

	// set deleted
	f._deleted = true

	return nil
}

// MStatus returns the MStatus associated with the Feedback's FeedbacksStatusID (feedbacks_status_id).
//
// Generated from foreign key 'feedbacks_status_id'.
func (f *Feedback) MStatus(db XODB) (*MStatus, error) {
	return MStatusByStatusID(db, f.FeedbacksStatusID)
}

// FeedbackByFeedbackID retrieves a row from 'feedbacks.feedbacks' as a Feedback.
//
// Generated from index 'feedbacks_feedback_id_pkey'.
func FeedbackByFeedbackID(db XODB, feedbackID int) (*Feedback, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`feedback_id, feedbacks_status_id, tanto_name, url, html, img, note, created, modified ` +
		`FROM feedbacks.feedbacks ` +
		`WHERE feedback_id = ?`

	// run query
	XOLog(sqlstr, feedbackID)
	f := Feedback{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, feedbackID).Scan(&f.FeedbackID, &f.FeedbacksStatusID, &f.TantoName, &f.URL, &f.HTML, &f.Img, &f.Note, &f.Created, &f.Modified)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

// FeedbacksByFeedbacksStatusID retrieves a row from 'feedbacks.feedbacks' as a Feedback.
//
// Generated from index 'feedbacks_status_id'.
func FeedbacksByFeedbacksStatusID(db XODB, feedbacksStatusID int) ([]*Feedback, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`feedback_id, feedbacks_status_id, tanto_name, url, html, img, note, created, modified ` +
		`FROM feedbacks.feedbacks ` +
		`WHERE feedbacks_status_id = ?`

	// run query
	XOLog(sqlstr, feedbacksStatusID)
	q, err := db.Query(sqlstr, feedbacksStatusID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Feedback{}
	for q.Next() {
		f := Feedback{
			_exists: true,
		}

		// scan
		err = q.Scan(&f.FeedbackID, &f.FeedbacksStatusID, &f.TantoName, &f.URL, &f.HTML, &f.Img, &f.Note, &f.Created, &f.Modified)
		if err != nil {
			return nil, err
		}

		res = append(res, &f)
	}

	return res, nil
}
