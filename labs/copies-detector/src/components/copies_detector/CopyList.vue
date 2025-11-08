<script setup lang="ts">
import { useGetCopiesDetectorsIds } from '../../queries/copiesDetector';
import Copy from './Copy.vue';

const props = defineProps<{
  httpPort: number,
}>()

const { data, asyncStatus } = useGetCopiesDetectorsIds(props.httpPort)

</script>

<template>
  <div v-if="data">
    <Copy v-for="c in data" :key="c" :id="c" :http-port="httpPort" />
  </div>
</template>

<style scoped>
</style>