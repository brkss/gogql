package token

import (
	"fmt"
	"net/http"
)

//func SendRefreshToken()

func RefreshToken(w http.ResponseWriter, r *http.Request){
		
	w.Header().Set("Content-Type", "application/json")
	token, err :=  r.Cookie("rtk")
	if err != nil {
		w.Write([]byte(`{"error" : ""}`))		
		return;		
	}
	fmt.Println("token : ", token.Value)
	w.Write([]byte(`{"error" : ""}`))		
}
