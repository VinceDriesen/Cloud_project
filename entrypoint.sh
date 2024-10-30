#!/bin/bash

# Print statement om aan te geven dat het script start
echo "Starting entrypoint script..."

# Zorg ervoor dat de EF tools zijn geïnstalleerd
echo "Installing dotnet-ef tools..."
dotnet tool install --global dotnet-ef
export PATH=$PATH:/root/.dotnet/tools

# Print statement om te bevestigen dat de EF tools zijn geïnstalleerd
echo "dotnet-ef tools installed."

# Voer de database migraties uit
echo "Running database migrations..."
dotnet ef database update --project /app/services/userProfileAPI/userProfileAPI/userProfileAPI.csproj

# Controleer of de migratie succesvol was
if [ $? -eq 0 ]; then
    echo "Database migrations completed successfully."
else
    echo "Database migrations failed."
    exit 1
fi

# Start de ASP.NET Core applicatie
echo "Starting the ASP.NET Core application..."
exec dotnet userProfileAPI.dll
