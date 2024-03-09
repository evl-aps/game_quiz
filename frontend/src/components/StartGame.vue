<template>
	<div class="input">
		<label for="name">Имя участника</label>
		<input type="text" id="name" v-model="computedName" @keypress.enter="startGame">

		<button class="greenBtn" @click="startGame">START</button>
	</div>
</template>

<script>
export default {
	props: {
		startGame: Function,
	},

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