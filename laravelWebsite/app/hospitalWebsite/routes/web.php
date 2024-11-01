<?php

use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('home');
});

Route::post('/change-theme', [App\Http\Controllers\ThemeController::class, 'changeTheme'])->name('change-theme');

Auth::routes();

Route::get('/signout', [App\Http\Controllers\Auth\SignoutController::class, 'index'])->name('signout');
Route::get('/home', [App\Http\Controllers\HomeController::class, 'index'])->name('home');
Route::get('/profileSettings', [\App\Http\Controllers\user\ProfileSettingsController::class, 'index'])->name('profileSettings');
Route::get('/dashboard', [\App\Http\Controllers\user\DashboardController::class, 'index'])->name('dashboard');
Route::post('/updateProfile', [\App\Http\Controllers\user\ProfileSettingsController::class, 'updateProfile'])->name('updateProfile');

Route::get('/doctor/login', [\App\Http\Controllers\auth\DoctorLoginController::class, 'index'])->name('doctor.login');
Route::post('/doctor/login', [\App\Http\Controllers\auth\DoctorLoginController::class, 'login'])->name('doctor.login');

Route::get('/doctor/register', [\App\Http\Controllers\auth\DoctorRegisterController::class, 'index'])->name('doctor.register');
Route::post('/doctor/register', [\App\Http\Controllers\auth\DoctorRegisterController::class, 'register'])->name('doctor.register');
