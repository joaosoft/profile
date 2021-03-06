package profile

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/joaosoft/dbr"
	"github.com/joaosoft/web"
)

type ErrorResponse struct {
	Code    web.Status `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
	Cause   string     `json:"cause,omitempty"`
}

type GetSectionRequest struct {
	SectionKey string `json:"section_key" validate:"notzero"`
}

type GetSectionContentsRequest struct {
	SectionKey string `json:"section_key" validate:"notzero"`
}

type SectionList []*Section

type Section struct {
	IdSection   string `json:"id_section" db:"id_section"`
	Key         string `json:"key" db:"key"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type ContentList []*Content

type Content struct {
	IdContent string           `json:"id_content" db:"id_content"`
	Key       string           `json:"key" db:"key"`
	Type      string           `json:"type" db:"type"`
	Content   *json.RawMessage `json:"content" db:"content"`
}

type SectionsContentsList []*SectionContents

type SectionContents struct {
	Section
	Contents ContentList `json:"contents" db:"contents"`
}

func (l *ContentList) Value() (driver.Value, error) {
	j, err := dbr.Marshal(l)
	return j, err
}
func (l *ContentList) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return ErrorInvalidType
	}

	err := dbr.Unmarshal(source, l)
	if err != nil {
		return err
	}

	return nil
}
