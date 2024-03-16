package worldcommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbedCode(t *testing.T) {
	t.Run("NewEmbedCode", func(t *testing.T) {
		t.Run("Should throw error when embed code has no src", func(t *testing.T) {
			_, err := NewEmbedCode("<iframe></iframe>")
			assert.Error(t, err)
		})

		t.Run("Should throw error when embed code has no closing tag", func(t *testing.T) {
			_, err := NewEmbedCode("<iframe src=\"htps://google.com\">")
			assert.Error(t, err)
		})

		t.Run("Should throw error when embed code does not start or end with iframe tag", func(t *testing.T) {
			_, err := NewEmbedCode(" <iframe src=\"htps://google.com\"></iframe>")
			assert.Error(t, err)

			_, err = NewEmbedCode("<iframe src=\"htps://google.com\"></iframe> ")
			assert.Error(t, err)
		})

		t.Run("Should not throw error when embed code is valid", func(t *testing.T) {
			_, err := NewEmbedCode("<iframe src=\"https://quizlet.com/249824575/learn/embed?i=181ujs&x=1jj1\" height=\"500\" width=\"100%\" style=\"border:0\"></iframe>")
			assert.NoError(t, err)
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("Should return the source embed code", func(t *testing.T) {
			embedCodeString := "<iframe src=\"https://quizlet.com/249824575/learn/embed?i=181ujs&x=1jj1\" height=\"500\" width=\"100%\" style=\"border:0\"></iframe>"
			embedCode, err := NewEmbedCode(embedCodeString)
			if err != nil {
				assert.NoError(t, err)
			}

			assert.Equal(t, embedCode.String(), embedCodeString)
		})
	})
}
