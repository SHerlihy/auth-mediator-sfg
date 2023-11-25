package routes

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func Setup(router fiber.Router) {
    router.Get("/user/auth-status", func(c *fiber.Ctx)error{
        return authProxy(c,"/user/auth-status")
    })
    router.Get("/user/logout", func(c *fiber.Ctx)error{
        return authProxy(c,"/user/logout")
    })
    
    router.Post("/user/signup", func(c *fiber.Ctx)error{
        return authProxy(c,"/user/signup")
    })
    router.Post("/user/login", func(c *fiber.Ctx)error{
        return authProxy(c,"/user/login")
    })
}

func authProxy(c *fiber.Ctx, endpoint string) error {
    clientURL := os.Getenv("CLIENT_URL")
    authServiceURL := os.Getenv("AUTH_SERVICE_URL")
    reqURL := fmt.Sprintf("%s%s", authServiceURL, endpoint)

    if err := proxy.Do(c,reqURL); err!=nil{
        return err
    }

    c.Response().Header.Set("Access-Control-Allow-Origin", clientURL)
    c.Response().Header.Set("Access-Control-Allow-Credentials", "true")

    return nil
}
