<template>
	<h1>QUIZ</h1>

	<StartGame v-if="!$store.getters.getStartGame" :startGame="startGame"/>
	<QuestionGame
		:index="index"
		:clickToAnswer="clickToAnswer"
		:nextQuestion="nextQuestion" 
		v-else-if="$store.getters.getStartGame && !$store.getters.getFinishGame"
	/>

	<ResultGame v-else />

	<AlertApp v-if="errors.length" :errors="errors"  @click="errors = []"/>
</template>

<script>
import StartGame from './components/StartGame.vue'
import QuestionGame from './components/QuestionGame.vue'
import ResultGame from './components/ResultGame.vue'
import AlertApp from './components/AlertApp.vue'

export default {
	components: { StartGame, QuestionGame, ResultGame, AlertApp },

	data() {
		return {
			errors: [],
			index: 0,
		}
	},

	methods: {
		startGame() {
			if(!this.$store.getters.getPlayerName) {
				this.updateErrors('Введите ваше имя!')
				return
			}
			this.errors = []
			this.$store.commit('updateStartGame', true)
		},

		clickToAnswer(answer, index) {
			this.$store.getters.getQuestions.forEach( question => {
				question.answers.forEach( el => {
					if(el.id == this.parseProxy(answer).id) {
						el.choisen = true
						this.$store.getters.getQuestions[index].userAnswer = el.value
					} else {
						el.choisen = false
					}
				})
			})
		},

		nextQuestion() {
			if(!this.$store.getters.getQuestions[this.index].userAnswer) {
				this.updateErrors('Необходимо выбрать ответ!')
				return
			}
			this.errors = []
			this.updateAnswersToFalse();
			if(this.index + 1 != this.$store.getters.getQuestions.length) {
				this.index++
			} else {
				this.$store.commit('updateFinishGame', true)
			}
		},

		parseProxy(value) {
			return JSON.parse(JSON.stringify(value))
		},

		updateAnswersToFalse() {
			this.$store.getters.getQuestions.forEach( question => {
				question.answers.forEach( el => {
					el.choisen = false
				})
			})
		},

		updateErrors(err) {
			this.errors = [ ...this.errors, err ]
		},
	}
}
</script>

<style lang="scss">
* {
	font-family: Arial, Helvetica, sans-serif;
	margin: 0; 
	padding: 0;
	color: #DEE4EA;
}

body {
	background: #161A1D;
}

h1 {
	text-align: center;
	margin: 20px 0;
}

h2 {
	text-align: center;
}

.greenBtn {
	font-size: 22px;
	font-weight: 600;
	padding: 10px 30px;
	outline: none;
	margin: 0 auto;
	display: block;
	background: #1F845A;
	border: none;
	border-radius: 5px;
	cursor: pointer;
	transition: all .25s ease;

	&:hover {
		transform: scale(1.1);
	}
}
</style>