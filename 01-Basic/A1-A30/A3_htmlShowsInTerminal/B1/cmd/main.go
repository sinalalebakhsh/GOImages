package main

import (
	cms "cmsOne"
	"os"
)

func main()  {
	p := cms.Page{	
		Title: "Hello, world!",
		Content: "This is body of our website",	
	} 
	
	cms.Templ.ExecuteTemplate(os.Stdout, "index", p)
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