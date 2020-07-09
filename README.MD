
# Supercharger API Documentation

The Tesla Supercharger API is RESTful service that returns information about charging stations. The API uses standard HTTP response codes, authentication, and verbs.

A successful request will return a `200 Successful` response, an unauthorized request will return a `403 Forbidden` response, and an incorrect request will return a `404 Error` response.

# API Endpoints

## Search All Superchargers
Returns all supercharger locations.

**Endpoint:**
`GET`  `/api/stations`

**Parameters:**
None

## Search Superchargers by Zip Code
Retrieves an item with a given ID.

**Endpoint:**
`GET`  `/api/stations/:zipcode`

**Parameters:**
|Name | Type | Description |
|--|--|--|
|zipcode|string|A 5 digit zip code.|

## Search Superchargers by State.
List all superchargers in a specific state.

**Endpoint:**
`GET`  `/api/stations/:state`

 **Parameters:**
|Name | Type | Description |
|--|--|--|
|state|string|The State, formatted as a 2 letter abbreviation.|

# Models

## Station
A single Tesla Supercharger station.

**Attributes:**
|Name | Type | Description |
|--|--|--|
|Name|string|Label the item as a Shoe, Clothing, or Other.|
|URL|string|The origin or manufacturer of the item.|
|Address|Address|A brief description of the item.|
|Active|bool|Returns True if the station is active. Otherwise, False|

## Stations
A list of Tesla Supercharger stations.

**Attributes:**
|Name | Type | Description |
|--|--|--|
|Stations|[]Station|A list of `station objects.|

## Address
The address of a Tesla Supercharger station.

**Attributes:**
|Name | Type | Description |
|--|--|--|
|Street|string|The street address.|
|City|string|The city.|
|State|string|The state, abbrieviated to 2 letters.|
|Zipcode|string|The zip code|