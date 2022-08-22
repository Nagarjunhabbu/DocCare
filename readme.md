- Chkdin Project Description

Follow these setup to run the Project:

    1. Get the code from Github: 
                $cd go/src
                $ git clone https://github.com/Nagarjunhabbu/chkdin.git


    2. Install docker :
            - $ sudo apt install docker.io
            - $ sudo snap install docker
            - $ sudo docker images
            - $ sudo docker ps -a
            - $ sudo docker ps  (To check for containers in a running state, use the following command)

    3. To run Mysql Queries (create table and insert data)
         - go inside project directory
              $cd {project directory}
    
         - $ docker-compose up -d
         - $ docker ps  (To get the mariaDB container Id)
         - $ sudo docker exec -it {containerId} bash

         - $mysql -u root -p
         - $use chkdin;       (to go inside chkdin DB)
         
         - run the scripts present in schema.sql file
            
    4. Run the Server: $go run main.go
  
  -------------------------------------------------------------------------------------------------------------------

  All the CRUDL operations for users  performed using Mysql DB

           To perform the CRUD operation use below link

    1. User Login - POST -  http://localhost:8000/login - (send JSOn body along with this { "name":"abc", "password":"abc12"}) 
          This will provide the AuthToken
          -use the same Auth token and send Auth token in header field for all the CRUD operation
          -auth = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VyIjp7ImlkIjowLCJuYW1lIjoiIiwicGxhY2UiOiIiLCJlbWFpbCI6IiJ9fQ.Cn9ldjHPqFPTMC3-h21OZG5J-OO13quX8Xp3E6DRn5A

    2. GET request- http://localhost:8000/api/v1/user   -  (/user - to get all user data)

    3. POST Request - http://localhost:8000/api/v1/user - (create new user send JSON body along with this { "name":"abc", "password":"abc12",  "place":"Bangalore","email":"abc@gmail.com" })

    4. PUT Request - http://localhost:8000/api/v1/user/{id} - (update specific user data send Id in header and JSON body { "place":"xyz","email":"xyz@gmail.com" })

    5. DELETE Request - http://localhost:8000/api/v1/user/{id} - (delete specific user data send UserId in header)

    6. GET Specific user Data - http://localhost:8000/api/v1/user/{id}  -  (/user/{id} - to get specified user data by sending id in request header)