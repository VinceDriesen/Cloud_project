<?php

use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('home');
});

Route::post('/change-theme', [App\Http\Controllers\ThemeController::class, 'changeTheme'])->name('change-theme');

Auth::routes();

Route::get('/signout', [App\Http\Controllers\Auth\SignoutController::class, 'index'])->name('signout');
Route::get('/home', [App\Http\Controllers\HomeController::class, 'index'])->name('home');
Route::get('/dashboard', [\App\Http\Controllers\user\DashboardController::class, 'index'])->name('dashboard');
