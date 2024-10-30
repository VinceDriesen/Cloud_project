using SoapCore;
using Microsoft.EntityFrameworkCore;
using userProfileAPIService.Models;


var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var connectionString = "Host=postgres;Port=5433;Database=user_profile_database;Username=postgres;Password=postgres;";
builder.Services.AddDbContext<ProfileContext>(options =>
    options.UseNpgsql(connectionString, npgsqlOptions => 
        npgsqlOptions.EnableRetryOnFailure()));

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();      

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

// app.UseHttpsRedirection();

app.UseHttpsRedirection();
app.UseAuthorization();
app.MapControllers();

app.Run();