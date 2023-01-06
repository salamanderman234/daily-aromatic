package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	model "github.com/salamanderman234/daily-aromatic/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Seeder struct {
	users    []model.User
	products []model.Product
	reviews  []model.Review
}

func (s *Seeder) Generate(conn *gorm.DB) error {
	s.userSeed(5)
	s.productSeed(10)
	s.reviewSeed(20)

	conn.Transaction(func(tx *gorm.DB) error {
		err := conn.Create(s.users)
		if err.Error != nil {
			tx.Rollback()
			return err.Error
		}
		err = conn.Create(s.products)
		if err.Error != nil {
			tx.Rollback()
			return err.Error
		}
		err = conn.Create(s.reviews)
		if err.Error != nil {
			tx.Rollback()
			return err.Error
		}
		return nil
	})
	return nil
}

func (s *Seeder) userSeed(amount int) {
	for i := 0; i < amount; i++ {
		res, err := http.Get("https://random-data-api.com/api/users/random_user")
		if err != nil {
			panic(err)
		}
		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		result := model.User{}
		json.Unmarshal(responseData, &result)
		pass, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
		result.Password = string(pass)
		s.users = append(s.users, result)
	}
}

func (s *Seeder) productSeed(amount int) {
	pabrikan := []string{"Napura", "Putri Bulan", "Kemenyan"}
	warna := []string{"Hitam", "Coklat", "Biru"}
	aroma := []string{"Kemenyan", "Rampe", "Jepun"}
	for i := 0; i < amount; i++ {
		result := model.Product{
			Nama:       "product" + strconv.Itoa(i),
			ImageUrl:   "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/6/19/f5c15304-a4e6-4334-b02f-4b3ee27c1501.jpg",
			Pabrikan:   pabrikan[rand.Intn(3)],
			Warna:      warna[rand.Intn(3)],
			LamaNyala:  3,
			Asal:       "Bali",
			JenisDupa:  "Aromatic",
			JenisAbu:   "tidak panas",
			Jumlah:     20,
			Aroma:      aroma[rand.Intn(3)],
			JenisAroma: "Aromatic",
			Panjang:    30,
			Berat:      400,
			Rating:     0,
		}
		s.products = append(s.products, result)
	}
}

func (s *Seeder) reviewSeed(amount int) {
	pl := len(s.products)
	ul := len(s.users)
	if pl > 0 && ul > 0 {
		for i := 0; i < amount; i++ {
			user := s.users[rand.Intn(ul)]
			product := s.products[rand.Intn(pl)]
			rate := float64(rand.Intn(6))
			user.ReviewTotal++
			product.JumlahReview++
			product.Rating = (product.TotalRate + rate) / float64(product.JumlahReview)
			review := model.Review{
				Product: product,
				User:    user,
				Rate:    rate,
				Comment: "loreflasdhfalksddjfhhasdlkjfhasdfuhlasdkjfhaslkjdhhfasdkjhfaskjfh akksahfkjashdfakjh fskadljjhhlkj",
			}
			s.reviews = append(s.reviews, review)
		}
	}
}
