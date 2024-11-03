# LoLTeamMaker
```
    __          __  ______                     __  ___      __            
   / /   ____  / / /_  __/__  ____ _____ ___  /  |/  /___ _/ /_____  _____
  / /   / __ \/ /   / / / _ \/ __ `/ __ `__ \/ /|_/ / __ `/ //_/ _ \/ ___/
 / /___/ /_/ / /___/ / /  __/ /_/ / / / / / / /  / / /_/ / ,< /  __/ /    
/_____/\____/_____/_/  \___/\__,_/_/ /_/ /_/_/  /_/\__,_/_/|_|\___/_/     v0.0.1

```
Simple LoLTeamBuilder


# Docker


## build
```
docker build -t team-builder .
```

## deploy 1
```
docker run --rm -p 8000:8000 --name team-builder-container -v ./container/meta:/app/meta team-builder
```

## deploy 2
```
cd scripts
docker-compos up
```