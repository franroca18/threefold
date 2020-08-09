-----------------------------------------------
Threefold Challenge Task - Francisco Rodriguez
-----------------------------------------------

This is a project to handle a full CRUDL microservice for a customer.

Base path is: localhost:9000/customers/

4 end points are accesible from this path:
- **Create or Update** PUT - localhost:9000/customers/ 
	JSON as body is expected, there are 2 options, 
	- Without idNumber, a new customer will be created
		{
			"name":"First",
			"surname":"01",
			"email":"test@gmail.com",
			"initials":"TTT",
			"mobile":"08123456"
		}
		
		Response
		{
			"code": 200,
			"message": "Successful",
			"payload": [
				{
					"idNumber": "5f305669d1ad38801baeb286",
					"name": "First",
					"surname": "01",
					"email": "test@gmail.com",
					"initials": "TTT",
					"mobile": "08123456"
				}
			],
			"status": "OK"
		}
	
	
	- With idNumber, customer matching will be update
		{
			"idNumber": "5f305669d1ad38801baeb286",
			"name":"FirstUpdate",
			"surname":"01U",
			"email":"test@gmail.com",
			"initials":"TTT",
			"mobile":"08123456"
		}
		
		Response
		{
			"code": 200,
			"message": "Successful",
			"payload": [
				{
					"idNumber": "5f305669d1ad38801baeb286",
					"name": "FirstUpdate",
					"surname": "01U",
					"email": "test@gmail.com",
					"initials": "TTT",
					"mobile": "08123456",
					"lastUptade": 1597003592
				}
			],
			"status": "OK"
		}
		
- **Read single record** GET - localhost:9000/customers/id=5f305669d1ad38801baeb286
	Response
	{
		"code": 200,
		"message": "Successful",
		"payload": [
			{
				"idNumber": "5f305669d1ad38801baeb286",
				"name": "FirstUpdate",
				"surname": "01U",
				"email": "test@gmail.com",
				"initials": "TTT",
				"mobile": "08123456",
				"lastUptade": 1597003592
			}
		],
		"status": "OK"
	}
	
- **Listing** GET - localhost:9000/customers/ or localhost:9000/customers?page=2&limit=3
	Where page and limit are to handle pagination, these parameters are optionals. By defaul page=1 and limit=2
	
	--Response localhost:9000/customers/ 
	{
		"code": 200,
		"message": "Successful",
		"payload": [
			{
				"idNumber": "5f305669d1ad38801baeb286",
				"name": "FirstUpdate",
				"surname": "01U",
				"email": "test@gmail.com",
				"initials": "TTT",
				"mobile": "08123456",
				"lastUptade": 1597003592
			},
			{
				"idNumber": "5f305ac752b978693b4c0c6e",
				"name": "Second",
				"surname": "02",
				"email": "test@gmail.com",
				"initials": "TTT",
				"mobile": "08123456"
			}
		],
		"status": "OK"
	}
	
	
	--Response localhost:9000/customers?page=2&limit=3
	{
		"code": 200,
		"message": "Successful",
		"payload": [
			{
				"idNumber": "5f305b0f52b978693b4c0c70",
				"name": "Fourth",
				"surname": "04",
				"email": "test@gmail.com",
				"initials": "TTT",
				"mobile": "08123456"
			},
			{
				"idNumber": "5f305b1a52b978693b4c0c71",
				"name": "Fifth",
				"surname": "05",
				"email": "test@gmail.com",
				"initials": "TTT",
				"mobile": "08123456"
			},
			{
				"idNumber": "5f305b2852b978693b4c0c72",
				"name": "Sixth",
				"surname": "06",
				"email": "test@gmail.com",
				"initials": "TTT",
				"mobile": "08123456"
			}
		],
		"status": "OK"
	}
	
- **Delete** DELETE - localhost:9000/customers/id=5f305b2852b978693b4c0c72
	Response
	{
		"code": 200,
		"message": "Successful",
		"payload": null,
		"status": "OK"
	}
	
	
Also, there is other end point where you can cosult **documentation** http://localhost:9000/swagger/index.html



I hope you like work I develop in less of 1 week without full time dedication due to I'm working, I know there are some points of the task incomplete, I would have liked could finish all of them but unfortunately it was not possible, there are 2 main point:
	- Docker deployment: I've had several issues with that, I couldn't make work app with mongoDB, It works fine in local but when I deploy it to docker it doesn't work, I explored different possible issues like wrong variables, Dockerfile or visibility between containers. I researched but nothing I found worked. I deployed to openshifth but didn´t work there either.
	- Sucurity, Test and sonarqube: I didn't have time to learn how those work in Golang
	
Probably this is an disadvantage to promote next stage, but after all I'm proud what I achieve in no time working with technologies I've never work before and knowledge I acquired thankful this project. Happen whatever happen, it was worh it :)


------------
	End
------------
