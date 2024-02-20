package models

type Collaboration struct {
	MsmeID          string  `json:"mse_id"`
	StudentID       string  `json:"student_id"`
	InCollaboration bool    `json:"in_collaboration"`
	Progress        int8    `json:"progress"`
	Description     string  `json:"description"`
	Finished        bool    `json:"finished"`
	Feedback        string  `json:"feedback"`
	Rating          float32 `json:"rating"`
}

type CollabMsmeResp struct {
	Student         UserResponse `json:"student"`
	InCollaboration bool         `json:"in_collaboration"`
}

type CollabStudentResp struct {
	Msme            UserResponse `json:"msme"`
	InCollaboration bool         `json:"in_collaboration"`
}

type UpdateCollaboration struct {
	MsmeID      string `json:"mse_id"`
	StudentID   string `json:"student_id"`
	Progress    int8   `json:"progress"`
	Description string `json:"description"`
}

type GetStudentCollaboration struct {
	Msme            UserResponse `json:"msme"`
	InCollaboration bool         `json:"in_collaboration"`
	Progress        int8         `json:"progress"`
	Description     string       `json:"description"`
	Finished        bool         `json:"finished"`
	Feedback        string       `json:"feedback"`
	Rating          float32      `json:"rating"`
}
