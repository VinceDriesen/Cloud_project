using System.Collections.Generic;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using userProfileAPIService.Models;

namespace userProfileAPIService.Services;

public class ProfileService : IProfileService
{
    private readonly ProfileDbContext _context;

    public ProfileService(ProfileDbContext context)
    {
        _context = context;
    }

    public Profile GetProfileById(long userId)
    {
        return _context.Profiles.Include(p => p.Address).FirstOrDefault(p => p.UserId == userId);
    }

    public List<Profile> GetAllProfiles()
    {
        return _context.Profiles.Include(p => p.Address).ToList();
    }

    public bool CreateProfile(Profile profile)
    {
        _context.Profiles.Add(profile);
        return _context.SaveChanges() > 0;
    }

    public bool UpdateProfile(Profile profile)
    {
        _context.Profiles.Update(profile);
        return _context.SaveChanges() > 0;
    }

    public bool DeleteProfile(long userId)
    {
        var profile = _context.Profiles.Find(userId);
        if (profile == null) return false;
        _context.Profiles.Remove(profile);
        return _context.SaveChanges() > 0;
    }

    public List<Gender> GetGenders()
    {
        return Enum.GetValues(typeof(Gender)).Cast<Gender>().ToList();
    }
}