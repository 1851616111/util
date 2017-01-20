package email

import "testing"

func Test_Validate(t *testing.T) {
	validateEmails := []string{"abc@163.com", "18511175984@163.com", "827384738@qq.com"}
	invalidateEmails := []string{"", "xxx@.com", "@xcxzcsaac.com", "123", "asc", " "}

	for _, email := range validateEmails {
		if err := Validate(email); err != nil {
			t.Fatal(err)
		}
	}

	for _, email := range invalidateEmails {
		if err := Validate(email); err == nil {
			t.Fatal(err)
		}
	}
}
