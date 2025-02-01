package response

import "github.com/Anabol1ks/music-library/internal/models"

type CreateSuccessResponse struct {
	Message string `json:"message" example:"Песня успешно добавлена"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"текст ошибки"`
}

type DeleteSuccessResponse struct {
	Message string `json:"message" example:"Песня успешно удалена"`
}

type UpdateResponse struct {
	Message string      `json:"message" example:"Данные песни успешно обновлены"`
	Song    models.Song `json:"song"`
}

type TextResponse struct {
	Page        int      `json:"page" example:"1"`
	Per_page    int      `json:"per_page" example:"5"`
	SongID      int      `json:"song_id" example:"1"`
	TotalVerses int      `json:"total_verses" example:"2"`
	Verser      []string `json:"verses" example:"[Куплет 1, Куплет 2]"`
}
