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
        Schema::create('trash', function (Blueprint $table) {
            $table->integer('game_id');
            $table->integer('card_id');
            $table->string('name_jp', 100);
            $table->integer('coin_cost');
            $table->string('card_type', 100);
            $table->string('description', 200);
            $table->foreign('game_id')->references('id')->on('game');
            $table->foreign('card_id')->references('id')->on('card');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('trash');
    }
};
