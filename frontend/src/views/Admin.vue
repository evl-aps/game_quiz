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

	<div v-else>
		<div class="adminBtnsSelectTabs">
			<button
				:class="{ 'activeTab': addQuestionPage }"
				@click="addQuestionPage = true, changeQuestionPage = false"
			>
				Добавить новые вопросы
			</button>

			<button
				:class="{ 'activeTab': changeQuestionPage }"
				@click="addQuestionPage = false, changeQuestionPage = true"
			>
				Изменить текущие вопросы
			</button>
		</div>

		<AddQuestions v-if="addQuestionPage" :addQuestion="addQuestion" />
	</div>

	<AlertApp v-if="errors.length" :errors="errors"  @click="errors = []"/>
</div>
</template>

<script>
import AlertApp from '../components/AlertApp.vue'
import AddQuestions from '../components/admin/AddQuestions.vue'

export default {
	components: { AlertApp, AddQuestions },

	data() {
		return {
			name: '',
			password: '',
			errors: [],
			successLogin: false,
			questionList: [],
			addQuestionPage: false,
			changeQuestionPage: false,
		}
	},

	mounted() {
		this.successLogin = sessionStorage.getItem('successLogin')
		this.GetQuestionsList()
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
			formData.append('question', this.$store.getters.getQuestion)
			const { data } = await this.$http.post('http://localhost:8000/admin/add-question', formData)

			if(data.status != 200) {
				this.errors = [ ...this.errors, data.msg ]
			} else {
				this.$store.commit('updateQuestionsList', data.list)
			}
			this.$store.commit('updateQuestion', '')
		},

		async GetQuestionsList() {
			const { data } = await this.$http.get('http://localhost:8000/get-question')
			this.$store.commit('updateQuestionsList', data.list)
		}
	}
}
</script>

<style lang="scss">
.adminBtnsSelectTabs,
.adminContent {
	display: flex;
	align-items: baseline;
	justify-content: space-between;
	flex-wrap: wrap;
	width: 90%;
	margin: 0 auto;
}

.adminContent {
	margin-top: 20px;
	padding-top: 20px;
	border-top: 2px solid #DEE4EA;
}

.adminBtnsSelectTabs {
	button {
		cursor: pointer;
		width: 49%;
		font-weight: 900;
		background: transparent;
		border: 1px solid #DEE4EA;
		border-radius: 5px;
		outline: none;
		font-size: 54px;
		transition: all .35s ease;
		padding: 15px;

		&:hover {
			border-color: #1F845A;
			color: #1F845A;
		}
	}
}

.activeTab {
	border-color: #1F845A !important;
	color: #1F845A;
}

.adminContentItem {
	width: 49%;
	border: 1px solid #DEE4EA;
	border-radius: 5px;
	padding: 15px;
	box-sizing: border-box;
}
</style>