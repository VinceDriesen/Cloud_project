using System.ServiceModel;
namespace userProfileAPIService.Models;

[ServiceContract]
public interface IProfileService
{
    [OperationContract]
    Profile GetProfileById(long userId);

    [OperationContract]
    List<Profile> GetAllProfiles();

    [OperationContract]
    bool CreateProfile(Profile profile);

    [OperationContract]
    bool UpdateProfile(Profile profile);

    [OperationContract]
    bool DeleteProfile(long userId);

    [OperationContract]
    List<Gender> GetGenders();
}