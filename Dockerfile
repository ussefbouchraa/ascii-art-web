FROM golang:latest

# Image metadata
LABEL maintainer="Tamine Ybouchraa"
LABEL project="Ascii-Art-Web"
LABEL version="1.23.4"
LABEL description="A web-based ASCII art generator built with Golang."
LABEL environment="production"

WORKDIR /app

# Copy all files into the folder app
COPY . .

# Create an executable file named Asciiart
RUN go build -o Asciiart .

# Expose container port
EXPOSE 8080

# Run the compiled app
CMD ["./Asciiart"]
