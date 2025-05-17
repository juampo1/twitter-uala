package db

import (
	"log"
	"math/rand"
	follow "twitter-uala/internal/domain/follow/models"
	"twitter-uala/internal/domain/user/models"

	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) Seed() {
	users := []models.User{
		{
			ID:       "1",
			Username: "@juanito",
		},
		{
			ID:       "2",
			Username: "@mary",
		},
		{
			ID:       "3",
			Username: "@carlitos",
		},
		{
			ID:       "4",
			Username: "@anita",
		},
		{
			ID:       "5",
			Username: "@pedrito",
		},
		{
			ID:       "6",
			Username: "@sofi",
		},
		{
			ID:       "7",
			Username: "@luisito",
		},
		{
			ID:       "8",
			Username: "@lau",
		},
		{
			ID:       "10",
			Username: "@mike",
		},
		{
			ID:       "11",
			Username: "@carmi",
		},
		{
			ID:       "12",
			Username: "@joseto",
		},
		{
			ID:       "13",
			Username: "@isa",
		},
		{
			ID:       "14",
			Username: "@dieguito",
		},
		{
			ID:       "15",
			Username: "@vale",
		},
		{
			ID:       "16",
			Username: "@sebas",
		},
		{
			ID:       "17",
			Username: "@gabi",
		},
		{
			ID:       "18",
			Username: "@manu",
		},
		{
			ID:       "19",
			Username: "@cami",
		},
		{
			ID:       "20",
			Username: "@rodrigo",
		},
	}

	// Crear mapa para almacenar relaciones de seguimiento y evitar duplicados
	followMap := make(map[string]map[string]bool)

	// Inicializar el mapa
	for _, user := range users {
		followMap[user.ID] = make(map[string]bool)
	}

	// Asignar seguidores a cada usuario
	for _, user := range users {
		// Generar un número aleatorio de seguidores entre 5 y 10
		numFollowers := rand.Intn(6) + 5

		// insertar usuario en la base de datos
		if err := s.db.Create(&user).Error; err != nil {
			log.Fatalf("Error al insertar usuario %s: %v", user.Username, err)
		}

		// Obtener una lista de otros usuarios (excluyendo al usuario actual)
		var otherUsers []models.User
		for _, u := range users {
			if u.ID != user.ID {
				otherUsers = append(otherUsers, u)
			}
		}

		// Mezclar la lista de otros usuarios
		rand.Shuffle(len(otherUsers), func(i, j int) {
			otherUsers[i], otherUsers[j] = otherUsers[j], otherUsers[i]
		})

		// Ajustar si el número de seguidores excede el número de usuarios disponibles
		if numFollowers > len(otherUsers) {
			numFollowers = len(otherUsers)
		}

		// Seleccionar los primeros 'numFollowers' usuarios como seguidores
		followers := otherUsers[:numFollowers]

		// Asignar seguidores al usuario actual
		for _, follower := range followers {
			// Verificar si ya existe la relación para evitar duplicados
			if !followMap[user.ID][follower.ID] {
				// Crear la relación de seguimiento en la base de datos
				follow := follow.Follow{
					UserID:     user.ID,
					FollowedID: follower.ID,
				}
				if err := s.db.Create(&follow).Error; err != nil {
					log.Printf("Error al insertar seguimiento de %s a %s: %v", follower.Username, user.Username, err)
				}

				// Marcar la relación como existente
				followMap[user.ID][follower.ID] = true
			}
		}
	}
}
