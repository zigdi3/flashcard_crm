// flashcard_crm.go

package main

import (
	"net/http"

	
    "github.com/gin-gonic/gin"

)

type Flashcard struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`         // Updated field for Name
    GithubURL   string `json:"github_url"`   // Updated field for GitHub URL
    LinkedinURL string `json:"linkedin_url"` // Updated field for LinkedIn URL
    QRCode      string `json:"qrcode"`
}

var flashcards []Flashcard

func main() {
    r := gin.Default()
    
    // ... existing routes ...

    r.GET("/flashcards", getFlashcards)
    r.POST("/flashcards", createFlashcard)

    // ... existing routes ...

    r.Run()
}

// ... existing functions ...

// getFlashcards retrieves all flashcards
func getFlashcards(c *gin.Context) {
    c.JSON(http.StatusOK, flashcards)
}

// createFlashcard adds a new flashcard
func createFlashcard(c *gin.Context) {
    var newFlashcard Flashcard
    if err := c.ShouldBindJSON(&newFlashcard); err == nil {
        flashcards = append(flashcards, newFlashcard)
        c.JSON(http.StatusCreated, newFlashcard)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}