package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/AryaWgna/go-restapi-gin/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var mahasiswa []models.Mahasiswa

	models.DB.Find(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"mahasiswa": mahasiswa})

}

func Show(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := models.DB.First(&mahasiswa, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data mahasiswa tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"mahasiswa": mahasiswa})
}

func Create(c *gin.Context) {

	var mahasiswa models.Mahasiswa

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"mahasiswa": mahasiswa})
}

func Update(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&mahasiswa).Where("id = ?", id).Updates(&mahasiswa).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data mahasiswa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data mahasiswa berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var mahasiswa models.Mahasiswa

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&mahasiswa, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data mahasiswa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data mahasiswa berhasil dihapus"})
}
