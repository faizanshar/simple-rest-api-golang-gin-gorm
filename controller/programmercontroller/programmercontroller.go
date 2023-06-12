package programmercontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"programmerData/models"
)

func Index(c *gin.Context) {
	var products []models.Programmer

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}

func Show(c *gin.Context) {
	var programmer []models.Programmer

	id := c.Param("id")

	if err := models.DB.First(&programmer, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   programmer,
	})
}

func Create(c *gin.Context) {
	var programmer models.Programmer

	if err := c.ShouldBindJSON(&programmer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	models.DB.Create(&programmer)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   programmer,
	})
}

func Update(c *gin.Context) {
	var programmer models.Programmer

	id := c.Param("id")

	if err := c.ShouldBindJSON(&programmer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if models.DB.Model(&programmer).Where("id = ?", id).Updates(&programmer).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "tidak dapat mengubah data!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data berhasil diperbarui",
	})

}

func Delete(c *gin.Context) {
	var programmer models.Programmer

	id := c.Param("id")

	//if err := c.ShouldBindJSON(&programmer); err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}

	if models.DB.Model(&programmer).Where("id = ?", id).Delete(&programmer).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "tidak dapat menghapus data!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data berhasil dihapus",
	})
}
