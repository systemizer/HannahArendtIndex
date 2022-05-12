# Hannah Arendt Archives

The Hannah Arendt Archives project integrates with the [Hannah Arendt Library of Congress collection](https://www.loc.gov/collections/hannah-arendt-papers). It runs OCR on the documents in the collections and provides text search functionality for the documents.

See [arendtindex.com](https://arendtindex.com) for production version.

## Getting Started

This project requires [Golang](https://go.dev/doc/install), [Typescript](https://www.typescriptlang.org/download), and [Sqlite](https://www.sqlitetutorial.net/download-install-sqlite/).

To run the project locally, use the following make commands

```
# install dependences
make install

# Run backend
make run-backend

# Run frontend. This will open a browser
make run-frontend
```

## How it works

This project has no external service dependencies. The backend process serves a compiled frontend and runs SQLite in memory. The SQLite database implements FTS for text-based search. There is no need to run OCR or SYNC, which populates the database (this is already done). An example database file is included in the project. A more full database is run in production. I maintain this database outside the repository, but it could be recreated using the scripts available in the repository.

## Contribution

All are welcome to contribute. I'm not actively working on the project, but if you have feature ideas, feel free to open a ticket.