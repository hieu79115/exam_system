package dto

// --- REQUEST ---

type CreateSelectionReq struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type CreateQuestionRequest struct {
	ExaminationID    string  `json:"examinationId" binding:"required"`
	ReadingPassageID *string `json:"readingPassageId"`

	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	MaxText     int    `json:"maxText"`

	Selections []CreateSelectionReq `json:"selections" binding:"dive"`
}

type UpdateQuestionRequest struct {
	ExaminationID    string               `json:"examinationId" binding:"required"`
	ReadingPassageID *string              `json:"readingPassageId"`
	Type             string               `json:"type" binding:"required"`
	Description      string               `json:"description" binding:"required"`
	MaxText          int                  `json:"maxText"`
	Selections       []CreateSelectionReq `json:"selections" binding:"dive"`
}

// --- RESPONSE ---

type SelectionRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type QuestionRes struct {
	ID               string         `json:"id"`
	ExaminationID    string         `json:"examinationId"`
	ReadingPassageID *string        `json:"readingPassageId,omitempty"`
	Type             string         `json:"type"`
	Description      string         `json:"description"`
	Selections       []SelectionRes `json:"selections,omitempty"`
}
