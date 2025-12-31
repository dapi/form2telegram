package formatter

import (
	"testing"
)

func TestFormatForm_SingleField(t *testing.T) {
	form := &FormData{
		Answers: []Answer{
			{Key: "email", Value: "test@example.com"},
		},
	}

	result := FormatForm(form)

	expected := "*email:* test@example.com"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestFormatForm_MultipleFields(t *testing.T) {
	form := &FormData{
		Answers: []Answer{
			{Key: "Имя", Value: "Иван"},
			{Key: "Email", Value: "ivan@example.com"},
			{Key: "Телефон", Value: "+7 999 123 45 67"},
		},
	}

	result := FormatForm(form)

	expected := "*Имя:* Иван\n*Email:* ivan@example.com\n*Телефон:* +7 999 123 45 67"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestFormatForm_EmptyAnswers(t *testing.T) {
	form := &FormData{
		Answers: []Answer{},
	}

	result := FormatForm(form)

	if result != "" {
		t.Errorf("got %q, want empty string", result)
	}
}

func TestFormatForm_EscapesMarkdown(t *testing.T) {
	form := &FormData{
		Answers: []Answer{
			{Key: "comment", Value: "Hello *world* and _test_"},
		},
	}

	result := FormatForm(form)

	expected := "*comment:* Hello \\*world\\* and \\_test\\_"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}
