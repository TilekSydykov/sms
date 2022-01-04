package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"solar-faza/entity"
	"solar-faza/repository"
	"solar-faza/utils"
	"time"
)

type MainController struct {
	BasicController
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Connection struct {
	Conn *websocket.Conn
}

var connections map[string]Connection

func init() {
	connections = make(map[string]Connection)
}

func (m *MainController) WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	// user connected
	uuids := c.Request.URL.Query()["uuid"]
	if len(uuids) < 1 {
		return
	}
	var uuid = uuids[0]
	connections[uuid] = Connection{conn}

	for {
		_, p, err := conn.ReadMessage()

		if err != nil {
			if _, ok := connections[uuid]; ok {
				delete(connections, uuid)
			}
			return
		}

		messageResolver(p, conn, uuid)
	}

}

func messageResolver(s []byte, conn *websocket.Conn, uuid string) {
	var m Message
	err := json.Unmarshal(s, &m)
	if err != nil {
		print(err)
	}

	switch m.Type {
	case "test":
		err := conn.WriteJSON(Message{Type: "test", Data: "ok"})
		if err != nil {
			return
		}
	}
}

func (m *MainController) MainPage(c *gin.Context) {
	m.FetchUser(c)
	m.SetTitle("Гонкиии!!!!")
	m.Render(c, "index.gohtml")
}

func (m *MainController) LoginPage(c *gin.Context) {
	m.InitData()
	m.SetTitle("Логин")
	m.Render(c, "login.gohtml")
}

func (m *MainController) LoginPost(c *gin.Context) {
	m.InitData()
	m.SetTitle("Логин")
	user := entity.User{}
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	if user.Email == "" || user.Password == "" {
		m.AddData("Message", "Пустые поля")
		m.Render(c, "login.gohtml")
		return
	}
	in, err := utils.SingIn(user.Email, user.Password)
	if err != nil {
		m.AddData("Message", "Неправильные данные для входа")
		m.Render(c, "login.gohtml")
		return
	}
	user.LastLogin = time.Now()
	repository.NewUserRepository().Update(&user)
	utils.GenerateCookies(c, in.AccessToken)
	c.Redirect(302, "/profile")
}

func (m *MainController) RegisterPage(c *gin.Context) {
	m.InitData()
	m.SetTitle("Регистрация")
	m.Render(c, "register.gohtml")
}

func (m *MainController) ProfilePage(c *gin.Context) {
	m.FetchUser(c)
	m.ShouldBeLoggedIn(c)
	m.AddData("Books", repository.NewParticipantRepository().GetByUserId(m.user.Id))
	m.SetTitle("Профиль")
	m.Render(c, "profile.gohtml")
}

func (m *MainController) RegisterPost(c *gin.Context) {
	m.InitData()
	user := entity.User{}
	user.Email = c.PostForm("email_reg")
	user.Password = c.PostForm("password_reg")
	if user.Email == "" || user.Password == "" {
		m.AddData("Message", "Пустые поля")
		m.Render(c, "register.gohtml")
		return
	}
	err := utils.RegisterUser(&user)
	if err != nil {
		m.AddData("Message", "Такой пользователь уже существует")
		m.Render(c, "register.gohtml")
		return
	}
	m.AddData("Message", "Зарегистрирован")
	m.Render(c, "login.gohtml")
}

func (m *MainController) LogOut(c *gin.Context) {
	m.InitData()
	utils.GenerateCookies(c, "")
	c.Redirect(302, "/")
}
