package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SadamIRM/be-smoke-store/config"
	"github.com/SadamIRM/be-smoke-store/models"
	"github.com/SadamIRM/be-smoke-store/services"
)

type TransactionHandler struct {
	txService *services.TransactionService
}

func NewTransactionHandler() *TransactionHandler {
	return &TransactionHandler{txService: services.NewTransactionService()}
}

func (h *TransactionHandler) Checkout(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak ditemukan"})
		return
	}

	var req models.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Metode pembayaran wajib diisi"})
		return
	}

	tx, err := h.txService.Checkout(userID, req.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Checkout berhasil",
		"data":    tx,
	})
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak ditemukan"})
		return
	}

	transactions, err := h.txService.GetTransactions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal mengambil riwayat transaksi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    transactions,
	})
}

func (h *TransactionHandler) Callback(c *gin.Context) {
	status := c.Query("status")
	reference := c.Query("reference")

	if reference == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Reference is required"})
		return
	}

	if status == "success" {
		var tx models.Transaction
		if err := config.DB.Where("transaction_number = ?", reference).First(&tx).Error; err == nil {
			tx.Status = "Selesai"
			config.DB.Save(&tx)
		}
	} else if status == "failed" || status == "cancelled" {
		var tx models.Transaction
		if err := config.DB.Where("transaction_number = ?", reference).First(&tx).Error; err == nil {
			tx.Status = "Gagal"
			config.DB.Save(&tx)
		}
	}

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Pembayaran Sukses</title>
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<style>
				body {
					font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
					background-color: #FAF6F1;
					color: #3E2723;
					text-align: center;
					padding: 50px 20px;
				}
				.card {
					background-color: #FFFFFF;
					border-radius: 20px;
					padding: 40px 30px;
					max-width: 400px;
					margin: 0 auto;
					box-shadow: 0 10px 25px rgba(111, 78, 55, 0.1);
					border: 1px solid rgba(139, 111, 71, 0.15);
				}
				.icon {
					font-size: 60px;
					color: #4CAF50;
					margin-bottom: 20px;
				}
				h1 {
					font-size: 24px;
					margin-bottom: 10px;
				}
				p {
					font-size: 14px;
					color: #795548;
					margin-bottom: 30px;
				}
				.btn {
					background-color: #6F4E37;
					color: white;
					text-decoration: none;
					padding: 12px 30px;
					border-radius: 10px;
					font-weight: bold;
					display: inline-block;
				}
			</style>
		</head>
		<body>
			<div class="card">
				<div class="icon">✓</div>
				<h1>Pembayaran Diterima!</h1>
				<p>Pembayaran untuk transaksi <b>`+reference+`</b> telah berhasil dikonfirmasi. Anda dapat kembali ke aplikasi Toko Material sekarang.</p>
				<a href="javascript:window.close();" class="btn">Kembali ke Aplikasi</a>
			</div>
		</body>
		</html>
	`)
}
