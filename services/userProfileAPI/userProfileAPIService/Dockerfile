# docker/laravel-Dockerfile
FROM docker.io/bitnami/dotnet-sdk:latest

# Install system dependencies
RUN apt-get update && apt-get install -y

# Set working directory
WORKDIR /app

# Copy the project files and the solution file (if applicable)
COPY *.csproj ./
COPY *.sln ./

# Copy Code
RUN dotnet restore
COPY . ./

# Build the application
RUN dotnet build -c Release -o out

# # Expose port 5109 and start dotnet
EXPOSE 5109
CMD ["dotnet", "run", "--urls=http://0.0.0.0:5109"]