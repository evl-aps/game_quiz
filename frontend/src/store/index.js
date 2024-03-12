import { createStore } from 'vuex'

export default createStore({
	state: {
		playerName: '',
		startGame: false,
		finishGame: false,
		question: '',
		questionsList: [],
	},
	getters: {
		getPlayerName(state) {
			return state.playerName
		},
		getStartGame(state) {
			return state.startGame
		},
		getFinishGame(state) {
			return state.finishGame
		},
		getQuestion(state) {
			return state.question
		},
		getQuestionsList(state) {
			return state.questionsList
		},
	},
	mutations: {
		updatePlayerName(state, name) {
			state.playerName = name
		},
		updateStartGame(state, value) {
			state.startGame = value
		},
		updateFinishGame(state, value) {
			state.finishGame = value
		},
		updateQuestion(state, value) {
			state.question = value
			console.log(state.question)
		},
		updateQuestionsList(state, questionsList) {
			state.questionsList = questionsList
		}
	}
})
