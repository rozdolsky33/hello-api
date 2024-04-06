package translation_test

import (
	"hello-api/translation"
	"testing"
)

func TestTranslate(t *testing.T) {
	//Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "buy",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "buy",
			Language:    "german",
			Translation: "",
		},
		{
			Word:        "hello",
			Language:    "German",
			Translation: "hallo",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello ",
			Language:    "german",
			Translation: "hallo",
		},
	}
	underTest := translation.NewStaticService()
	for _, test := range tt {
		//Act
		res := underTest.Translate(test.Word, test.Language)

		//Assert
		if res != test.Translation {
			t.Errorf(`expected "%s" to be "%s" from "%s" but received "%s"`, test.Word, test.Language, test.Translation, res)
		}

	}

}
