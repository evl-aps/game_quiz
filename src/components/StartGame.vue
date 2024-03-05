<template>
	<div class="nameInput">
		<label for="name">Имя участника</label>
		<input type="text" id="name" v-model="computedName" @keypress.enter="$store.commit('updateStartGame', true)">

		<button class="startBtn" @click="$store.commit('updateStartGame', true)">START</button>
	</div>
</template>

<script>
export default {
	computed: {
		computedName: {
			get() {
				return this.$store.getters.getPlayerName
			},
			set(name) {
				this.$store.commit('updatePlayerName', name)
				return this.$store.getters.getPlayerName
			}
		}
	}
}
</script>

<style lang="scss" scoped>
.nameInput {
	width: 300px;
	margin: 0 auto;
	border: 1px solid #DEE4EA;
	border-radius: 10px;
	padding: 30px;
	box-sizing: border-box;

	label {
		font-size: 14px;
	}

	input {
		display: block;
		width: 100%;
		margin: 2px auto 10px auto;
		font-size: 18px;
		background: #DEE4EA;
		color: #161A1D;
		padding: 10px;
		border: none;
		outline: none;
		border-radius: 10px;
		box-sizing: border-box;
	}
}

.startBtn {
	font-size: 22px;
	font-weight: 600;
	padding: 10px 30px;
	outline: none;
	margin: 0 auto;
	display: block;
	background: #1F845A;
	border: none;
	border-radius: 10px;
	cursor: pointer;
	transition: all .25s ease;

	&:hover {
		transform: scale(1.1);
	}
}
</style>