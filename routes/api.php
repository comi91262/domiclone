<?php

use Illuminate\Http\Request;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:api')->get('/user', function (Request $request) {
    return $request->user();
});



Route::get('/init_parent', 'GameController@initParent');
Route::get('/init_child', 'GameController@initChild');

Route::get('/start'      , 'GameController@start');

Route::get('/get_name' ,  'GameController@getName');
Route::get('/dummy' ,  'GameController@dummy');


Route::get('/hands',     'GameController@showHands');
Route::get('/supplies' , 'GameController@showSupplies');
Route::get('/playarea' , 'GameController@showPlayArea');
Route::get('/trashes',  'GameController@showTrashes');
Route::get('/hands_and_playarea','GameController@getHandsAndPlayArea');


Route::get('/action_phase/exist', 'GameController@containActionCards');
Route::get('/action_phase/is_action', 'GameController@isActionCards');
Route::get('/action_phase/action', 'GameController@action');

Route::get('/buy_phase/estimate' ,  'GameController@estimate');
Route::get('/buy_phase/hands' ,  'GameController@showHands');
Route::get('/buy_phase/check' ,  'GameController@checkSelectedCards');
Route::get('/buy_phase/buy' ,  'GameController@buy');

Route::get('/clean' ,  'GameController@clean');

//for debug
Route::get('/init_playdata', 'DebugController@init');

Route::get('/debug/entry', 'GameController@entryOffline');
Route::get('/debug/id' ,     'DebugController@get_id');
Route::get('/debug/hand' ,   'DebugController@get_hand');
Route::get('/debug/deck',    'DebugController@get_deck');
Route::get('/debug/discard', 'DebugController@get_discard');
Route::get('/debug/playarea','DebugController@get_playarea');
Route::get('/debug/coin','DebugController@get_coin');
Route::get('/debug/action_counts','DebugController@get_action_counts');
Route::get('/debug/buy_counts'   ,'DebugController@get_buy_counts');




