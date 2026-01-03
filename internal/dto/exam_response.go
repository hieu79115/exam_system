package dto

type ExamResponse struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Duration      int    `json:"duration"`
	QuestionCount int    `json:"questionCount"`
}

type SelectionResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type QuestionResponse struct {
	ID               string              `json:"id"`
	Type             string              `json:"type"`
	Description      string              `json:"description"`
	MaxText          int                 `json:"maxText,omitempty"`
	ReadingPassageID *string             `json:"readingPassageId,omitempty"`
	Selections       []SelectionResponse `json:"selections,omitempty"`
}

type ExamDetailResponse struct {
	Exam      ExamResponse       `json:"exam"`
	Questions []QuestionResponse `json:"questions"`
	Passages  []PassageResponse  `json:"passages"`
}

type ExamListItemResponse struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Duration      int    `json:"duration"`
	QuestionCount int    `json:"questionCount"`
}

type CreateExamRequest struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Duration    int    `json:"duration" binding:"required,min=1"`
}

type UpdateExamRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Duration    int    `json:"duration" binding:"min=1"`
}
