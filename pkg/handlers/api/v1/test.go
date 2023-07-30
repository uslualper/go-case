package handlers

import (
	"github.com/gofiber/fiber/v2"

	payloadSchema "yuka-case/pkg/schema/payload/v1"
	responseSchema "yuka-case/pkg/schema/response/v1"
	"yuka-case/pkg/utils/validate"
)

type Test struct{}

// Test
// @Summary Test
// @Description get test data
// @Tags Test
// @Accept json
// @Produce json
// @Param payload body payloadSchema.Test true "Test Payload"
// @Response 200 {object} responseSchema.Test
// @Router /v1/test [post]
func (t *Test) Test(c *fiber.Ctx) error {
	testValid := new(payloadSchema.Test)
	c.BodyParser(testValid)
	if err := validate.ValidateStruct(testValid); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(err)
	}

	testResponse := responseSchema.Test{
		Message: testValid.Message,
	}

	return c.JSON(testResponse)
}
