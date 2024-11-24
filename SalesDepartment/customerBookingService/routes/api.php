<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

Route::get('/bookings', function () {
    return response()->json([
        ['id' => 1, 'customer' => 'John Doe', 'flight' => 'AB123'],
        ['id' => 2, 'customer' => 'Jane Smith', 'flight' => 'CD456'],
    ]);
});

Route::get('/bookings/{id}', function ($id) {
    return response()->json([
        'id' => $id,
        'customer' => 'John Doe',
        'flight' => 'AB123',
        'details' => 'Booking details for flight AB123',
    ]);
});

Route::post('/bookings', function () {
    return response()->json(['message' => 'Booking created successfully!'], 201);
});

Route::put('/bookings/{id}', function ($id) {
    return response()->json(['message' => "Booking with ID {$id} updated successfully!"]);
});

Route::delete('/bookings/{id}', function ($id) {
    return response()->json(['message' => "Booking with ID {$id} deleted successfully!"]);
});
