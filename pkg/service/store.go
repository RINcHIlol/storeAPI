package service

import (
	"fmt"
	"gopkg.in/mail.v2"
	"os"
	"storeApi"
	"storeApi/pkg/repository"
)

type StoreService struct {
	repo repository.Store
}

func NewStoreService(repo repository.Store) *StoreService {
	return &StoreService{repo: repo}
}

func (s *StoreService) AddNewProduct(product storeApi.Product) (bool, error) {
	return s.repo.CreateProduct(product)
}

func (s *StoreService) GetProducts() ([]storeApi.Product, error) {
	return s.repo.GetProducts()
}

func (s *StoreService) GetProductById(productId int) (storeApi.Product, error) {
	return s.repo.GetProductById(productId)
}

//func (s *StoreService) BuyProduct(productId int) (bool, error) {
//	product, err := s.repo.DeleteProductById(productId)
//	if err != nil {
//		return false, err
//	}
//
//	// Сообщение покупателю
//	buyerMsg := fmt.Sprintf("It is your order number: %d, Save it, you need to show it to cap in future.", product.OrderNum)
//	buyer := mail.NewMessage()
//	buyer.SetHeader("From", "galimatron229@gmail.com")
//	buyer.SetHeader("To", "workemailsvarka123@gmail.com")
//	buyer.SetHeader("Subject", "Shopping in Svarka_Shop")
//	buyer.SetBody("text/plain", buyerMsg)
//
//	// Сообщение себе (владелец магазина)
//	ownerMsg := fmt.Sprintf("Customer bought smth with ID: %d. Name: %s. OrderNum: %d. Price: %.2f.", productId, product.Name, product.OrderNum, product.Price)
//	owner := mail.NewMessage()
//	owner.SetHeader("From", "galimatron229@gmail.com")
//	owner.SetHeader("To", "workemailsvarka123@gmail.com")
//	owner.SetHeader("Subject", "Shopping in Svarka_Shop")
//	owner.SetBody("text/plain", ownerMsg)
//
//	// Настройка SMTP-сервера
//	d := mail.NewDialer("smtp.gmail.com", 587, "galimatron229@gmail.com", "wplgvcwvcsvxxfxp")
//
//	// Отправка двух писем
//	if err := d.DialAndSend(buyer, owner); err != nil {
//		return false, err
//	}
//
//	return true, nil
//}

func (s *StoreService) BuyProduct(productId int) (bool, error) {
	product, err := s.repo.DeleteProductById(productId)
	if err != nil {
		return false, err
	}

	// 1. Сохраняем временный файл картинки
	tmpFile, err := os.CreateTemp("", "product-*.jpg")
	if err != nil {
		return false, err
	}
	defer os.Remove(tmpFile.Name()) // удаляем файл после отправки
	defer tmpFile.Close()

	_, err = tmpFile.Write(product.Image)
	if err != nil {
		return false, err
	}

	// 2. Создаем письмо покупателю с HTML и встроенной картинкой
	buyer := mail.NewMessage()
	buyer.SetHeader("From", "galimatron229@gmail.com")
	buyer.SetHeader("To", "workemailsvarka123@gmail.com")
	buyer.SetHeader("Subject", "Shopping in Svarka_Shop")

	// Устанавливаем CID вручную
	cid := "product-image"
	buyer.Embed(tmpFile.Name(), mail.SetHeader(map[string][]string{
		"Content-ID": {"<" + cid + ">"},
	}))

	buyerBody := fmt.Sprintf(`
		<h2>Thank you for your purchase!</h2>
		<p><strong>Order Number:</strong> %d</p>
		<p><strong>Product Name:</strong> %s</p>
		<p><strong>Price:</strong> %.2f</p>
		<p><strong>Description:</strong> %s</p>
		<p>Here is your product:</p>
		<img src="cid:%s" alt="Product Image"/>
	`, product.OrderNum, product.Name, product.Price, product.Description, cid)

	buyer.SetBody("text/html", buyerBody)

	// 3. Письмо владельцу (обычное текстовое)
	ownerMsg := fmt.Sprintf("Customer bought smth with ID: %d. Name: %s. OrderNum: %d. Price: %.2f.",
		productId, product.Name, product.OrderNum, product.Price)

	owner := mail.NewMessage()
	owner.SetHeader("From", "galimatron229@gmail.com")
	owner.SetHeader("To", "workemailsvarka123@gmail.com")
	owner.SetHeader("Subject", "Shopping in Svarka_Shop")
	owner.SetBody("text/plain", ownerMsg)

	// 4. Отправляем оба письма
	d := mail.NewDialer("smtp.gmail.com", 587, "galimatron229@gmail.com", "wplgvcwvcsvxxfxp")

	if err := d.DialAndSend(buyer, owner); err != nil {
		return false, err
	}

	return true, nil
}

func (s *StoreService) UpdateProductById(productId int, product storeApi.Product) (bool, error) {
	return s.repo.UpdateProductById(productId, product)
}
