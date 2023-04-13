# Base image
FROM golang:1.19

# Install nodemon
RUN apt-get update && apt-get install -y npm
RUN npm install -g nodemon

# Create a working directory
WORKDIR /go/src/app

# Copy the source code to the container
COPY . .

# Build the Go binary
RUN go build -o app

# Expose the port the application will listen on
EXPOSE 3000

# Set the environment variable
ENV ENVIRONMENT=prod

# Set the entry point for the container
CMD ["nodemon", "app.js"]
