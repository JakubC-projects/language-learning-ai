source ./.env

IMAGE_NAME=europe-central2-docker.pkg.dev/gifted-decker-395914/main/server

docker build -t $IMAGE_NAME .
docker push $IMAGE_NAME

gcloud run deploy server --region=europe-central2 --project=gifted-decker-395914 --image $IMAGE_NAME