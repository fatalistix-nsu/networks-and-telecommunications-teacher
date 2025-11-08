<script setup lang="ts">
import { Subprocess } from '../../model/subprocess';
import SubprocessEnd from './SubprocessEnd.vue';
import SubprocessKill from './SubprocessKill.vue';

const props = defineProps<{
  subprocess: Subprocess,
}>()

</script>

<template>
  <div class="grid grid-cols-2">
    <div class="flex flex-row gap-2">
      <div class="grid place-items-center text-gray-400">
        PID: {{ subprocess.pid }}
      </div>
      <div class="grid place-items-center text-gray-400">
        HTTP port: {{ subprocess.httpPort }}
      </div>
    </div>
    <div class="flex flex-row-reverse gap-3">
      <SubprocessKill :pid="subprocess.pid" />
      <SubprocessEnd :pid="subprocess.pid" />
    </div>
  </div>
</template>

<style scoped></style>
