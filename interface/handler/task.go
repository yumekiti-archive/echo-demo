package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/echo-demo/usecase"
)

// TaskHandler task handlerのinterface
type TaskHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

// NewTaskHandler task handlerのコンストラクタ
func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase: taskUsecase}
}

type requestTask struct {
	Title   string `json:"title"`
	Context string `json:"context"`
	Status  bool   `json:"status"`
}

type responseTask struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Context string `json:"context"`
	Status  bool   `json:"status"`
}

// Post taskを保存するときのハンドラー
func (th *taskHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdTask, err := th.taskUsecase.Create(req.Title, req.Context, req.Status)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      createdTask.ID,
			Title:   createdTask.Title,
			Context: createdTask.Context,
			Status:  createdTask.Status,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get taskを取得するときのハンドラー
func (th *taskHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundTask, err := th.taskUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      foundTask.ID,
			Title:   foundTask.Title,
			Context: foundTask.Context,
			Status:  foundTask.Status,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put taskを更新するときのハンドラー
func (th *taskHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedTask, err := th.taskUsecase.Update(id, req.Title, req.Context, req.Status)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      updatedTask.ID,
			Title:   updatedTask.Title,
			Context: updatedTask.Context,
			Status:  updatedTask.Status,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete taskを削除するときのハンドラー
func (th *taskHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.taskUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
