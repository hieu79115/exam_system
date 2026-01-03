package dto

type PassageResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Text        string `json:"text"`
}

type CreatePassageRequest struct {
	ID          string `json:"id"`
	Text        string `json:"text" binding:"required"`
	Description string `json:"description"`
}

type UpdatePassageRequest struct {
	Text        string `json:"text" binding:"required"`
	Description string `json:"description"`
}
