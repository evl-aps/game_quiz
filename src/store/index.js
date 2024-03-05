import { createStore } from 'vuex'

export default createStore({
	state: {
		playerName: '',
		startGame: false,
	},
	getters: {
		getPlayerName(state) {
			return state.playerName
		},
		getStartGame(state) {
			return state.startGame
		}
	},
	mutations: {
		updatePlayerName(state, name) {
			state.playerName = name
		},
		updateStartGame(state, value) {
			state.startGame = value
		}
	},
	actions: {
	},
	modules: {
	}
})
