
bookservice : 
	go build -o bs_server  bookservice_server/main.go 
	go build -o bs_client  bookservice_client/main.go 

clean : 
	rm bs_server
	rm bs_client