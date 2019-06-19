# _UrbanAirship_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-urbanairship.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-urbanairship)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-urbanairship/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-urbanairship)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com)

An OMG service for UrbanAirship, it allows to push messaging to multiple device.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Send Message By Tag
```coffee
>>> urbanairship send tag:'username-tag' message:'messageTest' deviceList:'[ios,android]'
{"ok": true/false,"operation_id": "operationId","push_ids": ["listOfPushIDs"]}
```
##### Send Message By Named User
```coffee
>>> urbanairship send namedUser:'username' message:'messageTest' deviceList:'[ios,android]'
{"ok": true/false,"operation_id": "operationId","push_ids": ["listOfPushIDs"]}
```
##### Send Message By Channel Id
```coffee
>>> urbanairship send channelId:'channelId' channelType:'android/ios' message:'messageTest' deviceList:'[ios,android]'
{"ok": true/false,"operation_id": "operationId","push_ids": ["listOfPushIDs"]}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)
##### Awesome ID
##### Send Message By Tag
```shell
$ omg run send -a tag=<TAG> -a message=<MESSAGE> -a deviceList=<DEVICE LIST ARRAY> -e APP_KEY=<APP_KEY> -e MASTER_SECRET=<MASTER_SECRET> 
```
##### Example
```shell
$ omg run send -a tag="rohit-tag" -a message="This is test message" -a deviceList="[\"android\"]" -e APP_KEY=<APP_KEY> -e MASTER_SECRET=<MASTER_SECRET> 
```
##### Send Message By Named User
```shell
$ omg run send -a namedUser=<NAMED_USER> -a message=<MESSAGE> -a deviceList=<DEVICE LIST ARRAY> -e APP_KEY=<APP_KEY> -e MASTER_SECRET=<MASTER_SECRET>
```
##### Send Message By Channel Id
```shell
$ omg run send -a channelId=<CHANNEL_ID> -a channelType=<CHANNEL_TYPE> -a message=<MESSAGE> -a deviceList=<DEVICE LIST ARRAY> -e APP_KEY=<APP_KEY> -e MASTER_SECRET=<MASTER_SECRET>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/urbanairship/blob/master/LICENSE).
