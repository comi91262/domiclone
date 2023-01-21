import { createStore } from 'vuex'
import phase from './store/modules/phase';
import log from './store/modules/log';
import action_phase from './store/modules/action_phase';
import buy_phase from './store/modules/buy_phase';
import clean_phase from './store/modules/clean_phase';
import player from './store/modules/player';


export default createStore({
  modules: {
      phase,
      log,
      player,
      action_phase,
      buy_phase,
      clean_phase,
  },
  strict: true,
});
