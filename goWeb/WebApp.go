package main
import (
	"net/http"
	"fmt"
)
func handlerFunc(w http.ResponseWriter, r * http.Request)  {
	fmt.Fprint(w, "<html>" +
		"<head>" +
		"<style>" +
		"h1{" +
		"color: red;" +
		"}" +
		"button{" +
		"padding: 15px;" +
		"margin-left: 20px;" +
		"font-weight: bold" +
		"}" +
		"</style>" +
		"</head>" +
		"<body>" +
		"<h1>Welcome In GOLang Server!" +
		"<button onClick=\"show()\">Click Me</button>" +
		"<script>" +
		"function show(){" +
		"alert('Clicked');" +
		"}" +
		"</script>" +
		"</body>" +
		"</html> ")
}
func main()  {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000",nil)
}