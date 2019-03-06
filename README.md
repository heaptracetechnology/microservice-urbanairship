# UrbanAirship as a microservice
An OMG service for UrbanAirship, it allows to push messaging to multiple device


## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Send Message By Tag
```sh
$ omg run send -a tag=<TAG> -a message=<MESSAGE> -a device_list=<DEVICE LIST ARRAY> -e APP_KEY=<APP_KEY> -e MASTER_SECRET=<MASTER_SECRET> 
```
example  
omg run send -a tag="rohit-tag" -a message="This is test message" -a device_list="[\"android\"]" -e APP_KEY="_i3ZHwoUSxKJzD_oA1QuCQ" -e MASTER_SECRET="rPOZp9WsQ1i-bQV6nYJpSA" 

##### Send Message By Named User
```sh
$ omg run send -a named_user=<NAMED_USER> -a message=<MESSAGE> -a device_list=<DEVICE LIST ARRAY> -e APP_KEY=<APP_KEY> -e MASTER_SECRET=<MASTER_SECRET>
```
##### Send Message By Channel Id
```sh
$ omg run send -a channel_id=<CHANNEL_ID> -a message=<MESSAGE> -a device_list=<DEVICE LIST ARRAY> -e APP_KEY=<APP_KEY> -e MASTER_SECRET=<MASTER_SECRET>
```


## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-urbanairship .
```
### RUN
```
docker run -p 3000:3000 microservice-urbanairship
```
