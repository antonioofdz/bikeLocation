**Login User**
----
  Realiza el login y, en caso de ser satisfactorio, retorna el accessToken del usuario.

* **URL**

  /login

* **Method:**

  `POST`
  
* **Data Body**

  Se pasará por el cuerpo de la petición un objeto del tipo:
  ```javascript
    {
        "Username" : "example_user",
        "Password" : "example_password"
    }
  ```
* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"Token":"a01bb83c-7f5g-5321-93e9-0242ac120003"}`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

----
**Register User**
----
  Realiza el proceso de registro del usuario. ¿DEVUELVE LOS DATOS DEL PROPIO USER?

* **URL**

  /signin

* **Method:**

  `POST`
  
* **Data Body**

  Se pasará por el cuerpo de la petición un objeto del tipo:
  ```javascript
    {
        "Username" : "example_user",
        "Password" : "example_password",
        "Name" : "example_name",
        "Surname" : "example_surname",
        "Email" : "example_email"
    }
  ```

    La contraseña deberá cumplir las siguientes condiciones:
    - Minimo 8 caracteres
	- Al menos una letra mayúscula
	- Al menos una letra minucula
	- Al menos un dígito

* **Success Response:**

  * **Code:** 200 <br />
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "Username Already exists" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`

----
**GET User**
----
  Obtiene los datos básicos de un usuario.

* **URL**

  /user

* **Method:**

  `GET`
* **Data header:**
    **Content:** `{"Token":"a01bb83c-7f5g-5321-93e9-0242ac120003"}`
  
* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
     ```javascript
    {
        "Username" : "example_user",
        "Password" : "example_password",
        "Name" : "example_name",
        "Surname" : "example_surname",
        "Email" : "example_email"
    }
    ```
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`


----
**GET BIKES LIST**
----
  Obtiene el listado de bicicletas disponibles.

* **URL**

  /bikes

* **Method:**

  `GET`
  
* **Data Body**

  Se pasará por el cuerpo de la petición un objeto del tipo:

  ```javascript
    {
        "Username" : "example_user",
        "Password" : "example_password"
    }
  ```

* **Data header:**
    **Content:** `{"Token":"a01bb83c-7f5g-5321-93e9-0242ac120003"}`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** ```javascript
    [
      {
      "Id":1,
      "Model":"model1",
      "Lat":1,
      "Lon":2,
      "Address":"address1",
      "UserBook":0,
      "Booked":false,
      "DateReturn":{
      "Time":"2019-05-07T10:53:36Z",
      "Valid":true
      },
      "DateRent":{
      "Time":"0001-01-01T00:00:00Z",
      "Valid":false
      }
      },
      {
      "Id":1,
      "Model":"model1",
      "Lat":1,
      "Lon":2,
      "Address":"address1",
      "UserBook":0,
      "Booked":false,
      "DateReturn":{
      "Time":"2019-05-07T10:53:36Z",
      "Valid":true
      },
      "DateRent":{
      "Time":"2019-05-06T19:55:13Z",
      "Valid":true
      }
      }
    ]
  ```
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`


----
**BOOK Bike**
----
  Realiza una reserva de una bicicleta concreta.

* **URL**

  /bikes/book

* **Method:**

  `POST`
  
* **Data Body**

  Se pasará por el cuerpo de la petición un objeto del tipo:
  ```javascript
    {
        "Username" : "example_user",
        "Password" : "example_password"
    }
  ```
* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"Token":"a01bb83c-7f5g-5321-93e9-0242ac120003"}`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`


----
**RETURN Bike**
----
  Realiza el login y, en caso de ser satisfactorio, retorna el accessToken del usuario.

* **URL**

  /login

* **Method:**

  `POST`
  
* **Data Body**

  Se pasará por el cuerpo de la petición un objeto del tipo:
  ```javascript
    {
        "Username" : "example_user",
        "Password" : "example_password"
    }
  ```
* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{"Token":"a01bb83c-7f5g-5321-93e9-0242ac120003"}`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "User doesn't exist" }`

  OR

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "You are unauthorized to make this request." }`


