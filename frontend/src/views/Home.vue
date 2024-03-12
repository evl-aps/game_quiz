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
		async startGame() {
			const formData = new FormData()
			formData.append('name', this.$store.getters.getPlayerName)
			const { data } = await this.$http.post('http://localhost:8000/start-game', formData)
			if(data.status != 200) {
				this.errors = [...this.errors, data.msg]
				return
			}
			this.errors = []
			this.$store.commit('updateStartGame', true)
		},

		clickToAnswer(answer, index) {
			this.$store.getters.getQuestionsList.forEach( question => {
				question.answers.forEach( el => {
					if(el.id == this.parseProxy(answer).id) {
						el.choisen = true
						this.$store.getters.getQuestionsList[index].userAnswer = el.value
					} else {
						el.choisen = false
					}
				})
			})
		},

		nextQuestion() {
			if(!this.$store.getters.getQuestionsList[this.index].userAnswer) {
				this.updateErrors('Необходимо выбрать ответ!')
				return
			}
			this.errors = []
			this.updateAnswersToFalse();
			if(this.index + 1 != this.$store.getters.getQuestionsList.length) {
				this.index++
			} else {
				this.$store.commit('updateFinishGame', true)
			}
		},

		parseProxy(value) {
			return JSON.parse(JSON.stringify(value))
		},

		updateAnswersToFalse() {
			this.$store.getters.getQuestionsList.forEach( question => {
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