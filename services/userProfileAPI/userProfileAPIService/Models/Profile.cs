using System.ComponentModel.DataAnnotations;

namespace userProfileAPIService.Models;

public class Profile
{
    [Key]
    public long UserId { get; set; }
    public int Age { get; set; }
}