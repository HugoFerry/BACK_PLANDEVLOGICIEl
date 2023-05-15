package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Structure de données pour stocker les informations sur les étages
type Etage struct {
	ID     int    `json:"id"`
	Nom    string `json:"nom"`
	Salles []Salle
}

// Structure de données pour stocker les informations sur les salles
type Salle struct {
	ID           int           `json:"id"`
	Nom          string        `json:"nom"`
	Disponible   bool          `json:"disponible"`
	Reservations []Reservation `json:"reservations"`
}

// Structure de données pour stocker les informations sur les réservations
type Reservation struct {
	ID      int    `json:"id"`
	SalleID int    `json:"salle_id"`
	UserID  int    `json:"user_id"`
	Debut   string `json:"debut"`
	Fin     string `json:"fin"`
}

var etages []Etage
var reservations []Reservation

func main() {
	router := gin.Default()

	// Route pour obtenir la liste des étages du bâtiment
	router.GET("/etages", func(c *gin.Context) {
		c.JSON(http.StatusOK, etages)
	})

	// Route pour obtenir la liste des salles d'un étage
	router.GET("/etages/:id_etage/salles", func(c *gin.Context) {
		idEtage, err := strconv.Atoi(c.Param("id_etage"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID d'étage invalide"})
			return
		}
		if idEtage >= len(etages) || idEtage < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID d'étage invalide"})
			return
		}
		c.JSON(http.StatusOK, etages[idEtage].Salles)
	})

	// Route pour obtenir les informations détaillées d'une salle et ses disponibilités
	router.GET("/salles/:id_salle", func(c *gin.Context) {
		idSalle, err := strconv.Atoi(c.Param("id_salle"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de salle invalide"})
			return
		}
		var salle Salle
		for _, etage := range etages {
			for _, s := range etage.Salles {
				if s.ID == idSalle {
					salle = s
					break
				}
			}
		}
		if salle.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de salle invalide"})
			return
		}
		c.JSON(http.StatusOK, salle)
	})

	// Initialisation des données
	//initEtages()
	//initSalles()

	// Lancement du serveur sur le port 8080
	router.Run(":8080")
}
