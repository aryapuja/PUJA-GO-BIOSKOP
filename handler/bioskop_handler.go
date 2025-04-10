package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	database "puja_go_bioskop/db"
	"puja_go_bioskop/model"

	"github.com/gin-gonic/gin"
)

// POST
func CreateBioskop(ctx *gin.Context) {
	var bioskop model.Bioskop

	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":  "Semua field wajib diisi dengan benar",
			"detail": err.Error(),
		})
		ctx.Abort()
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	err := database.DB.QueryRow(
		"INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id",
		bioskop.Nama, bioskop.Lokasi, *bioskop.Rating,
	).Scan(&bioskop.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Bioskop dengan ID %v berhasil ditambahkan", bioskop.ID),
		"id":      bioskop.ID,
	})
}

//POST

// GET
func GetBioskop(ctx *gin.Context) {
	rows, err := database.DB.Query("SELECT id, nama, lokasi, rating FROM bioskop")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var bioskops []model.Bioskop
	for rows.Next() {
		var b model.Bioskop
		if err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		bioskops = append(bioskops, b)
	}

	ctx.JSON(http.StatusOK, bioskops)
}

//GET

// GET
func GetBioskopInfo(ctx *gin.Context) {
	var bioskop model.Bioskop
	id := ctx.Param("id")

	// Cek ID bioskop
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM bioskop WHERE id = $1)", id).Scan(&exists)
	if err != nil || !exists {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Bioskop dengan ID %v tidak ditemukan", id),
		})
		return
	}

	err = database.DB.QueryRow("SELECT id, nama, lokasi, rating FROM bioskop WHERE id=$1", id).Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bioskop)
}

//GET

// PUT
func UpdateBioskop(ctx *gin.Context) {
	var bioskop model.Bioskop

	id := ctx.Param("id")

	// Cek ID bioskop
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM bioskop WHERE id = $1)", id).Scan(&exists)
	if err != nil || !exists {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Bioskop dengan ID %v tidak ditemukan", id),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&bioskop); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi wajib diisi"})
		return
	}

	var result sql.Result

	if bioskop.Rating == nil {
		result, err = database.DB.Exec(
			"UPDATE bioskop SET nama=$1, lokasi=$2 WHERE id=$3",
			bioskop.Nama, bioskop.Lokasi, id,
		)

	} else {
		result, err = database.DB.Exec(
			"UPDATE bioskop SET nama=$1, lokasi=$2, rating=$3 WHERE id=$4",
			bioskop.Nama, bioskop.Lokasi, *bioskop.Rating, id,
		)

	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Bioskop dengan ID %v tidak ditemukan", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Bioskop dengan ID %v berhasil di update", id),
		"id":      id,
	})
}

//PUT

// DEL
func DeleteBioskop(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := database.DB.Exec("DELETE FROM bioskop WHERE id = $1", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Bioskop dengan ID %v tidak ditemukan", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Bioskop dengan ID %v berhasil dihapus", id),
		"id":      id,
	})
}

//DEL
