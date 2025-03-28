package handlers

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/internal/api/presenters"
	"Go-Starter-Template/pkg/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	UserHandler interface {
		RegisterUser(c *fiber.Ctx) error
		Login(c *fiber.Ctx) error
		GetProfile(c *fiber.Ctx) error
		UpdateProfile(c *fiber.Ctx) error
		PostEducation(c *fiber.Ctx) error
		UpdateEducation(c *fiber.Ctx) error
		DeleteEducation(c *fiber.Ctx) error
		PostExperience(c *fiber.Ctx) error
		UpdateExperience(c *fiber.Ctx) error
		DeleteExperience(c *fiber.Ctx) error
		PostSkill(c *fiber.Ctx) error
		DeleteSkill(c *fiber.Ctx) error
		SearchUser(c *fiber.Ctx) error
		GetSkills(c *fiber.Ctx) error
	}
	userHandler struct {
		UserService user.UserService
		Validator   *validator.Validate
	}
)

func NewUserHandler(userService user.UserService, validator *validator.Validate) UserHandler {
	return &userHandler{
		UserService: userService,
		Validator:   validator,
	}
}

func (h *userHandler) RegisterUser(c *fiber.Ctx) error {
	req := new(domain.UserRegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	req.ProfilePicture, _ = c.FormFile("profile_picture")
	req.Headline, _ = c.FormFile("headline")

	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	res, err := h.UserService.RegisterUser(c.Context(), *req)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedRegister, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusCreated, domain.MessageSuccessRegister)
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	req := new(domain.UserLoginRequest)
	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	res, err := h.UserService.Login(c.Context(), *req)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedLogin, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusOK, domain.MessageSuccessLogin)
}

func (h *userHandler) GetProfile(c *fiber.Ctx) error {
	slug := c.Params("slug")
	res, err := h.UserService.GetProfile(c.Context(), slug)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedGetProfile, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusOK, domain.MessageSuccessGetProfile)
}

func (h *userHandler) UpdateProfile(c *fiber.Ctx) error {
	req := new(domain.UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	req.ProfilePicture, _ = c.FormFile("profile_picture")
	req.Headline, _ = c.FormFile("headline")

	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	userid := c.Locals("user_id").(string)

	if err := h.UserService.UpdateProfile(c.Context(), *req, userid); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedUpdateUser, err)
	}

	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessUpdateUser)
}

func (h *userHandler) PostEducation(c *fiber.Ctx) error {
	req := new(domain.PostUserEducationRequest)
	userid := c.Locals("user_id").(string)
	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	if err := h.UserService.PostEducation(c.Context(), *req, userid); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedAddEducation, err)
	}
	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessAddEducation)
}

func (h *userHandler) UpdateEducation(c *fiber.Ctx) error {
	req := new(domain.UpdateUserEducationRequest)
	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}
	userid := c.Locals("user_id").(string)
	if err := h.UserService.UpdateEducation(c.Context(), *req, userid); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedUpdateEducation, err)
	}
	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessUpdateEducation)
}

func (h *userHandler) DeleteEducation(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.UserService.DeleteEducation(c.Context(), id); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedDeleteEducation, err)
	}
	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessDeleteEducation)
}

func (h *userHandler) PostExperience(c *fiber.Ctx) error {
	req := new(domain.PostUserExperienceRequest)

	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	userid := c.Locals("user_id").(string)

	if err := h.UserService.PostExperience(c.Context(), *req, userid); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedAddEducation, err)
	}

	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessAddEducation)
}

func (h *userHandler) UpdateExperience(c *fiber.Ctx) error {
	req := new(domain.UpdateUserExperienceRequest)

	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	userid := c.Locals("user_id").(string)

	if err := h.UserService.UpdateExperience(c.Context(), *req, userid); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedAddEducation, err)
	}

	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessAddEducation)
}

func (h *userHandler) DeleteExperience(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.UserService.DeleteExperience(c.Context(), id); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedDeleteExperience, err)
	}

	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessDeleteExperience)
}

func (h *userHandler) PostSkill(c *fiber.Ctx) error {
	req := new(domain.PostUserSkillRequest)

	if err := c.BodyParser(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	if err := h.Validator.Struct(req); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedBodyRequest, err)
	}

	userid := c.Locals("user_id").(string)

	if err := h.UserService.PostSkill(c.Context(), *req, userid); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedAddSkill, err)
	}

	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessAddSkill)
}

func (h *userHandler) DeleteSkill(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.UserService.DeleteSkill(c.Context(), id); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedDeleteEducation, err)
	}

	return presenters.SuccessResponse(c, nil, fiber.StatusOK, domain.MessageSuccessDeleteEducation)
}

func (h *userHandler) SearchUser(c *fiber.Ctx) error {
	query := domain.UserSearchRequest{
		Keyword: c.Query("keyword"),
	}

	res, err := h.UserService.SearchUser(c.Context(), query)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedSearchUser, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusOK, domain.MessageSuccessSearchUser)
}

func (h *userHandler) GetSkills(c *fiber.Ctx) error {
	res, err := h.UserService.GetSkills(c.Context())
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadRequest, domain.MessageFailedGetSkills, err)
	}
	return presenters.SuccessResponse(c, res, fiber.StatusOK, domain.MessageSuccessGetSkills)
}
