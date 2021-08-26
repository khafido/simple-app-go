package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/khafido/simple-app-go/api/models"
)

var users = []models.User{
	models.User{
		Nama: "User 1",
		Password: "password1",
    Username: "user1",
	},
	models.User{
		Nama: "User 2",
		Password: "password2",
    Username: "user2",
	},
	models.User{
		Nama: "khafido",
		Password: "khaf123",
    Username: "khaf",
	},
}

var products = []models.Product{
	models.Product{
		ID_Produk: 1,
		Nama_Produk: "Produk 1",
		Kode_Produk: "A1",
    Foto_Produk: "https://blog.golang.org/lib/godoc/images/go-logo-blue.svg",
	},
	models.Product{
		ID_Produk: 2,
		Nama_Produk: "Produk 2",
		Kode_Produk: "A2",
    Foto_Produk: "https://blog.golang.org/lib/godoc/images/go-logo-blue.svg",
	},
}

var prices = []models.Price{
	models.Price{
		ID_Harga: 1,
		Harga_Produk: 1000,
		ID_Outlet: 1,
		ID_Produk: 1,
	},
	models.Price{
		ID_Harga: 2,
		Harga_Produk: 2000,
		ID_Outlet: 2,
		ID_Produk: 1,
	},
	models.Price{
		ID_Harga: 3,
		Harga_Produk: 3000,
		ID_Outlet: 1,
		ID_Produk: 2,
	},
	models.Price{
		ID_Harga: 4,
		Harga_Produk: 4000,
		ID_Outlet: 2,
		ID_Produk: 2,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}, &models.Product{}, &models.Price{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Product{}, &models.Price{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
  //
	// err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	// PRODUCTS
	for i, _ := range products {
		err = db.Debug().Model(&models.User{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		// posts[i].AuthorID = users[i].ID

		// err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
	}

	for i, _ := range prices {
		err = db.Debug().Model(&models.Price{}).Create(&prices[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
