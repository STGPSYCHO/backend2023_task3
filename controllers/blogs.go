package controllers

import (
	"net/http"

	// "log"
	// "strconv"
	// "strings"

	"github.com/STGPSYCHO/backend2023_task3/models"
	"github.com/gin-gonic/gin"
)

// GET /posts/:id
// Получаем блог по id
func GetPost(context *gin.Context) {
	var blogs models.Blog
	var comments []models.Comment

	if err := models.DB.Where("id = ?", context.Param("id")).First(&blogs).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Запись не существует"})
		return
	}
	if err := models.DB.Where("blog_id = ?", context.Param("id")).Find(&comments).Error; err != nil {
	}
	// models.DB.Joins("Comment").Joins("User").Find()
	context.HTML(
		http.StatusOK,
		"post.html",
		gin.H{
			"ID":       blogs.ID,
			"Text":     blogs.Text,
			"Title":    blogs.Title,
			"comments": comments,
		},
	)
}

// type AddProductInput struct {
// 	Product  uint   `json:"product"`
// 	Quantity string `json:"quantity"`
// }
// type CartProducts struct {
// 	ID       uint   `json:"id" gorm:"primary_key"`
// 	Name     string `json:"name"`
// 	Price    string `json:"price"`
// 	Picture  string `json:"pic_link"`
// 	Quantity string `json:"quantity"`
// }

// // GET /products
// // Получаем список всех продуктов
// func GetAllProducts(context *gin.Context) {

// 	var products []models.Product

// 	if err := models.DB.Find(&products).Error; err != nil {
// 		context.JSON(http.StatusNotFound, gin.H{"error": "Нет подходящих записей"})
// 		return
// 	}

// 	context.HTML(
// 		http.StatusOK,
// 		"products.html",
// 		gin.H{"products": products},
// 	)
// }

// // POST /add-cart
// // Добавляем корзину в куку
// func AddProductToCart(context *gin.Context) {

// 	var cart []string
// 	cookie, err := context.Cookie("cart")
// 	if err != nil {
// 		if err != http.ErrNoCookie {
// 			context.JSON(http.StatusBadRequest, gin.H{"Bad Request": "Не могу распарсить куки"})
// 		}
// 	}

// 	qu, _ := context.GetPostForm("quantity")
// 	intqu, err := strconv.Atoi(qu)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	pr, _ := context.GetPostForm("product")
// 	idpr, err := strconv.Atoi(pr)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	change_string := ""
// 	if cookie != "" { // 1:1,2:1,3:5,4:20,5:10
// 		cart = strings.Split(cookie, ",") // 1:1 2:1 3:5
// 		for idx, str := range cart {      // 1:1
// 			split_str := strings.Split(str, ":")      // 1 1
// 			id, _ := strconv.Atoi(split_str[0])       // 1
// 			quantity, _ := strconv.Atoi(split_str[1]) // 1
// 			if idpr == id {
// 				quantity += intqu
// 				change_string = strconv.Itoa(idpr) + ":" + strconv.Itoa(quantity)
// 				cart[idx] = change_string
// 			}
// 		}
// 		if change_string == "" {
// 			cart = append(cart, pr+":"+qu)
// 		}
// 		cookie = strings.Join(cart, ",")
// 		change_string = ""
// 	} else {
// 		cookie = pr + ":" + qu
// 	}

// 	context.SetCookie("cart", cookie, 3600, "/", context.Request.URL.Hostname(), false, true)
// 	// context.JSON(http.StatusOK, gin.H{"Успех": cookie})
// 	context.Redirect(http.StatusMovedPermanently, "/products")
// }

// // POST /remove-cart
// // Удаляем из корзины куку
// func RemoveProductFromCart(context *gin.Context) {

// 	cookie, err := context.Cookie("cart")
// 	if err != nil {
// 		if err != http.ErrNoCookie {
// 			context.JSON(http.StatusBadRequest, gin.H{"Bad Request": "Не могу распарсить куки"})
// 			return
// 		}
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "Нет таких товаров в корзине"})
// 		return
// 	}

// 	qu, _ := context.GetPostForm("quantity")
// 	pr, _ := context.GetPostForm("product")

// 	if cookie != "" {
// 		cart := strings.Split(cookie, ",")
// 		idx := Find(cart, pr+":"+qu)
// 		cart = remove(cart, idx)
// 		cookie = strings.Join(cart, ",")
// 	} else {
// 		context.JSON(http.StatusBadRequest, gin.H{"Bad Request": "Пустые куки"})
// 		return
// 	}
// 	context.SetCookie("cart", cookie, 3600, "/", context.Request.URL.Hostname(), false, true)
// 	context.Redirect(http.StatusMovedPermanently, "/cart")
// }

// // GET /cookie
// // Удаляем из корзины куку
// func GetCookie(context *gin.Context) {

// 	cookie, err := context.Cookie("cart")
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"cookie": err.Error()})
// 		return
// 	}
// 	context.JSON(http.StatusOK, gin.H{"cookie": cookie})
// }

// // GET /cart
// // Получаем корзину товаров
// func GetCart(context *gin.Context) {

// 	var arr []AddProductInput       // массив объектов инпута, которые пришли из кук
// 	var cartProducts []CartProducts // массив конечных продуктов в корзине
// 	var products []models.Product

// 	if err := models.DB.Find(&products).Error; err != nil {
// 		context.JSON(http.StatusNotFound, gin.H{"error": "Нет подходящих записей"})
// 		return
// 	}

// 	cookie, err := context.Cookie("cart")
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"cookie": "Ошибка получения корзины"})
// 		return
// 	}
// 	cart := strings.Split(cookie, ",")

// 	for _, value := range cart { // формирование массив объектов инпута, которые пришли из кук
// 		b, a, f := strings.Cut(value, ":")
// 		if f {
// 			val, _ := strconv.Atoi(b)
// 			input := AddProductInput{Product: uint(val), Quantity: a}
// 			arr = append(arr, input)
// 		}
// 	}
// 	for _, value := range arr {
// 		for _, value2 := range products {
// 			if value.Product == value2.ID {
// 				input := CartProducts{ID: value2.ID, Name: value2.Name, Picture: value2.Picture, Price: value2.Price, Quantity: value.Quantity}
// 				cartProducts = append(cartProducts, input)
// 			}
// 		}
// 	}
// 	if cartProducts != nil {
// 		context.HTML(
// 			http.StatusOK,
// 			"cart.html",
// 			gin.H{"cart": cartProducts},
// 		)
// 	} else {
// 		context.HTML(
// 			http.StatusOK,
// 			"cartEmpty.html",
// 			gin.H{"cart": "Корзина пуста"},
// 		)
// 	}

// }

// func Find(a []string, x string) int {
// 	for i, n := range a {
// 		if x == n {
// 			return i
// 		}
// 	}
// 	return len(a)
// }
// func remove(slice []string, s int) []string {
// 	return append(slice[:s], slice[s+1:]...)
// }
