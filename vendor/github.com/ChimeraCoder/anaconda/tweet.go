package anaconda

import (
	"fmt"
	"time"
)

type Tweet struct {
	Contributors         []Contributor          `json:"contributors"` // Not yet generally available to all, so hard to test
	Coordinates          *Coordinates           `json:"coordinates"`
	CreatedAt            string                 `json:"created_at"`
	Entities             Entities               `json:"entities"`
	ExtendedEntities     Entities               `json:"extended_entities"`
	FavoriteCount        int                    `json:"favorite_count"`
	Favorited            bool                   `json:"favorited"`
	FilterLevel          string                 `json:"filter_level"`
	Id                   int64                  `json:"id"`
	IdStr                string                 `json:"id_str"`
	InReplyToScreenName  string                 `json:"in_reply_to_screen_name"`
	InReplyToStatusID    int64                  `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr string                 `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64                  `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   string                 `json:"in_reply_to_user_id_str"`
	Lang                 string                 `json:"lang"`
	Place                Place                  `json:"place"`
	PossiblySensitive    bool                   `json:"possibly_sensitive"`
	RetweetCount         int                    `json:"retweet_count"`
	Retweeted            bool                   `json:"retweeted"`
	RetweetedStatus      *Tweet                 `json:"retweeted_status"`
	Source               string                 `json:"source"`
	Scopes               map[string]interface{} `json:"scopes"`
	Text                 string                 `json:"text"`
	Truncated            bool                   `json:"truncated"`
	User                 User                   `json:"user"`
	WithheldCopyright    bool                   `json:"withheld_copyright"`
	WithheldInCountries  []string               `json:"withheld_in_countries"`
	WithheldScope        string                 `json:"withheld_scope"`

	//Geo is deprecated
	//Geo                  interface{} `json:"geo"`
}

// CreatedAtTime is a convenience wrapper that returns the Created_at time, parsed as a time.Time struct
func (t Tweet) CreatedAtTime() (time.Time, error) {
	return time.Parse(time.RubyDate, t.CreatedAt)
}

// It may be worth placing these in an additional source file(s)

// Could also use User, since the fields match, but only these fields are possible in Contributor
type Contributor struct {
	Id         int64  `json:"id"`
	IdStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"` // Coordinate always has to have exactly 2 values
	Type        string     `json:"type"`
}

// HasCoordinates is a helper function to easily determine if a Tweet has coordinates associated with it
func (t Tweet) HasCoordinates() bool {
	if t.Coordinates != nil {
		if t.Coordinates.Type == "Point" {
			return true
		}
	}
	return false
}

// The following provide convenience and eliviate confusion about the order of coordinates in the Tweet

// Latitude is a convenience wrapper that returns the latitude easily
func (t Tweet) Latitude() (float64, error) {
	if t.HasCoordinates() {
		return t.Coordinates.Coordinates[1], nil
	}
	return 0, fmt.Errorf("No Coordinates in this Tweet")
}

// Longitude is a convenience wrapper that returns the longitude easily
func (t Tweet) Longitude() (float64, error) {
	if t.HasCoordinates() {
		return t.Coordinates.Coordinates[0], nil
	}
	return 0, fmt.Errorf("No Coordinates in this Tweet")
}

// X is a concenience wrapper which returns the X (Longitude) coordinate easily
func (t Tweet) X() (float64, error) {
	return t.Longitude()
}

// Y is a convenience wrapper which return the Y (Lattitude) corrdinate easily
func (t Tweet) Y() (float64, error) {
	return t.Latitude()
}
