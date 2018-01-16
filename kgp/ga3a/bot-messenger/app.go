package botMessenger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	."./models"
)

const (
	FACEBOOK_API = "https://graph.facebook.com/v2.6/me/messages?access_token=%s"
)

func SendBill() {
	client := &http.Client{}
	response := Response{
		Recipient: User{
			ID: "1614179648639790",
		},
		Message: Message{
			Attachment: &Attachment{
				Type: "template",
				Payload: Payload{
					TemplateType: "receipt",
					RecipientName: "Nguyen Trung Nghia",
					OrderNumber: "123456789",
					Currency: "VND",
					PaymentMethod: "ATM",
					OrderUrl: "http://petersapparel.parseapp.com/order?order_id=123456",
					Timestamp: time.Now().Unix(),
					Address: Address{
						Street1: "70, Nguyễn trọng Tuyển",
						Street2:"",
						City:"Hồ Chí Minh",
						PostalCode:"94025",
						State:"Phu Nhuan",
						Country: "Vietnam",
					},
					Summary: Summary{
						// Subtotal: 125.000,
						// ShippingCost: 10000,
						// TotalTax: 12000,
						TotalCost: 900000,
					},
					// Adjustments: []Adjustment{
					// 	{
					// 		Name: "asdddd",
					// 		Amount: "1223333",
					// 	},
					// 	{
					// 		Name: "Qưeqưeqưeqưe",
					// 		Amount: "123123123123",
					// 	},
					// },
					Elements: []Element{
						{
							Title:"Áo trắng sơ mi",
							Subtitle:"Hàng việt nam chât lượng cao",
							Quantity:1,
							Price:400000,
							Currency:"VND",
							ImageURL:"https://media3.scdn.vn/img1/2015/1_28/ao-so-mi-trang-cong-so-1m4G3-somitrang4_2k8q7ejtc17pi.jpg",
						},
						{
							Title:"Quần tây đen",
							Subtitle:"100% Soft and Luxurious Cotton",
							Quantity:1,
							Price:500000,
							Currency:"VND",
							ImageURL:"https://media3.scdn.vn/img1/2016/7_4/O5mK9L_simg_b5529c_250x250_maxb.jpg",
						},
					},
				},
			},
		},
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)
	url := fmt.Sprintf(FACEBOOK_API, os.Getenv("PAGE_ACCESS_TOKEN"))
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	log.Println(resp)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func StartServer() {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	log.Println("Starting...")
	// SendBill()
	r := mux.NewRouter()
	r.HandleFunc("/webhook", VerificationEndpoint).Methods("GET")
	r.HandleFunc("/webhook", MessagesEndpoint).Methods("POST")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Println(err)
	}
	log.Println("Server started, litstening 8080.")
}
