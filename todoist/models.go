package todoist

import "time"

type Project struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	CommentCount int    `json:"comment_count"`
	IsShared     bool   `json:"is_shared"`
	IsFavorite   bool   `json:"is_favorite"`
	IsTeamInbox  bool   `json:"is_team_inbox,omitempty"`
	ViewStyle    string `json:"view_style,omitempty"`
	Url          string `json:"url"`

	ParentId     *string `json:"parent_id,omitempty"`
	Order        *int    `json:"order,omitempty"`
	InboxProject *bool   `json:"inbox_project,omitempty"`
	TeamInbox    *bool   `json:"team_inbox,omitempty"`
	SyncId       *int    `json:"sync_id,omitempty"`
}

type Due struct {
	Date      string    `json:"date,omitempty"`
	Datetime  time.Time `json:"datetime,omitempty"`
	Recurring bool      `json:"recurring,omitempty"`
	String    string    `json:"string,omitempty"`
	Timezone  string    `json:"timezone,omitempty"`
}

type Task struct {
	Assignee     int     `json:"assignee,omitempty"`
	Assigner     int     `json:"assigner,omitempty"`
	CommentCount int     `json:"comment_count,omitempty"`
	Completed    bool    `json:"completed,omitempty"`
	Content      string  `json:"content,omitempty"`
	Description  string  `json:"description,omitempty"`
	Due          *Due    `json:"due,omitempty"`
	Id           int64   `json:"id,omitempty"`
	LabelIds     []int64 `json:"label_ids,omitempty"`
	Order        int     `json:"order,omitempty"`
	Priority     int     `json:"priority,omitempty"`
	ProjectId    string  `json:"project_id,omitempty"`
	SectionId    int     `json:"section_id,omitempty"`
	ParentId     int64   `json:"parent_id,omitempty"`
	Url          string  `json:"url,omitempty"`
}
