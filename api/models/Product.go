package models

import (
	"errors"
	"html"
	"strings"
	_"log"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator/v10"
	"github.com/khafido/simple-app-go/api/utils"
)

type Product struct {
	ID_Produk uint32 `gorm:"primary_key;auto_increment" json:"id_produk"`
	Nama_Produk string `gorm:"size:255;not null;unique" json:"nama_produk" validate:"required"`
	Kode_Produk string `gorm:"size:100;not null;unique" json:"kode_produk" validate:"required"`
	Foto_Produk string `gorm:"size:150;not null;" json:"foto_produk"`
	Harga []Price `json:"harga,omitempty"`
}

type Price struct {
	ID_Harga uint32 `gorm:"primary_key;auto_increment" json:"id_harga"`
	Harga_Produk uint32 `gorm:"size:11;" json:"harga_produk"`
	ID_Outlet uint32 `gorm:"size:6;" json:"id_outlet"`
	ID_Produk uint32 `"gorm:"size:6;" json:id_produk;"`
}

type Outlet struct {
	ID_Outlet uint32 `gorm:"primary_key;auto_increment" json:"id_outlet"`
	Outlet string `gorm:"size:100;" json:"nama_outlet"`
}

func (p *Product) Prepare() {
	p.Kode_Produk = html.EscapeString(strings.TrimSpace(p.Kode_Produk))
	p.Nama_Produk = html.EscapeString(strings.TrimSpace(p.Nama_Produk))
	p.Foto_Produk = html.EscapeString(strings.TrimSpace(p.Foto_Produk))
}

func (p *Product) ValidateStruct() map[string]string {
    validate := validator.New()
    err := validate.Struct(p)
    if err != nil {
    	return utils.ParseValidator(err.(validator.ValidationErrors))
    }
    return nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Select("Nama_Produk","Kode_Produk","Foto_Produk").Create(&p).Error
	if err != nil {
		return &Product{}, err
	} else {
		for _, harga := range p.Harga {
			harga.ID_Produk = p.ID_Produk
			err = db.Debug().Create(&harga).Error
			if err != nil {
				return &Product{}, err
			}
		}
	}
	return p, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}

	var newProducts []Product
	for _, product := range products {
		err = db.Debug().Model(Price{}).Where("id_produk = ?", product.ID_Produk).Limit(100).Find(&product.Harga).Error
		newProducts = append(newProducts, product)
	}
	fmt.Println(newProducts)
	return &newProducts, err
}

func (p *Product) FindProductByID(db *gorm.DB, pid uint32) (*Product, error) {
	var err error
	err = db.Debug().Model(Product{}).Where("id_produk = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}

	err = db.Debug().Model(Price{}).Where("id_produk = ?", pid).Limit(100).Find(&p.Harga).Error
	fmt.Println(p)
	if gorm.IsRecordNotFoundError(err) {
		return &Product{}, errors.New("Product Not Found")
	}

	return p, err
}
//
// func (p *Price) FindHarga(db *gorm.DB, pid uint32) (*Price, error){
//
// }

func (p *Product) UpdateProduct(db *gorm.DB, pid uint32) (*Product, error) {
	db = db.Debug().Model(&Product{}).Where("id_produk = ?", pid).Take(&Product{}).UpdateColumns(
		map[string]interface{}{
			"kode_produk": p.Kode_Produk,
			"nama_produk": p.Nama_Produk,
			"foto_produk": p.Foto_Produk,
		},
	)

	var harga Price
	var err = db.Where("id_produk = ?", pid).Delete(&harga).Error
	if(err!=nil){
		return &Product{}, err
	} else {
		for _, hrg := range p.Harga {
			hrg.ID_Produk = pid
			var err = db.Debug().Create(&hrg).Error
			if err != nil {
				return &Product{}, err
			}
		}
	}

	if db.Error != nil {
		return &Product{}, db.Error
	}

	return p, nil
}

func (p *Product) DeleteProduct(db *gorm.DB, pid uint32) (int64, error) {
	db = db.Debug().Model(&Product{}).Where("id_produk = ?", pid).Take(&Product{}).Delete(&Product{})

	db = db.Debug().Model(&Price{}).Where("id_produk = ?", pid).Find(&Price{}).Delete(&Price{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
