import { createStore } from 'vuex'

export default createStore({
	state: {
		playerName: '',
		startGame: false,
		finishGame: false,
		questions: [
			{
				question: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
							Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.`,
				answers: [
					{
						id: 1,
						text: 'Вариант 1',
						value: 1,
					},
					{
						id: 2,
						text: 'Вариант 2',
						value: 2,
					},
					{
						id: 3,
						text: 'Вариант 3',
						value: 3,
					},
					{
						id: 4,
						text: 'Вариант 4',
						value: 4,
					}
				],
				correctAnswer: 1,
				userAnswer: null
			},
			{
				question: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua???`,
				answers: [
					{
						id: 1,
						text: 'Вариант 14',
						value: 1,
					},
					{
						id: 2,
						text: 'Вариант 21',
						value: 2,
					},
					{
						id: 3,
						text: 'Вариант 32',
						value: 3,
					},
					{
						id: 4,
						text: 'Вариант 43',
						value: 4,
					}
				],
				correctAnswer: 3,
				userAnswer: null,
			},
		]
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
		getQuestions(state) {
			return state.questions
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
	}
})
