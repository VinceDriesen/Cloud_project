<?php

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use App\Models\User;
use Illuminate\Foundation\Auth\RegistersUsers;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\DB;

class RegisterController extends Controller
{
    use RegistersUsers;

    protected $redirectTo = '/dashboard';

    public function __construct()
    {
        $this->middleware('guest');
    }

    protected function validator(array $data)
    {
        return Validator::make($data, [
            'firstname' => ['required', 'string', 'max:255'],
            'lastname' => ['required', 'string', 'max:255'],
            'email' => ['required', 'string', 'email', 'max:255', 'unique:users'],
            'password' => ['required', 'string', 'min:8', 'confirmed'],
            'password_confirmation' => ['required', 'string', 'min:8'],
        ], [
            'firstname.required' => 'Please provide a firstname',
            'lastname.required' => 'Please provide a lastname',
            'email.required' => 'Please provide an email',
            'password.required' => 'Please provide a password',
            'password_confirmation.required' => 'Please provide the same password',
            'password.min' => 'Please provide a valid password. Contain minimum of 8 characters consisting of a symbol and digit',
            'password.confirmed' => 'Verify password and password are not the same',
            'email.email' => 'Please provide a valid email',
        ]);
    }

    protected function create(array $data)
    {
        DB::beginTransaction();

        try {
            $user = User::create([
                'firstname' => $data['firstname'],
                'lastname' => $data['lastname'],
                'email' => $data['email'],
                'password' => Hash::make($data['password']),
            ]);

            $profileResult = $this->createProfile($user);

            if (!$profileResult) {
                throw new \Exception("Profiel aanmaken is mislukt.");
            }

            DB::commit();
            return $user;
        } catch (\Exception $e) {
            DB::rollBack();
            Log::error("Fout bij registratie: " . $e->getMessage());
            throw $e;
        }
    }

    public function createProfile(User $user)
    {
        Log::info("createProfile aangeroepen voor gebruiker ID: {$user->id}");

        if (!$user) {
            Log::error("User not found");
            return false;
        }

        $userId = $user->id;

        try {
            $soapClient = new \SoapClient("http://user_profile_api:5109/ProfileService.asmx?wsdl");

            $params = ['userId' => $userId];
            Log::info("SOAP Request Parameters (GetProfileById): " . json_encode($params));

            $testResponse = $soapClient->__soapCall('GetProfileById', [$params]);

            if (empty((array)$testResponse)) {
                Log::info("Profiel niet gevonden, maak een nieuw profiel aan");
                Log::info("SOAP Request Parameters (CreateProfile): " . json_encode($params));

                $mainResponse = $soapClient->__soapCall('CreateProfile', [$params]);

                if (!empty((array)$mainResponse)) {
                    Log::info("Profiel succesvol aangemaakt voor gebruiker ID: {$userId}");
                    return true;
                } else {
                    throw new \Exception("Fout bij aanmaken profiel voor gebruiker ID: {$userId}. Geen resultaat terug ontvangen.");
                }
            } else {
                throw new \Exception("Error bij aanmaken van Profiel voor gebruiker ID: {$userId}. Profiel bestaat al. Antwoord: " . json_encode($testResponse));
            }
        } catch (\Exception $e) {
            Log::error("Fout bij aanmaken profiel: " . $e->getMessage());
            return false;
        }
    }
}
