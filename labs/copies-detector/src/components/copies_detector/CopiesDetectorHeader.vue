<script setup lang="ts">
import { ref } from 'vue';
import CopiesDetectorLaunch from './CopiesDetectorLaunch.vue';

const props = defineProps<{
  httpPort: number,
}>()

const host = ref("")
const port = ref(0)
const name = ref("")

</script>

<template>
  <div class="grid grid-cols-2">
    <div class="grid grid-cols-3 gap-1 items-center">
      <div class="h-full w-full">
        <input type="text" v-model="host" class="w-full h-full border-2 rounded-xl p-1 border-gray-100 outline-none" />
      </div>
      <div class="h-full w-full">
        <input type="number" v-model="port" class="w-full h-full border-2 rounded-xl p-1 border-gray-100 outline-none" />
      </div>
      <div class="h-full w-full">
        <input type="text" v-model="name" class="w-full h-full border-2 rounded-xl p-1 border-gray-100 outline-none" />
      </div>
    </div>
    <div class="flex flex-row-reverse">
      <CopiesDetectorLaunch :http-port="httpPort" :host="host" :port="port" :name="name" />
    </div>
  </div>
</template>

<style scoped></style>