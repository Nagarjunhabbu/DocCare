- DocCare Project Description

Follow these setup to run the Project:

    1. Get the code from Github: 
                $cd go/src
                $ git clone https://github.com/Nagarjunhabbu/DocCare.git


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
         - $use doccare;       (to go inside chkdin DB)
         
         - run the scripts present in schema.sql file
            
    4. Run the Server: $go run main.go
  
  -------------------------------------------------------------------------------------------------------------------

  All the CRUDL operations for patient  (performed using Mysql DB)

           To perform the CRUD operation use below link

    1. Doctor Login - POST -  http://localhost:8000/login - (send JSOn body along with this { "name":"abc", "password":"abc12"}) 
          This will provide the AuthToken
          -use the same Auth token and send Auth token in header field for all the CRUD operation
          -auth = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VyIjp7ImlkIjowLCJuYW1lIjoiIiwicGxhY2UiOiIiLCJlbWFpbCI6IiJ9fQ.Cn9ldjHPqFPTMC3-h21OZG5J-OO13quX8Xp3E6DRn5A

    2. GET request- http://localhost:8000/api/v1/patient   -  (/patient - to get all patient data for that particular doctor)

    3. POST Request - http://localhost:8000/api/v1/patient - (create new Patient send JSON body along with this { "name":"abc", "place":"Bangalore" })- patient name should be unique (will create patient by mapping respective doctorId)

    4. PUT Request - http://localhost:8000/api/v1/patient/{id} - (update specific patient data, if patient belongs to respective doctor. send Id in header and JSON body { "place":"xyz","name":"xyz" })

    5. DELETE Request - http://localhost:8000/api/v1/patient/{id} - (delete specific patient data send patientId in header)

    6. GET Specific patient Data - http://localhost:8000/api/v1/patient/{id}  -  (/patient/{id} - to get specified patient data by sending id in request header)