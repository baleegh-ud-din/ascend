# Use a minimal base image for efficiency and security
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the pre-built binary
COPY app/server/ascend ./server/ascend

# Copy static files and environment variables
COPY app/static ./static
COPY .env .env

# Ensure the binary has execute permissions
RUN chmod +x ./server/ascend

# Expose the required port
EXPOSE 8443

# Start the application
CMD ["./server/ascend"]
