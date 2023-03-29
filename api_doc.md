


# Telemetry IoT
Telemetry REST API of IoT telemetry device,
  

## Informations

### Version

0.0

### Contact

Christian Jodra c.jodra14@gmail.com 

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  status

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /healthz | [get healthz](#get-healthz) | Server status |
  


###  telemetry

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /user/{user-id} | [get user user ID](#get-user-user-id) | This endpoint allows the client to get user telemetries |
| POST | /save | [post save](#post-save) | This endpoint allows the client to post the telemetry |
  


## Paths

### <span id="get-healthz"></span> Server status (*GetHealthz*)

```
GET /healthz
```

This endpoint allow to see the serves status if 200 is received everything is working OK

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-healthz-200) | OK | OK |  | [schema](#get-healthz-200-schema) |

#### Responses


##### <span id="get-healthz-200"></span> 200 - OK
Status: OK

###### <span id="get-healthz-200-schema"></span> Schema

### <span id="get-user-user-id"></span> This endpoint allows the client to get user telemetries (*GetUserUserID*)

```
GET /user/{user-id}
```

Client get the telemetry and it is saved on it's user on the database

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user-id | `path` | string | `string` |  | ✓ |  | User ID |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-user-user-id-200) | OK | OK |  | [schema](#get-user-user-id-200-schema) |
| [400](#get-user-user-id-400) | Bad Request | Bad Request |  | [schema](#get-user-user-id-400-schema) |
| [500](#get-user-user-id-500) | Internal Server Error | Internal Server Error |  | [schema](#get-user-user-id-500-schema) |

#### Responses


##### <span id="get-user-user-id-200"></span> 200 - OK
Status: OK

###### <span id="get-user-user-id-200-schema"></span> Schema

##### <span id="get-user-user-id-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="get-user-user-id-400-schema"></span> Schema

##### <span id="get-user-user-id-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="get-user-user-id-500-schema"></span> Schema

### <span id="post-save"></span> This endpoint allows the client to post the telemetry (*PostSave*)

```
POST /save
```

Client Post the telmetry and it is saved on it's user on the database

#### Consumes
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| telemetry | `body` | [ModelsTelemetry](#models-telemetry) | `models.ModelsTelemetry` | | ✓ | | Telemetry |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#post-save-200) | OK | OK |  | [schema](#post-save-200-schema) |
| [400](#post-save-400) | Bad Request | Bad Request |  | [schema](#post-save-400-schema) |
| [500](#post-save-500) | Internal Server Error | Internal Server Error |  | [schema](#post-save-500-schema) |

#### Responses


##### <span id="post-save-200"></span> 200 - OK
Status: OK

###### <span id="post-save-200-schema"></span> Schema

##### <span id="post-save-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="post-save-400-schema"></span> Schema

##### <span id="post-save-500"></span> 500 - Internal Server Error
Status: Internal Server Error

###### <span id="post-save-500-schema"></span> Schema

## Models

### <span id="models-accelerometer"></span> models.Accelerometer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| angle_x | number| `float64` |  | |  |  |
| angle_y | number| `float64` |  | |  |  |
| angle_z | number| `float64` |  | |  |  |
| g_force | number| `float64` |  | |  |  |



### <span id="models-g-p-s"></span> models.GPS


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| connected_satellites | integer| `int64` |  | |  |  |
| direction | string| `string` |  | |  |  |
| latitude | number| `float64` |  | |  |  |
| longitude | number| `float64` |  | |  |  |
| speed | number| `float64` |  | |  |  |



### <span id="models-telemetry"></span> models.Telemetry


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accelerometer | [ModelsAccelerometer](#models-accelerometer)| `ModelsAccelerometer` |  | |  |  |
| gps | [ModelsGPS](#models-g-p-s)| `ModelsGPS` |  | |  |  |
| timestamp | integer| `int64` |  | |  |  |
| user_id | string| `string` |  | |  |  |


