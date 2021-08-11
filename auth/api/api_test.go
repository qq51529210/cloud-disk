package api

// import (
// 	"bytes"
// 	"net/http"
// 	"testing"

// 	"github.com/qq51529210/cloud-service/authentication/db"
// )

// type testResponseWriter struct {
// 	header     http.Header
// 	statusCode int
// 	body       bytes.Buffer
// }

// func (r *testResponseWriter) Reset() {
// 	r.header = make(http.Header)
// 	r.body.Reset()
// }

// func (r *testResponseWriter) Header() http.Header {
// 	return r.header
// }

// func (r *testResponseWriter) Write(b []byte) (int, error) {
// 	return r.body.Write(b)
// }

// func (r *testResponseWriter) WriteString(s string) (int, error) {
// 	return r.body.WriteString(s)
// }

// func (r *testResponseWriter) WriteHeader(statusCode int) {
// 	r.statusCode = statusCode
// }

// func testInitDB(t *testing.T) {
// 	err := db.InitPhoneNumberRedis("")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = db.InitTokenRedis("")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = db.InitMysql("root:123123@tcp(127.0.0.1:3306)/anthentication")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
