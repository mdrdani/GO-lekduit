# GO-lekduit

<img align="center" alt="Golang" width="50%" src="https://github.com/mdrdani/GO-lekduit/assets/45899199/78c8ae34-5270-456b-8853-8c4bfdc4f4a8" />
<br>

![Go](https://img.shields.io/badge/Go-%3E%3D1.15-blue)
![Echo](https://img.shields.io/badge/Echo-%3E%3D4.0-green)
![GORM](https://img.shields.io/badge/GORM-%3E%3D1.21-lightgrey)

This is a GoLang Payment project that implements a simple web application for managing user transactions and payments.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Requirements](#requirements)
- [Usage](#usage)
- [Installation](#installation)

## Introduction

The GoLang Payment Project is a web application that allows users to manage their transactions and payments. It uses the Echo web framework for building APIs and GORM for database ORM (Object-Relational Mapping).

## Features

- User Management: Create, Read, Update, and Delete users.
- Transaction Management: Create, Read, Update user transactions.
- Payment Management: Create and Read payments associated with transactions.
- JWT Token Login and Logout.

## Requirements

- Go (>=1.15)
- Echo (>=4.0)
- GORM (>=1.21)
- PostgreSQL or MySQL (or any other supported GORM database)

## Usage

Once the application is running, you can access the API endpoints using a tool like curl, Postman, or any other API client.

## Installation

```bash
git clone https://github.com/mdrdani/GO-lekduit.git
cd GO-lekduit
go mod download
go run main.go
```
