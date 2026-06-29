package main

import (
	"github.com/SadamIRM/be-smoke-store/config"
	"github.com/SadamIRM/be-smoke-store/models"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()
	config.InitDatabase()
	products := []models.Product{
		{
			Name:        "Sampoerna Mild",
			Price:       75000,
			Category:    "Rokok",
			Stock:       50,
			Description: "Rokok mild dengan kualitas premium",
			ImageURL:    "https://cdn.ralali.id/assets/img/Libraries/SAMPOERNA-Mild-16-Per-Bungkus_prwDTuNKTCxluEJT_1572210269.png",
		},
		{
			Name:        "Marlboro",
			Price:       300000,
			Category:    "Rokok",
			Stock:       20,
			Description: "Rokok dengan kualitas premium",
			ImageURL:    "https://www.hstliquors.com/cdn/shop/files/Malboro100Red_1800x1800.webp?v=1704656285",
		},
		{
			Name:        "Gudang Garam",
			Price:       1200,
			Category:    "Rokok",
			Stock:       500,
			Description: "Rokok dengan kualitas premium",
			ImageURL:    "https://www.static-src.com/wcsstore/Indraprastha/images/catalog/full/MTA-2582269/gudang-garam_gudang-garam-filter-internasional--12-batang--bungkus-_full03.jpg",
		},
		{
			Name:        "Gudang Garam Signature",
			Price:       85000,
			Category:    "Rokok",
			Stock:       100,
			Description: "Rokok dengan kualitas premium",
			ImageURL:    "https://www.static-src.com/wcsstore/Indraprastha/images/catalog/full//98/MTA-2947382/gudang-garam-signature-black_gudang-garam-signature-black-rokok-filter--12-batang-per-bungkus-_full02.jpg",
		},
		{
			Name:        "Djarum Super",
			Price:       95000,
			Category:    "Rokok",
			Stock:       40,
			Description: "Rokok dengan kualitas premium",
			ImageURL:    "https://image.astronauts.cloud/product-images/2023/12/DjarumSuperRokokBoxFt12Stc_031bf86d-ec79-4f58-b9f0-5845a9af78dd_900x900.png",
		},
		{
			Name:        "Djarum Black",
			Price:       125000,
			Category:    "Rokok",
			Stock:       25,
			Description: "Rokok dengan kualitas premium",
			ImageURL:    "https://victorycigars.ca/wp-content/uploads/2023/03/DjarumBlackBlissEmerald-Package.png",
		},
		{
			Name:        "Sampoerna kretek",
			Price:       65000,
			Category:    "Rokok",
			Stock:       80,
			Description: "Rokok dengan kualitas premium",
			ImageURL:    "https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhYU-oYO1bMotPSCRNmXeCaJzY3WNHfUlgqH3OhE90Iw0-xDuoXEGrf-AOA3crTmPuKxBItkQLaEcM-7hcxSXG5bPmBihzfe-rkHtY7bnMZY4hyywfINty01g7fo1GPW9e4bNT8vxQvAxrt30_0BTT1xvAQH_X5cuuVX-6_pNnIyF4mAzOSnYgIGLDcK3og/s320/20240522_102855.jpg",
		},
		{
			Name:        "Marlboro Ice Blast mild",
			Price:       45000,
			Category:    "Rokok",
			Stock:       60,
			Description: "Rokok dengan kualitas premium",
			ImageURL:    "https://arti-assets.sgp1.cdn.digitaloceanspaces.com/megaswalayan/products/311222d6-f9ed-4589-898f-f90d6c7fcc45.jpg",
		},
		{
			Name:        "Djarum Black mild",
			Price:       150000,
			Category:    "Rokok",
			Stock:       30,
			Description: "Rokok dengan kualitas premium",
			ImageURL:    "https://arti-assets.sgp1.cdn.digitaloceanspaces.com/megaswalayan/products/e909f421-5375-4c77-8511-a1eee896449b.jpg",
		},
		{
			Name:        "Surya pro mild",
			Price:       500,
			Category:    "Rokok",
			Stock:       1000,
			Description: "Rokok dengan kualitas premium",
			ImageURL:    "https://image.astronauts.cloud/product-images/2025/10/8998989300391GudangGaramSuryaProfesionalMildFilterRokokBatang16sticks_9cc146f1-5a5d-4c81-abe9-828408d0f2cb_900x900.png",
		},
	}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
