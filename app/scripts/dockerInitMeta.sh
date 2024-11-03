mkdir -p ./container/meta
cp -r ../meta/ ./container

# docker run --rm -p 9000:8000 --name team-builder-container team-builder