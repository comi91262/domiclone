<template>
<div>
    <div class="row col-md-12">
        <log></log>
        <div class="row">
            <div class="row col-md-6">
                <supply></supply>
            </div>
            <div class="row col-md-6">
                <playarea></playarea>
                <hand></hand>
            </div>
        </div>
        <trash></trash>
        <public></public>
    </div>
</div>
</template>

<script>
import Entry from './Entry.vue'
import Card from './Card.vue'
import Balloon from './Balloon.vue'
import Selection from './Selection.vue'
import Hand from './Hand.vue'
import Supply from './Supply.vue'
import PlayArea from './PlayArea.vue'
import Trash from './Trash.vue'
import Public from './Public.vue'
import Debug from './Debug.vue'
import Log from './Log.vue'

export default {
    components: {
    Entry,
    Card,
    Balloon,
    Selection,
    Hand,
    Supply,
    PlayArea,
    Trash,
    Public,
    Debug,
    Log,
    },
    props: ['mode'],
    created: function (){
        if (this.mode !== 'debug'){
            Echo.join('game')
                .here((users) => {
                    this.$store.dispatch('start');
                })
                .joining((user) => {
                })
                .leaving((user) => {
                    console.log(user.name + 'がゲームから離れました');
                })
                .listen('TurnChanged', (e) => {
                    console.log('a');
                    this.$store.dispatch('start');
                })
        } else {
            axios.post('/users').then(res => {
                this.$store.dispatch('getSupplies').then(() =>{
                    this.$store.dispatch('startActionPhase');
                });
            });
        }
    }
}
</script>

