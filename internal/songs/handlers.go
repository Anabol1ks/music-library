package songs

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Anabol1ks/music-library/internal/models"
	"github.com/Anabol1ks/music-library/internal/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InputSong struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

// CreateSong добавляет новую песню в базу данных
// @Summary Добавление новой песни
// @Description Добавляет новую песню в базу данных
// @Tags songs
// @Accept json
// @Produce json
// @Param input body InputSong true "Добавление песни"
// @Success 201 {object} response.CreateSuccessResponse "Песня успешно добавлена"
// @Failure 400 {object} response.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} response.ErrorResponse "Ошибка добавления песни"
// @Router /songs [post]
func CreateSong(c *gin.Context) {
	log.Println("[DEBUG] Начало CreateSong")
	var input InputSong
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("[DEBUG] Ошибка при парсинге JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// Проверка на дублирование
	result := storage.DB.Where("song_group = ? AND title = ?", input.Group, input.Song).First(&models.Song{})
	if result.RowsAffected > 0 {
		log.Printf("[INFO] Песня с группой '%s' и названием '%s' уже существует", input.Group, input.Song)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Такая песня уже есть в базе"})
		return
	}

	song := models.Song{
		Group: input.Group,
		Title: input.Song,
	}

	if err := storage.DB.Create(&song).Error; err != nil {
		log.Printf("[DEBUG] Ошибка создания песни: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления песни"})
		return
	}

	log.Printf("[INFO] Песня успешно создана, ID: %d", song.ID)
	c.JSON(http.StatusCreated, gin.H{"message": "Песня успешно добавлена"})
}

// UpdateSongInput описывает поля для частичного обновления песни
type UpdateSongInput struct {
	Group       *string `json:"group"`
	Title       *string `json:"title"`
	ReleaseDate *string `json:"release_date"`
	Text        *string `json:"text"`
	Link        *string `json:"link"`
}

// UpdateSong позволяет обновлять данные песни частично.
// @Summary Обновление данных песни
// @Description Частичное обновление данных песни по её ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор песни"
// @Param input body UpdateSongInput true "Данные для обновления"
// @Success 200 {object} response.UpdateResponse "Успешное обновление песни"
// @Failure 400 {object} response.ErrorResponse "Ошибка в запросе"
// @Failure 404 {object} response.ErrorResponse "Песня не найдена"
// @Failure 500 {object} response.ErrorResponse "Ошибка обновления данных песни"
// @Router /songs/{id} [patch]
func UpdateSong(c *gin.Context) {
	log.Println("[DEBUG] Начало UpdateSong")
	// Получаем идентификатор песни из URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("[DEBUG] Неверный идентификатор: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор песни"})
		return
	}

	// Ищем песню по ID
	var song models.Song
	if err := storage.DB.First(&song, id).Error; err != nil {
		log.Printf("[INFO] Песня с ID %d не найдена", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Песня не найдена"})
		return
	}

	// Получаем данные для обновления
	var input UpdateSongInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("[DEBUG] Ошибка при парсинге JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	updatedFields := make(map[string]interface{})
	if input.Group != nil {
		updatedFields["song_group"] = *input.Group
	}
	if input.Title != nil {
		updatedFields["title"] = *input.Title
	}
	if input.ReleaseDate != nil {
		updatedFields["release_date"] = *input.ReleaseDate
	}
	if input.Text != nil {
		updatedFields["text"] = *input.Text
	}
	if input.Link != nil {
		updatedFields["link"] = *input.Link
	}

	if len(updatedFields) == 0 {
		log.Println("[DEBUG] Нет данных для обновления")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Нет данных для обновления"})
		return
	}

	// Выполняем обновление
	if err := storage.DB.Model(&song).Updates(updatedFields).Error; err != nil {
		log.Printf("[DEBUG] Ошибка обновления данных песни: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления данных песни"})
		return
	}

	log.Printf("[INFO] Песня с ID %d успешно обновлена", song.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Песня успешно обновлена",
		"song":    song,
	})
}

// Paginate возвращает скоуп для пагинации
func Paginate(page, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * perPage
		log.Printf("[DEBUG] Пагинация: страница %d, записей на странице %d, offset %d", page, perPage, offset)
		return db.Offset(offset).Limit(perPage)
	}
}

// GetSongsWithFilters возвращает список песен с фильтрацией и пагинацией
// @Summary Получение данных библиотеки с фильтрацией и пагинацией
// @Description Получает список песен по фильтрам: group, title, release_date, link; поддерживается пагинация
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string false "Группа"
// @Param title query string false "Название песни"
// @Param release_date query string false "Дата релиза"
// @Param link query string false "Ссылка"
// @Param page query int false "Номер страницы (начиная с 1)"
// @Param per_page query int false "Количество записей на страницу"
// @Success 200 {array} models.Song
// @Failure 400 {object} response.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} response.ErrorResponse "Ошибка получения данных"
// @Router /songs [get]
func GetSongsWithFilters(c *gin.Context) {
	log.Println("[DEBUG] Начало GetSongsWithFilters")
	var songs []models.Song

	// Получаем параметры фильтрации
	groupFilter := c.Query("group")
	titleFilter := c.Query("title")
	releaseDateFilter := c.Query("release_date")
	linkFilter := c.Query("link")

	// Параметры пагинации
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		log.Printf("[DEBUG] Неверный параметр page: %s", pageStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный параметр page"})
		return
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		log.Printf("[DEBUG] Неверный параметр per_page: %s", perPageStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный параметр per_page"})
		return
	}

	// Формируем запрос с фильтрами
	query := storage.DB.Model(&models.Song{})
	if groupFilter != "" {
		query = query.Where("song_group = ?", groupFilter)
	}
	if titleFilter != "" {
		query = query.Where("title ILIKE ?", "%"+titleFilter+"%")
	}
	if releaseDateFilter != "" {
		query = query.Where("release_date = ?", releaseDateFilter)
	}
	if linkFilter != "" {
		query = query.Where("link = ?", linkFilter)
	}

	// Применяем пагинацию и получаем записи
	if err := query.Scopes(Paginate(page, perPage)).Find(&songs).Error; err != nil {
		log.Printf("[DEBUG] Ошибка получения песен: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}

	log.Printf("[INFO] Получено %d песен", len(songs))
	c.JSON(http.StatusOK, songs)
}

// GetSongTextWithPagination возвращает текст песни, разбитый на куплеты, с пагинацией
// @Summary Получение текста песни с пагинацией по куплетам
// @Description Возвращает куплеты текста песни по указанной странице и количеству куплетов на странице
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор песни"
// @Param page query int false "Номер страницы (начиная с 1)" default(1)
// @Param per_page query int false "Количество куплетов на странице" default(5)
// @Success 200 {object} response.TextResponse "Текст песни с пагинацией по куплетам"
// @Failure 400 {object} response.ErrorResponse "Ошибка в запросе"
// @Failure 404 {object} response.ErrorResponse "Песня не найдена"
// @Router /songs/{id}/text [get]
func GetSongTextWithPagination(c *gin.Context) {
	log.Println("[DEBUG] Начало GetSongTextWithPagination")
	// Получаем ID песни из пути
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[DEBUG] Неверный идентификатор: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор"})
		return
	}

	// Находим песню
	var song models.Song
	if err := storage.DB.First(&song, id).Error; err != nil {
		log.Printf("[INFO] Песня с ID %d не найдена", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Песня не найдена"})
		return
	}

	// Получаем параметры пагинации для куплетов
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "5")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		log.Printf("[DEBUG] Неверный параметр page: %s", pageStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный параметр page"})
		return
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		log.Printf("[DEBUG] Неверный параметр per_page: %s", perPageStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный параметр per_page"})
		return
	}

	// Предполагаем, что куплеты разделены двумя переводами строки
	verses := strings.Split(song.Text, "\n\n")
	totalVerses := len(verses)
	start := (page - 1) * perPage
	if start >= totalVerses {
		log.Printf("[DEBUG] Запрошенная страница %d вне диапазона. Всего куплетов: %d", page, totalVerses)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Страница вне диапазона"})
		return
	}
	end := start + perPage
	if end > totalVerses {
		end = totalVerses
	}

	pagedVerses := verses[start:end]
	log.Printf("[INFO] Песня ID %d: возвращено куплетов %d (страница %d из %d)", song.ID, len(pagedVerses), page, (totalVerses+perPage-1)/perPage)
	c.JSON(http.StatusOK, gin.H{
		"song_id":      song.ID,
		"total_verses": totalVerses,
		"page":         page,
		"per_page":     perPage,
		"verses":       pagedVerses,
	})
}

// DeleteSong удаляет песню по её ID
// @Summary Удаление песни
// @Description Удаляет песню из базы данных по указанному идентификатору
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор песни"
// @Success 200 {object} response.DeleteSuccessResponse "Песня успешно удалена"
// @Failure 400 {object} response.ErrorResponse "Ошибка в запросе"
// @Failure 404 {object} response.ErrorResponse "Песня не найдена
// @Failure 500 {object} response.ErrorResponse "Ошибка удаления песни"
// @Router /songs/{id} [delete]
func DeleteSong(c *gin.Context) {
	log.Println("[DEBUG] Начало DeleteSong")
	// Получаем идентификатор песни из пути
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[DEBUG] Неверный идентификатор: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный идентификатор"})
		return
	}

	// Ищем песню
	var song models.Song
	if err := storage.DB.First(&song, id).Error; err != nil {
		log.Printf("[INFO] Песня с ID %d не найдена", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Песня не найдена"})
		return
	}

	// Удаляем песню
	if err := storage.DB.Delete(&song).Error; err != nil {
		log.Printf("[DEBUG] Ошибка удаления песни с ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления песни"})
		return
	}

	log.Printf("[INFO] Песня с ID %d успешно удалена", id)
	c.JSON(http.StatusOK, gin.H{"message": "Песня успешно удалена"})
}
