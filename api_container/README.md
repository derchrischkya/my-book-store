# Libary API

## Functions
### User Management
- Request personal id card
- Get personal id card

### Inventory Management
- Add book to libary
- Remove book from libary

### Rent book
- Extend deadline of book rent
- Rent a book
- Return a book
    - Track conditions of the book

### Search
- Return list of books
- Search for books based on criteria (titel, author, category, isbn)
- Recommandation to may interesting books

## Deployment
### Create local docker image
`docker build libary-webserver .`

### Run docker image
`docker run -p 3000:3000 libary-webserver`