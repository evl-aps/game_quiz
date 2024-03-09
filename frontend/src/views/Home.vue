<template>
<div>
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
</div>
</template>

<script>
import StartGame from '../components/StartGame.vue'
import QuestionGame from '../components/QuestionGame.vue'
import ResultGame from '../components/ResultGame.vue'
import AlertApp from '../components/AlertApp.vue'

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