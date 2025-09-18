package handlers

import (
	"temp-backend-at-kbtg/database"
	"temp-backend-at-kbtg/models"

	"github.com/gofiber/fiber/v2"
)

// GetProfile godoc
// @Summary Get user profile
// @Description Get current user's profile information
// @Tags Profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.ProfileResponse
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /profile [get]
func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(models.ProfileResponse{
		User: user,
	})
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update current user's profile information
// @Tags Profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param profile body models.UpdateProfileRequest true "Profile update data"
// @Success 200 {object} models.ProfileResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /profile [put]
func UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req models.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Update user fields
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	// Save updated user
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update profile",
		})
	}

	return c.JSON(models.ProfileResponse{
		User: user,
	})
}

// GetMembershipInfo godoc
// @Summary Get membership information
// @Description Get current user's membership details including points and level
// @Tags Profile
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /profile/membership [get]
func GetMembershipInfo(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"membership_id": user.MembershipID,
		"member_level":  user.MemberLevel,
		"points":        user.Points,
		"member_since":  user.CreatedAt.Format("2/1/2006"),
		"full_name":     user.FirstName + " " + user.LastName,
		"email":         user.Email,
		"phone":         user.Phone,
	})
}
