# FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
# USER app
# WORKDIR /app
# EXPOSE 5000
# EXPOSE 5001

# # Build stage
# FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
# WORKDIR /src

# # Copy project files and restore dependencies
# COPY ["services/userProfileAPI/userProfileAPI/userProfileAPI.csproj", "services/userProfileAPI/userProfileAPI/"]
# RUN dotnet restore "services/userProfileAPI/userProfileAPI/userProfileAPI.csproj"

# # Copy everything else and build
# COPY . .
# WORKDIR "/src/services/userProfileAPI/userProfileAPI"
# RUN dotnet build "userProfileAPI.csproj" -c Release -o /app/build


# # Publish the application
# RUN dotnet publish "userProfileAPI.csproj" -c Release -o /app/publish /p:UseAppHost=false

# # Runtime stage
# FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS final
# WORKDIR /app
# COPY --from=publish /app/publish .
# ENTRYPOINT ["dotnet", "userProfileAPI.dll"]



# -------------------

# FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build

# # Update en installeer eventuele vereiste pakketten
# RUN apt-get update && apt-get install -y

# # Stel de werkmap in
# WORKDIR /app

# # Voeg .NET tools pad toe aan de PATH-variabele
# ENV PATH="$PATH:/app/.dotnet/tools"

# # Installeer de dotnet-ef tool als globale tool (gebruik --tool-path)
# RUN dotnet tool install --tool-path /app/.dotnet/tools dotnet-ef --version 8.0.10

# # Kopieer het csproj-bestand en voer dotnet restore uit
# COPY ["services/userProfileAPI/userProfileAPI/userProfileAPI.csproj", "services/userProfileAPI/userProfileAPI/"]
# RUN dotnet restore "services/userProfileAPI/userProfileAPI/userProfileAPI.csproj"

# # Kopieer de rest van de bestanden en build de applicatie
# COPY services/userProfileAPI/userProfileAPI/. ./ 
# RUN dotnet build "userProfileAPI.csproj" -c Release -o out

# # Expose port 5109 en start de applicatie
# EXPOSE 5109
# CMD ["dotnet", "run", "--urls=http://0.0.0.0:5109"]



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
