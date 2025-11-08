<script setup lang="ts">
import SubprocessList from './SubprocessList.vue';
import SubprocessSpawn from './SubprocessSpawn.vue';
</script>

<template>
	<div class="flex flex-col gap-5 mt-5 p-5">
	<!-- <div class="shadow-lg border border-gray-100 rounded-xl p-5 m-5"> -->
		<SubprocessList />
		<SubprocessSpawn />
	</div>
</template>

<style scoped>
</style>