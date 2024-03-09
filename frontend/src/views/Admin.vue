<template>
<div>
	<h1>ADMIN</h1>

	<div class="input" v-if="!successLogin">
		<label for="name">Логин</label>
		<input type="text" id="name" v-model="name" @keypress.enter="login">

		<label for="pass">Пароль</label>
		<input type="password" id="pass" v-model="password" @keypress.enter="login">

		<button class="greenBtn" @click="login">Войти как админ</button>
	</div>

	<div v-else class="adminContent">
		<div class="adminContentItem">
			<h2>Добавить вопрос</h2>

			<label for="question">Вопрос</label>
			<input type="text" id="question" v-model="question">

			<button class="greenBtn" @click="addQuestion">Добавить</button>
		</div>
	</div>

	<AlertApp v-if="errors.length" :errors="errors"  @click="errors = []"/>
</div>
</template>

<script>
import AlertApp from '../components/AlertApp.vue'

export default {
	components: { AlertApp },

	data() {
		return {
			name: '',
			password: '',
			errors: [],
			successLogin: false,
			question: '',
			questionList: [],
		}
	},

	mounted() {
		this.successLogin = sessionStorage.getItem('successLogin')
	},

	methods: {
		async login() {
			const formData = new FormData()
			formData.append('name', this.name)
			formData.append('password', this.password)
			const { data } = await this.$http.post('http://localhost:8000/admin', formData)
			
			if(data.status != 200) {
				this.errors = [ ...this.errors, data.msg ]
			} else {
				sessionStorage.setItem('successLogin', true)
				this.successLogin = sessionStorage.getItem('successLogin')
			}
		},

		async addQuestion() {
			const formData = new FormData()
			formData.append('question', this.question)
			const { data } = await this.$http.post('http://localhost:8000/admin/add-question', formData)

			if(data.status != 200) {
				this.errors = [ ...this.errors, data.msg ]
			} else {
				console.log('success');
			}
		}
	}
}
</script>

<style lang="scss" scoped>
.adminContent {
	display: flex;
	align-items: center;
	justify-content: space-between;
	flex-wrap: wrap;
	width: 90%;
	margin: 0 auto;

	.adminContentItem {
		width: 48%;
		border: 1px solid #DEE4EA;
		border-radius: 5px;
		padding: 15px;
	}
}
</style>