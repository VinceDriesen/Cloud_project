using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.EntityFrameworkCore;
using userProfileAPIService.Models;

namespace userProfileAPIService.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class ProfileController : ControllerBase
    {
        private readonly ProfileContext _context;

        public ProfileController(ProfileContext context)
        {
            _context = context;
        }

        // GET: api/Profile
        [HttpGet]
        public async Task<ActionResult<IEnumerable<Profile>>> GetProfiles()
        {
            return await _context.Profiles.ToListAsync();
        }

        // GET: api/Profile/5
        [HttpGet("{id}")]
        public async Task<ActionResult<Profile>> GetProfile(long id)
        {
            var profile = await _context.Profiles.FindAsync(id);

            if (profile == null)
            {
                return NotFound();
            }

            return profile;
        }

        // POST: api/Profile
        [HttpPost]
        [ValidateAntiForgeryToken] // Dit is niet nodig voor een API, maar als je het wilt behouden...
        public async Task<ActionResult<Profile>> PostProfile(Profile profile)
        {
            if (ModelState.IsValid)
            {
                _context.Profiles.Add(profile);
                await _context.SaveChangesAsync();
                return CreatedAtAction(nameof(GetProfile), new { id = profile.UserId }, profile);
            }

            return BadRequest(ModelState); // Retourneer een foutmelding als de modelstatus ongeldig is
        }

        // PUT: api/Profile/5
        [HttpPut("{id}")]
        [ValidateAntiForgeryToken] // Ook niet nodig voor een API, maar behouden als je het wilt
        public async Task<IActionResult> PutProfile(long id, Profile profile)
        {
            if (id != profile.UserId)
            {
                return BadRequest();
            }

            if (ModelState.IsValid)
            {
                _context.Entry(profile).State = EntityState.Modified;
                try
                {
                    await _context.SaveChangesAsync();
                }
                catch (DbUpdateConcurrencyException)
                {
                    if (!ProfileExists(id))
                    {
                        return NotFound();
                    }
                    else
                    {
                        throw;
                    }
                }
                return NoContent(); // Voor een succesvolle PUT
            }
            return BadRequest(ModelState);
        }

        // DELETE: api/Profile/5
        [HttpDelete("{id}")]
        public async Task<IActionResult> DeleteProfile(long id)
        {
            var profile = await _context.Profiles.FindAsync(id);
            if (profile == null)
            {
                return NotFound();
            }

            _context.Profiles.Remove(profile);
            await _context.SaveChangesAsync();

            return NoContent(); // Voor een succesvolle DELETE
        }

        private bool ProfileExists(long id)
        {
            return _context.Profiles.Any(e => e.UserId == id);
        }
    }
}
