package controller

import (
	"go-programming-secure-your-go-apps/final_project/exception"
	"go-programming-secure-your-go-apps/final_project/helper"
	"go-programming-secure-your-go-apps/final_project/middleware"
	"go-programming-secure-your-go-apps/final_project/model/domain"
	"go-programming-secure-your-go-apps/final_project/model/response"
	"go-programming-secure-your-go-apps/final_project/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (cc *CommentControllerImpl) AddComment(writer http.ResponseWriter, request *http.Request) {
	user := middleware.ForContext(request.Context())
	id := strconv.Itoa(user.ID)

	var input domain.Comment
	helper.ReadFromRequestBody(request, &input)

	newComment, errCreate := cc.CommentService.AddComment(request.Context(), id, input)
	if errCreate != nil {
		panic(exception.NewBadRequestError(errCreate.Error()))
	}

	createComment := response.CreateCommentRespone{
		ID:        newComment.ID,
		Message:   newComment.Message,
		PhotoID:   newComment.PhotoID,
		UserID:    newComment.UserID,
		CreatedAt: newComment.CreatedAt,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   createComment,
	}

	helper.WriteToResponseBody(writer, http.StatusCreated, webResponse)
}

func (cc *CommentControllerImpl) GetAllComment(writer http.ResponseWriter, request *http.Request) {
	commentAll, errGetAll := cc.CommentService.GetAllComment(request.Context())
	if errGetAll != nil {
		panic(exception.NewBadRequestError(errGetAll.Error()))
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   commentAll,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}
func (cc *CommentControllerImpl) GetCommentById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	getById, errById := cc.CommentService.GetCommentById(request.Context(), id)
	if errById != nil {
		panic(exception.NewBadRequestError(errById.Error()))
	}

	getCommentById := response.GetCommentByIdRespone{
		ID:        getById.ID,
		Message:   getById.Message,
		PhotoID:   getById.PhotoID,
		UserID:    getById.UserID,
		CreatedAt: getById.CreatedAt,
		UpdatedAt: getById.UpdatedAt,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   getCommentById,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (cc *CommentControllerImpl) UpdateComment(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	var input domain.CommentInput
	helper.ReadFromRequestBody(request, &input)

	updatedComment, errUpdateComment := cc.CommentService.UpdateComment(request.Context(), id, input)
	if errUpdateComment != nil {
		panic(exception.NewBadRequestError(errUpdateComment.Error()))
	}

	newComment := response.UpdateCommentRespone{
		ID:        updatedComment.ID,
		Message:   updatedComment.Message,
		PhotoID:   updatedComment.PhotoID,
		UserID:    updatedComment.UserID,
		UpdatedAt: updatedComment.UpdatedAt,
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   newComment,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}

func (cc *CommentControllerImpl) DeleteComment(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]

	errDelete := cc.CommentService.DeleteComment(request.Context(), id)
	if errDelete != nil {
		panic(exception.NewBadRequestError(errDelete.Error()))
	}

	commentDelete := response.DeleteCommentRespone{
		Message: "Your comment has been successfully deleted",
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   commentDelete,
	}

	helper.WriteToResponseBody(writer, http.StatusOK, webResponse)
}
