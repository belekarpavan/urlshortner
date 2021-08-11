# Golang URL Shortner
This Project is used to create the short url from given url.

# Feature Overview
  1. Flat File Based Database
  2. API for create short URL
  3. API for get short URL information.
  4. Redirect short url into original url.
  
## Guide

### Installation

```sh
go get github.com/belekarpavan/urlshortner
```

### Run
Go to the project directory.
 i.e workspace/src/github.com/belekarpavan/urlshortner/
```sh
go run main.go
```

### API

#### 1. /create :
  This API is used to create the short url from given url.
    
  Request body : {"url":"Your url"}
    
  #### Example
    
    
    URL          : http://localhost:8000/o/create
    
    Request body : { "url" : "https://github.com/belekarpavan/urlshortner"}
    
    Response     : {
                      "alias": "ZjVmYWE3ZT",
                      "originalURL": "https://github.com/belekarpavan/urlshortner",
                      "createdOn": "2021-08-11T09:38:04.5839015+05:30",
                      "shortURL": "http://localhost:8000/ZjVmYWE3ZT"
                    }
    
    
    
#### 2. /get :
  This API is used to get the short url details.
    
  Request body : {"url":"Your short URL"}
    
  #### Example
    
    URL          : http://localhost:8000/o/get
    
    Request body : { "url" : "http://localhost:8000/ZjVmYWE3ZT"}
    
    Response     : {
                      "alias": "ZjVmYWE3ZT",
                      "originalURL": "https://github.com/belekarpavan/urlshortner",
                      "createdOn": "2021-08-11T09:38:04.5839015+05:30",
                      "shortURL": "http://localhost:8000/ZjVmYWE3ZT"
                    }
    
    
  
#### 3. /urlAliase :
  This API is used to redirect the orginal URL.
    
   #### Example
    
      URL :  http://localhost:8000/ZjVmYWE3ZT
      
      It redirected to the https://github.com/belekarpavan/urlshortner URL.
      
   ![image](https://user-images.githubusercontent.com/42233759/128974797-d4b429d7-448e-4363-a111-3d74dda6d436.png)
      
   ![image](https://user-images.githubusercontent.com/42233759/128974901-65c8960c-0aa4-4f87-9385-de537ed27e36.png)


    
    
     
    
