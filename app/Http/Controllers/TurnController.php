<?php

namespace App\Http\Controllers;

use App\Models\Turn;

class TurnController extends Controller
{
    /**
     *　現プレイヤーかを確認するメソッド
     *  セキュリティの都合によりplayer_idはセッションで管理
     */
    public function show()
    {
        $turns = new Turn();
        $playerId = session('player_id');

        return $turns->getPlayer($playerId);
    }

    /**
     *  ターンの更新を行うメソッド
     */
    public function edit()
    {
        $turns = new Turn();
        $turns->change(session('player_id'));

        broadcast(new \App\Events\TurnChanged());
    }
}
