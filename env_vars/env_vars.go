package env_vars

import (
	"os"
)

func ProdEnv(){
    authServiceURL := "http://localhost:5001/api/v1"

    os.Setenv("AUTH_SERVICE_URL", authServiceURL)
    os.Setenv("ALLOWED_ORIGINS", "http://localhost:5173")
}

func DevEnv(){
    authServiceURL := "http://localhost:5001/api/v1"

    os.Setenv("CLIENT_URL", "http://localhost:5173")
    os.Setenv("AUTH_SERVICE_URL", authServiceURL)
    os.Setenv("ALLOWED_ORIGINS", "http://localhost:5173")
}
