package dto

// 1. Mapping: interface Examination
type ExamResponse struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Duration      int    `json:"duration"`
	QuestionCount int    `json:"questionCount"`
}

// 2. Mapping: interface SelectionInfo
type SelectionResponse struct {
	ID   string `json:"id"` // Go dùng string cho ID uuid
	Name string `json:"name"`
	Code string `json:"code"`
}

// 3. Mapping: interface ReadingPassage
type PassageResponse struct {
	ID          string `json:"id"`
	Description string `json:"description,omitempty"` // omitempty: nếu rỗng thì không trả về field này
	Text        string `json:"text"`
}

// 4. Mapping: interface Question
type QuestionResponse struct {
	ID               string              `json:"id"`
	Type             string              `json:"type"` // QuestionType bên TS là string hoặc enum
	Description      string              `json:"description"`
	MaxText          int                 `json:"maxText,omitempty"`
	ReadingPassageID *string             `json:"readingPassageId,omitempty"` // Dùng con trỏ để có thể null
	Selections       []SelectionResponse `json:"selections,omitempty"`       // Nhúng mảng Selection vào đây
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
