package handlers

import (
	"net/http"
	"strconv"

	"github.com/SadamIRM/be-smoke-store/models"
	"github.com/SadamIRM/be-smoke-store/services"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService *services.CartService
}

func NewCartHandler() *CartHandler {
	return &CartHandler{cartService: services.NewCartService()}
}

func getUserID(c *gin.Context) (uint, bool) {
	val, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	switch v := val.(type) {
	case float64:
		return uint(v), true
	case uint:
		return v, true
	case int:
		return uint(v), true
	default:
		return 0, false
	}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak ditemukan"})
		return
	}

	items, err := h.cartService.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal mengambil keranjang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    items,
	})
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak ditemukan"})
		return
	}

	var req models.AddCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	item, err := h.cartService.AddToCart(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menambah produk ke keranjang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Produk ditambahkan ke keranjang",
		"data":    item,
	})
}

func (h *CartHandler) UpdateQuantity(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak ditemukan"})
		return
	}

	var req models.UpdateCartQuantityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	item, err := h.cartService.UpdateQuantity(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal memperbarui kuantitas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Kuantitas berhasil diperbarui",
		"data":    item,
	})
}

func (h *CartHandler) RemoveItem(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak ditemukan"})
		return
	}

	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Product ID tidak valid"})
		return
	}

	if err := h.cartService.RemoveItem(userID, uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Produk berhasil dihapus dari keranjang",
	})
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak ditemukan"})
		return
	}

	if err := h.cartService.ClearCart(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal membersihkan keranjang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Keranjang berhasil dibersihkan",
	})
}
