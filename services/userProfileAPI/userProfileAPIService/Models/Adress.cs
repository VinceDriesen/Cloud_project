using System.ComponentModel.DataAnnotations;

namespace userProfileAPIService.Models;

public class Address
{
    [Key]
    public int AddressId { get; set; }

    [Required]
    public string Street { get; set; }

    [Required]
    public string City { get; set; }
    
    public string State { get; set; }

    [Required]
    public string PostalCode { get; set; }

    [Required]
    public string Country { get; set; }
}