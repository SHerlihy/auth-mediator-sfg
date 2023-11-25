package routes

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func Setup(router fiber.Router) {
//	router.Get("/user/logout")
//	router.Get("/user/auth-status")

    //authServiceURL := os.Getenv("AUTH_SERVICE_URL")
    //reqURL := fmt.Sprintf("%s%s", authServiceURL, "/user/signup")
    //router.Post("/user/signup", preAuthCall, proxy.DomainForward(reqURL,"http://localhost:5000" ))
    //router.Post("/user/signup", preAuthCall, proxy.Forward(reqURL))
    router.Post("/user/signup", preAuthCall, proxyDo)
    //router.Post("/user/signup", preAuthCall, proxyAuthService)
//	router.Post("/user/login")
}

func proxyDo(c *fiber.Ctx)error{
    authServiceURL := os.Getenv("AUTH_SERVICE_URL")
    reqURL := fmt.Sprintf("%s%s", authServiceURL, "/user/signup")

    if err := proxy.Do(c,reqURL); err!=nil{
        return err
    }

    c.Response().Header.Set("Access-Control-Allow-Origin", "http://localhost:5173")
    c.Response().Header.Set("Access-Control-Allow-Credentials", "true")

    return nil
}

func proxyAuthService(c *fiber.Ctx) error {
    authServiceURL := os.Getenv("AUTH_SERVICE_URL")
    reqURL := fmt.Sprintf("%s%s", authServiceURL, "/api/v1/user/signup")
    var buff bytes.Buffer

    buff.Write(c.BodyRaw())

    fmt.Println(os.Stdout, fmt.Sprint("pre post req"))
    resp, err := http.Post(reqURL, "application/json", &buff)
    if err != nil {
    	// handle error
        fmt.Println(os.Stdout, fmt.Sprint("Errror"))
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)

    fmt.Println(os.Stdout, fmt.Sprint("post body %v", body))

    respCookies := resp.Cookies()

    for _, respCookie := range respCookies {
        c.CookieParser(respCookie)
    }

    return c.JSON(body)
}

func preAuthCall(c *fiber.Ctx)error{
    fmt.Println(os.Stdout, fmt.Sprint("Pre call response"))

    return c.Next()
}
