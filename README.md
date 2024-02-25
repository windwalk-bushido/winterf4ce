# Winterf4ce

## Dev setup

### Clone project

```sh
git clone https://github.com/windwalk-bushido/winterf4ce
cd winterf4ce/
```

### Run it

```sh
air server --port 55555 -c .air.toml
```

## Prod setup

### Build

```sh
docker build -t winterf4ce .
```

### Run it

```sh
docker run -p 55555:55555 winterf4ce
```

## How to write/edit websites?

Open up any file in "data" directory and write it in format:

- name
- URL
- _empty row_
