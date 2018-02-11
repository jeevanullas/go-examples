package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	http.HandleFunc("/HealthCheck", helloRequest)
	http.HandleFunc("/", getRequest)

	// start web server
	log.Println("Listening on http://localhost:5000/")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// basic handler for /HealthCheck request
func helloRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "It works!")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	svc := dynamodb.New(sess)
	req := &dynamodb.DescribeTableInput{
		TableName: aws.String("person"),
	}

	result, err := svc.DescribeTable(req)
	if err != nil {
		fmt.Printf("%s", err)
	}

	value := *(result.Table.TableStatus)
	fmt.Println(value)

	return
}

// this function simply serves up the file requested, or
// an index list if a directory is requested
func getRequest(w http.ResponseWriter, r *http.Request) {
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}
