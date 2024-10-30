using System.ComponentModel.DataAnnotations;

namespace userProfileAPI.Models;

public class Profile
{
    [Key]
    public long UserId { get; set; }
    public int Age { get; set; }
}