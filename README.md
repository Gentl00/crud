# Contact Diary
---
CLI tool for storing and managing contacts with basic CRUD operations.

## Table of Contents
---
- [Contact Diary](#contact-diary)
- [Table of Contents](#table-of-contents)
- [Prerequisites](#prerequisites)
- [Configuration](#configuration)
- [Installation](#installation)
- [Usage](#usage)
- [Basics Commands](#basics-commands)

## Prerequisites
---
- Go 1.21 or higher (required by cobra for modern Go features)
- All project dependancies are listed at [go.mod](go.mod)
- Git
- PostgreSQL

## Configuration
---
Before running contact diary, you need to create a .env file with your database credentials. 

1. copy the example file:
``` Bash
cp .env.example .env
```

## Installation 
---
1. Clone the github repository :
``` Bash
git clone https://github.com/Gentl00/crud
```
2. Install the CLI :
``` Bash
cd crud
go install ./
```

## Basics commands
---

### addcontact  
Add a contact to the diary.
#### Example-Name only :
`crud addcontact "name" "first name"`

You can also add a contact with email and phonenumber.
#### Example-Name, email, and phone :
`crud addcontact "name" "first name" --email "alice@example.com" --phone "+1234567890"`

### delete  
Delete contact by its id.
#### example :
`crud delete "id"`

### list  
List all contacts in the diary.
#### example :
`crud list`

### search  
Search contacts by name.
#### example :
`crud delete "name"`

### update  
Update a contact's detail (names, email or phonenumber) by its id.

#### example :
`crud "id" "name" "first name" --email "alice@example.com" --phone "+1234567890"`




