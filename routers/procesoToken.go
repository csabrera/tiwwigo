package routers

import (
	"errors"
	"strings"

	"github.com/csabrera/tiwwigo/bd"
	"github.com/csabrera/tiwwigo/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor usado en todos los EndPoints*/
var Email string

/*IDUsuario devuelto del modelo, que se usara en todos los Endpoints*/
var IDUsuario string

/*ProcesoToken para extraer sus valoroes*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdeDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token no valido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}
	return claims, false, string(""), err
}
