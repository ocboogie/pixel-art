package models

// func TestUserValidate(t *testing.T) {
// 	user := &User{}

// 	if err := user.Validate(); err == nil {
// 		t.Error("Expected an error but didn't receive one")
// 	}

// 	user.ID = "60aaf13d-8ddc-403b-ba42-960e18a22f6a"
// 	user.Email = "foo@bar.com"
// 	user.Password = "$argon2id$v=19$m=4096,t=192,p=12$enR8oMuzdWr8gBbGf8el3iSts7e2cDadP3gst5adBG0$PQlPYLvmSG5E12n97hpxLmZXqiaCoIvPGUAmx3WzW/M"

// 	if err := user.Validate(); err != nil {
// 		t.Error("Didn't expect an error: " + err.Error())
// 	}
// }

// func TestUserDefaults(t *testing.T) {
// 	user := &User{}

// 	user.Defaults()

// 	if len(user.ID) != 36 {
// 		t.Error("Didn't generate id")
// 	}
// }
