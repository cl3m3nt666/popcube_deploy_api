package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"os/exec"
	"math/rand"
	"time"
	"io"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome for popcube deploy api!\n")
}

func deploy_create(w http.ResponseWriter, r *http.Request) {

	var create_stack Create_stack
	rand.Seed(time.Now().UnixNano())
	if r.Body == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader( 400)
		json.NewEncoder(w).Encode(jsonErr{Code: 400, Text: "Please send a request body"})
		return
	}
	err := json.NewDecoder(r.Body).Decode(&create_stack)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader( 400)
		json.NewEncoder(w).Encode(jsonErr{Code: 400, Text: "Please send a request body : "+ err.Error()} )
		return
	}
	if create_stack.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(jsonErr{Code: 400, Text: "Name is empty"})
		return
	}

	//Check if dir exist org dir
	if _, err := os.Stat(os.Getenv("DEFAULT_ORG_PATH")); os.IsNotExist(err) {
		os.Mkdir(os.Getenv("DEFAULT_ORG_PATH"), 0755)
	}
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Tep directory"})
		return
	}
	// Create org dir
	_ = os.Mkdir(os.Getenv("DEFAULT_ORG_PATH"), 0755)

	//Check if dir exist org dir
	if _, err := os.Stat(os.Getenv("DEFAULT_ORG_PATH")+"/"+create_stack.Name); os.IsNotExist(err) {
		os.Mkdir(os.Getenv("DEFAULT_ORG_PATH"), 0755)
	}
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step directory"})
		return
	}
	// Create org dir
	_ = os.Mkdir(os.Getenv("DEFAULT_ORG_PATH")+"/"+create_stack.Name, 0755)

	srcFile, err := os.Open(os.Getenv("ORGANISATION_TEMPLATE")+"/docker-compose.yml")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step copy file : "+ err.Error()})
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(os.Getenv("DEFAULT_ORG_PATH")+"/"+create_stack.Name+"/docker-compose.yml") // creates if file doesn't exist
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step copy file : "+err.Error()})
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		fmt.Println("S", err.Error())
		os.Exit(1)
	}

	err = destFile.Sync()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step copy file "+err.Error()})
		return
	}

	// Create env file in org dir
	f, err := os.Create(os.Getenv("DEFAULT_ORG_PATH")+"/"+create_stack.Name+"/.env")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step write env file "+err.Error()})
		return
	}
	defer f.Close()

	// Generate env file
	//Gen mysql password
	mysql_pwd, _ := Generate(`p0pcUb3e_[a-Z]{12}`)
	// Generate env file
	_, err = f.WriteString("MYSQL_ROOT_PASSWORD="+mysql_pwd+"\n")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step write env file "+err.Error()})
		return
	}
	_, err = f.WriteString("MYSQL_DATABASE="+os.Getenv("DEFAULT_DATABASE")+"\n")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step write env file "+err.Error()})
		return
	}

	_, err = f.WriteString("MYSQL_HOST="+create_stack.Name+"_"+os.Getenv("BASE_NAME_HOST_DB")+"\n")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step write env file "+err.Error()})
		return
	}

	_, err = f.WriteString("ORG_ORGANISATIONNAME="+create_stack.Name+"\n")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step write env file "+err.Error()})
		return
	}
	_, err = f.WriteString("ORG_DESCRIPTION="+create_stack.Description+"\n")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step write env file "+err.Error()})
		return
	}
	_, err = f.WriteString("ORG_AVATAR=jesaispas"+"\n")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step write env file "+err.Error()})
		return
	}
	_, err = f.WriteString("ORG_DOMAIN="+create_stack.Name+"."+os.Getenv("DEFAULT_DOMAIN")+"\n")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Step write env file "+err.Error()})
		return
	}
	f.Sync()

	os.Setenv("ORGANISATION", create_stack.Name)
	os.Setenv("PATH_ENV_FILE", os.Getenv("DEFAULT_ORG_PATH")+"/"+create_stack.Name+"/.env")
	t := time.Now()
	os.Setenv("CREATED_DATE", t.Format("2006-01-02 15:04:05"))

	args := []string{"stack", "deploy", "--with-registry-auth", "--compose-file",os.Getenv("DEFAULT_ORG_PATH")+"/"+create_stack.Name+"/docker-compose.yml", create_stack.Name}
	if err := exec.Command("docker", args...).Run(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(jsonErr{Code: 500, Text: "Exec Deployement  :  "+err.Error()})
		return
	}
	json.NewEncoder(w).Encode(Response{Message: "Success", Url: create_stack.Name+"."+os.Getenv("DEFAULT_DOMAIN"), Stack_name: os.Getenv("ORGANISATION")})
}

func deploy_remove(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome deploy_remove\n")
}

func deploy_status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome deploy_status\n")
}
