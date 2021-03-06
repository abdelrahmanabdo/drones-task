package medication

import (
	"github.com/gin-gonic/gin"
)

func createMedication(c *gin.Context) error {
	u := &medication{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	medications[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getMedication(c *gin.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, medications[id])
}

func getMedications(c *gin.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, medications[id])
}

func updateMedication(c *gin.Context) error {
	u := new(drone)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	medications[id].Name = u.Name
	return c.JSON(http.StatusOK, medications[id])
}

func deleteMedication(c *gin.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(medications, id)
	return c.NoContent(http.StatusNoContent)
}