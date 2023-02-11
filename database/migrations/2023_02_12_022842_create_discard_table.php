<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('discard', function (Blueprint $table) {
            $table->integer('game_id');
            $table->integer('user_id');
            $table->integer('card_id');
            $table->foreign('game_id')->references('id')->on('game');
            $table->foreign('user_id')->references('id')->on('user');
            $table->foreign('card_id')->references('id')->on('card');
            $table->primary(['game_id', 'user_id', 'card_id']);
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('discard');
    }
};
