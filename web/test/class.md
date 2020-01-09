download - from the mongodb site 
    - https://www.mongodb.com/download-center/community



optionaly robo3t / compass (compass can be installed with mongodb)
                - https://robomongo.org/download


download go connector
    - go get -u github.com/golang/dep/cmd/dep
    - install 
    - go to project and dep init
    - dep ensure -add "go.mongodb.org/mongo-driver/mongo@~1.0.0"

download postman
    - https://www.getpostman.com/downloads/


create a folder in c:/data/db



add to enviroment path the mongo lib folder



start service using mongod and the client using mongo command




In MongoDB, databases hold collections of documents.





//--------usefull commands-------------

show dbs : show the database
db : show current database
use <dbname> : create or change to db

show collections

db.inventory.insert({ item: "journal", qty: 25, status: "A", size: { h: 14, w: 21, uom: "cm" }, tags: [ "blank", "red" ] })

db.inventory.insertMany([
   { item: "journal", qty: 25, status: "A", size: { h: 14, w: 21, uom: "cm" }, tags: [ "blank", "red" ] },
   { item: "notebook", qty: 50, status: "A", size: { h: 8.5, w: 11, uom: "in" }, tags: [ "red", "blank" ] },
   { item: "paper", qty: 10, status: "D", size: { h: 8.5, w: 11, uom: "in" }, tags: [ "red", "blank", "plain" ] },
   { item: "planner", qty: 0, status: "D", size: { h: 22.85, w: 30, uom: "cm" }, tags: [ "blank", "red" ] },
   { item: "postcard", qty: 45, status: "A", size: { h: 10, w: 15.25, uom: "cm" }, tags: [ "blue" ] }
]);

db.inventory.find({})
db.inventory.find({}).pretty()
db.inventory.find( { qty: 0, status: "D" } );




// MongoDB adds an _id field with an ObjectId value if the field is not present in the document



https://docs.mongodb.com/manual/crud/

//--------usefull commands-------------



// create connection in robo3t



setup and deploy database in mlab
mongodb://<dbuser>:<dbpassword>@ds253348.mlab.com:53348/class_go











//postman walk through 
    - request 
    - response 
    - headers
    -authentication
    - enviroment
    - collection
    - publishing documentation

//online docs https://explore.postman.com/

//swagger api documentation