# Birdazzone Full

Project for the University of Bologna Data Structure and Algorithms course
(a.y. 2020-21).

# Setup

When running both containers (via `docker-compose`) or the backend container
only, please make sure you set the `TOKEN` environment variable first. It will
be used as a bearer token for [Twitter API
v2](https://developer.twitter.com/en/support/twitter-api/v2)

# Run

Production:

```bash
docker-compose -f docker-compose.yml -f production.yml up
```

Development:

```bash
docker-compose -f docker-compose.yml -f development.yml up
```

# Authors

- [P. Ceroni](https://github.com/pazero)
- [G. Crestanello](https://github.com/crestaa)
- [F. Grisendi](https://github.com/fedegri)
- [M. Girolimetto](https://github.com/specialfish9)
- [S. Volpe](https://github.com/foxyseta)
