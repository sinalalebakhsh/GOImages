package main

import (
	"cms"
	"net/http"
	// "os"
)

func main()  {
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandlesNew)
	http.HandleFunc("/page", cms.ServePage)
	http.HandleFunc("/post", cms.ServePost)
	http.ListenAndServe(":3000", nil)
	// p := cms.Page{	
	// 	Title: "Hello, world!",
	// 	Content: "This is body of our website",	
	// } 
	
	// cms.Templ.ExecuteTemplate(os.Stdout, "index", p)
}

/* Output:

<!DOCTYPE html>
<html>
	<head>
		<title>
			Hello, world!
		</title>
	</head>
	<body>
		This is body of our website
	</body>
</html>

*/