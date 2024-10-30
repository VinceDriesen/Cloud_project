using Microsoft.EntityFrameworkCore;

namespace userProfileAPI.Models;

public class ProfileContext : DbContext
{
    public ProfileContext(DbContextOptions<ProfileContext> options) : base(options)
    {
    }
    
    // protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
    // {
    //     // Use the connection string from the configuration
    //     optionsBuilder.UseNpgsql("Host=localhost; Database=user_profile_database; Username=postgres; Password=postgres");
    // }
    
    public DbSet<Profile> Profiles { get; set; } = null!;
}