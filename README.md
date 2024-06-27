# Blogify

Blogify is a comprehensive platform designed for users to create, edit, and delete blog posts. It also allows users to engage with posts through likes and comments. The project is structured following the principles of Onion Architecture, ensuring a clean and maintainable codebase.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Technologies Used](#technologies-used)
- [Setup and Installation](#setup-and-installation)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Features

- **User Authentication**: Secure sign-up and login functionality.
- **Blog Management**: Create, edit, and delete blog posts.
- **Engagement Tools**: Like and comment on posts.
- **Onion Architecture**: Structured for maintainability and scalability.

## Architecture

Blogify follows the Onion Architecture, which emphasizes the separation of concerns and dependency inversion. The architecture is divided into several layers:

- **Domain Layer**: Contains the core business logic and entities.
- **Application Layer**: Manages use cases and orchestrates the application logic.
- **Infrastructure Layer**: Handles external interactions such as databases and APIs.
- **Presentation Layer**: Manages user interface and API endpoints.

## Technologies Used

- **Backend**: Go (Golang) with Gin framework
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **API Documentation**: Postman

## Setup and Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/Blogify.git 
   cd Blogify
2. **Install dependencies:**
   ```bash
   go mod download
3. **Run the application:**
   ```bash
   go run main.go
## API Documentation

For detailed API documentation and testing, please visit the [Blogify Postman Workspace](https://documenter.getpostman.com/view/27431083/2sA3dsmDY2).

## Contributing

Contributions are welcome! Please read the [contributing guidelines](CONTRIBUTING.md) to get started.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
